package services

import (
	"errors"
	"strconv"

	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/arya237/foodPilot/pkg/logger"
	"github.com/arya237/foodPilot/pkg/reservations"
	"github.com/arya237/foodPilot/pkg/reservations/samad"
)

type UserService interface {
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

func (u *userService) Login(userName, password string) (string, string, error) {

	user, err := u.GetByUserName(userName)
	var id int

	if err != nil {
		u.logger.Info(err.Error())
		id, err = u.Save(userName, password)

		if err != nil {
			u.logger.Info(err.Error())
			return "", "", err
		}
	} else if user.Password != password {
		return "", "", errors.New("username or password is wrong")

	}

	token, err := u.samad.GetAccessToken(userName, password)

	if err != nil {
		u.logger.Info(err.Error())
		return "", "", err
	}
	userID := strconv.Itoa(id)
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
