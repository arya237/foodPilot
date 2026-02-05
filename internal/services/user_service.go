package services

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/arya237/foodPilot/internal/getways/reservations"
	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/arya237/foodPilot/pkg/logger"
	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	SignUp(userName, password string) (*models.User, error)
	Login(userName, password string) (*models.User, error)
	ToggleAutoSave(userID int, autoSave bool) error
	ViewFoods() ([]*models.Food, error)
	RateFoods(ID string, foods map[string]int) (string, error)
	ViewRating(ID int) (map[string]int, error)

	ConnectToResturant(id int, userName, password string) error
}

type userService struct {
	userStorage repositories.User
	foodStorge  repositories.Food
	rateStorage repositories.Rate
	userCred    repositories.RestaurantCredentials
	samad       reservations.ReserveFunctions
	logger      logger.Logger
}

func NewUserService(userRepo repositories.User, foodRepo repositories.Food,
	rateRepo repositories.Rate, userCred repositories.RestaurantCredentials, samd reservations.ReserveFunctions) UserService {
	return &userService{
		userStorage: userRepo,
		userCred:    userCred,
		foodStorge:  foodRepo,
		rateStorage: rateRepo,
		logger:      logger.New("userService"),
		samad:       samd,
	}
}

func (u *userService) ConnectToResturant(id int, userName, password string) error {
	token, err := u.samad.GetAccessToken(userName, password)
	if err != nil {
		return ErrTokenGeneration
	}

	if ok := checkToken(token); !ok {
		return ErrTokenGeneration
	}

	userCred := &models.RestaurantCredentials{
		UserID:      id,
		Username:    userName,
		Password:    password,
		AccessToken: token,
		AutoSave:    true,
	}
	_, err = u.userCred.Save(userCred)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) SignUp(userName, password string) (*models.User, error) {
	existingUser, err := u.userStorage.GetByUserName(userName)
	if err == nil && existingUser != nil {
		return nil, ErrUserAlreadyExists
	}

	token, err := u.samad.GetAccessToken(userName, password)
	if err != nil {
		u.logger.Info(err.Error())
		return nil, ErrTokenGeneration
	}

	if ok := checkToken(token); !ok {
		return nil, ErrTokenGeneration
	}

	// Create new user with default role
	user := &models.User{
		Username: userName,
		Role:     models.RoleUser, // Default role is user
	}
	user, err = u.userStorage.Save(user)
	if err != nil {
		return nil, err
	}

	userCred := &models.RestaurantCredentials{
		UserID:      user.Id,
		Username:    userName,
		Password:    password,
		AccessToken: token,
		AutoSave:    true,
	}
	_, err = u.userCred.Save(userCred)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) Login(userName, password string) (*models.User, error) {

	user, err := u.userStorage.GetByUserName(userName)
	if err != nil {
		u.logger.Info(err.Error())
		return nil, ErrUserNotRegistered
	}
	cred, err := u.userCred.GetByUserID(user.Id)
	if err != nil {
		return nil, err
	}

	if cred.Password != password {
		return nil, ErrInvalidCredentials
	}

	_, err = u.samad.GetAccessToken(userName, password)
	if err != nil {
		u.logger.Info(err.Error())
		return nil, ErrTokenGeneration
	}

	return user, nil
}

func (u *userService) ToggleAutoSave(userID int, autoSave bool) error {
	user, err := u.userStorage.GetById(userID)
	if err != nil {
		u.logger.Info(err.Error())
		return err
	}

	err = u.userStorage.Update(user)
	if err != nil {
		u.logger.Info(err.Error())
		return err
	}
	return nil
}

func (u *userService) ViewFoods() ([]*models.Food, error) {
	return u.foodStorge.GetAll()
}

func (u *userService) RateFoods(userID string, foods map[string]int) (string, error) {

	foodList, err := u.foodStorge.GetAll()

	if err != nil {
		return "", err
	}

	for key, value := range foods {

		foodID, err := findFoodID(foodList, key)

		if err != nil {
			return "", err
		}

		userID, err := strconv.Atoi(userID)
		if err != nil {
			log.Println(err)
		}

		err = u.rateStorage.Save(userID, foodID, value)

		if err != nil {
			return "", err
		}
	}

	return "all Rates save successfully", nil
}

func (u *userService) ViewRating(ID int) (map[string]int, error) {
	rates, err := u.rateStorage.GetByUser(ID)
	if err != nil {
		u.logger.Info(err.Error())
		return nil, err
	}

	userRates := make(map[string]int, len(rates))
	for _, rate := range rates {
		food, _ := u.foodStorge.GetById(rate.FoodID)
		userRates[food.Name] = rate.Score
	}

	return userRates, nil
}

// ------------------------ HELPERS ----------------------------------------

func findFoodID(foods []*models.Food, foodName string) (int, error) {
	for _, food := range foods {
		if food.Name == foodName {
			return food.Id, nil
		}
	}

	return 0, errors.New("food not found")
}

func checkToken(samadToken string) bool {

	log := logger.New("check")

	token, _, err := jwt.NewParser().ParseUnverified(samadToken, jwt.MapClaims{})
	if err != nil {
		log.Info("Error parsing token (even unverified): " + err.Error())
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if expFloat, ok := claims["exp"].(float64); ok {
			expTime := time.Unix(int64(expFloat), 0)
			now := time.Now()

			if now.Before(expTime) {
				log.Info(fmt.Sprintf("exp: %v \ntoken is valid", expTime))
				//
				return true
			}

		} else {
			log.Info("No 'exp' claim found or it's not a number")
			return false
		}
	} else {
		log.Info("Failed to parse claims")
		return false
	}

	return false
}
