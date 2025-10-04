package services

import (
	"strconv"

	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/arya237/foodPilot/pkg/logger"
	"github.com/arya237/foodPilot/pkg/reservations"
	"github.com/arya237/foodPilot/pkg/reservations/samad"
)

type UserService interface {
	SignUp(userName, password string) (*models.User, error)
	Login(userName, password string) (string, string, error)
	Save(username, password string) (int, error)
	GetById(id int) (*models.User, error)
	GetByUserName(username string) (*models.User, error)
	GetAll() ([]*models.User, error)
	Delete(id int) error
	Update(new *models.User) error
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
func (u *userService) SignUp(userName, password string) (*models.User, error) {
	// Check if user already exists
	existingUser, err := u.GetByUserName(userName)
	if err == nil && existingUser != nil {
		return nil, ErrUserAlreadyExists
	}

	// Create new user with default role
	user := &models.User{
		Username: userName,
		Password: password,
		Role:     models.RoleUser, // Default role is user
		AutoSave: false,
	}

	// Save user to database
	id, err := u.repo.SaveUser(userName, password)
	if err != nil {
		u.logger.Info(err.Error())
		return nil, ErrUserRegistration
	}

	user.Id = id
	return user, nil
}
func (u *userService) Login(userName, password string) (string, string, error) {
	// Get user by username
	user, err := u.GetByUserName(userName)
	if err != nil {
		u.logger.Info(err.Error())
		return "", "", ErrUserNotRegistered
	}

	// Validate password
	if user.Password != password {
		return "", "", ErrInvalidCredentials
	}

	// Generate access token
	token, err := u.samad.GetAccessToken(userName, password)
	if err != nil {
		u.logger.Info(err.Error())
		return "", "", ErrTokenGeneration
	}

	userID := strconv.Itoa(user.Id)
	return userID, token, nil
}

func (u *userService) Save(userName, password string) (int, error) {

	id, err := u.repo.SaveUser(userName, password)

	if err != nil {
		u.logger.Info(err.Error())
		return 0, err
	}
	return id, nil
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
