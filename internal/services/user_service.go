package services

import (
	"fmt"
	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/arya237/foodPilot/pkg/logger"
	"github.com/arya237/foodPilot/pkg/reservations"
	"github.com/arya237/foodPilot/pkg/reservations/samad"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UserService interface {
	SignUp(userName, password string) (*models.User, error)
	Login(userName, password string) (*models.User, error)
	Save(user *models.User) (int, error)
	ToggleAutoSave(userID int, autoSave bool) error
	Delete(id int) error
	GetAll() ([]*models.User, error)

	// IDEA: repo like functions -> i think it is better to delete them all :)
	GetById(id int) (*models.User, error)
	GetByUserName(username string) (*models.User, error)
	Update(new *models.User) error
}

type userService struct {
	repo   repositories.User
	samad  reservations.RequiredFunctions
	logger logger.Logger
}

func NewUserService(repo repositories.User, config *samad.Config) UserService {
	return &userService{
		repo:   repo,
		logger: logger.New("userService"),
		samad:  samad.NewSamad(config),
	}
}

// this functio need a huge refactoring.... in package repo
func (u *userService) SignUp(userName, password string) (*models.User, error) {
	// Check if user already exists
	existingUser, err := u.repo.GetUserByUserName(userName)
	if err == nil && existingUser != nil {
		return nil, ErrUserAlreadyExists
	}

	// Generate access token if Needed
	// TODO: fucking arya see this line.................
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
		Password: password,
		Role:     models.RoleUser, // Default role is user
		AutoSave: true,
		Token:    token,
	}

	// Save user to database
	user, err = u.repo.SaveUser(user)
	if err != nil {
		u.logger.Info(err.Error())
		return nil, ErrUserRegistration
	}

	return user, nil
}
func (u *userService) Login(userName, password string) (*models.User, error) {
	// Get user by username
	user, err := u.repo.GetUserByUserName(userName)
	if err != nil {
		u.logger.Info(err.Error())
		return nil, ErrUserNotRegistered
	}

	// Validate password
	if user.Password != password {
		return nil, ErrInvalidCredentials
	}

	if ok := checkToken(user.Token); !ok {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

func (u *userService) Save(user *models.User) (int, error) {

	SaveUser, err := u.repo.SaveUser(user)
	if err != nil {
		u.logger.Info(err.Error())
		return -1, err
	}

	return SaveUser.Id, nil
}

func (u *userService) GetById(id int) (*models.User, error) {
	user, err := u.repo.GetUserById(id)
	if err != nil {
		u.logger.Info(err.Error())
		return nil, err
	}
	return user, nil
}

func (u *userService) GetByUserName(username string) (*models.User, error) {
	user, err := u.repo.GetUserByUserName(username)
	if err != nil {
		u.logger.Info(err.Error())
		return nil, err
	}
	return user, nil
}

func (u *userService) GetAll() ([]*models.User, error) {
	users, err := u.repo.GetAllUsers()
	if err != nil {
		u.logger.Info(err.Error())
		return nil, err
	}
	return users, nil
}

func (u *userService) Delete(id int) error {
	err := u.repo.DeleteUser(id)
	if err != nil {
		u.logger.Info(err.Error())
	}
	return err
}

func (u *userService) Update(new *models.User) error {
	err := u.repo.UpdateUser(new)
	if err != nil {
		u.logger.Info(err.Error())
	}
	return err
}

// change auto save is better
func (u *userService) ToggleAutoSave(userID int, autoSave bool) error {
	user, err := u.repo.GetUserById(userID)
	if err != nil {
		return err
	}

	user.AutoSave = autoSave
	err = u.repo.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
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
