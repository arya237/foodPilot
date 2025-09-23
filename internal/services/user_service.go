package services

import (
	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/arya237/foodPilot/pkg/logger"
)

type UserService interface {
	Save(user *models.User) (int, error)
	GetById(id int) (*models.User, error)
	GetByUserName(username string) (*models.User, error)
	GetAll() ([]*models.User, error)
	Delete(id int) error
	Update(new *models.User) error
}

type userService struct {
	repo   repositories.UserRepo
	logger logger.Logger
}

func NewUserService(repo repositories.UserRepo) UserService {
	return &userService{
		repo:   repo,
		logger: logger.New("userService"),
	}
}

func (u *userService) Save(user *models.User) (int, error) {
	_, err := u.GetByUserName(user.Username)

	if err != nil {
		u.logger.Info(err.Error())
		return 0, err
	}

	id, err := u.repo.SaveUser(user.Username, user.Password)

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
