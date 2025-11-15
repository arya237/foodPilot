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
	ToggleAutoSave(userID int, autoSave bool) error
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
	existingUser, err := u.repo.GetByUserName(userName)
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
	user, err = u.repo.Save(user)
	if err != nil {
		u.logger.Info(err.Error())
		return nil, ErrUserRegistration
	}

	return user, nil
}
func (u *userService) Login(userName, password string) (*models.User, error) {
	// Get user by username
	user, err := u.repo.GetByUserName(userName)
	if err != nil {
		u.logger.Info(err.Error())
		return nil, ErrUserNotRegistered
	}

	// Validate password
	if user.Password != password {
		return nil, ErrInvalidCredentials
	}

	if ok := checkToken(user.Token); !ok {
		token, err := u.samad.GetAccessToken(user.Username, user.Password)
		if err != nil {
			u.logger.Info(err.Error())
			return nil, ErrTokenGeneration
		}

		user.Token = token
		err = u.repo.Update(user)
		if err != nil {
			u.logger.Info(err.Error())
			return nil, err
		}
	}

	return user, nil
}


// change auto save is better
func (u *userService) ToggleAutoSave(userID int, autoSave bool) error {
	user, err := u.repo.GetById(userID)
	if err != nil {
		return err
	}

	user.AutoSave = autoSave
	err = u.repo.Update(user)
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
