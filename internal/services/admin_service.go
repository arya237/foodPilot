package services

import (
	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/pkg/logger"
)

import "errors"

type AdminService interface {
	AddUser(userName, password, role string) (int, error)
	DeleteUser(id int) error
	GetUsers() ([]*models.User, error)
	GetFoods() ([]*models.Food, error)
	AddFood(foodName string) (int, error)
	DeleteFood(id int) error
}

type adminService struct {
	user   UserService
	food   FoodService
	logger logger.Logger
}

func NewAdminService(user UserService, food FoodService) AdminService {
	return &adminService{
		user:   user,
		food:   food,
		logger: logger.New("admin_service"),
	}
}

func (s *adminService) AddUser(userName, password, role string) (int, error) {

	var userRole models.UserRole

	if role == "user" {
		userRole = models.RoleUser
	} else if role == "admin" {
		userRole = models.RoleAdmin
	} else {
		s.logger.Info("role is invalid")
		return -1, errors.New("invalid role")
	}

	user := &models.User{
		Username: userName,
		Password: password,
		Role:     userRole,
		AutoSave: true,
		Token:    "empty",
	}

	id, err := s.user.Save(user)
	if err != nil {
		s.logger.Info(err.Error())
		return id, err
	}

	return id, nil
}

func (s *adminService) DeleteUser(id int) error {
	err := s.user.Delete(id)
	if err != nil {
		s.logger.Info(err.Error())
		return err
	}

	return nil
}

func (s *adminService) GetUsers() ([]*models.User, error) {
	users, err := s.user.GetAll()
	if err != nil {
		s.logger.Info(err.Error())
		return nil, err
	}
	return users, nil
}

func (s *adminService) GetFoods() ([]*models.Food, error) {
	foods, err := s.food.GetAll()
	if err != nil {
		s.logger.Info(err.Error())
		return nil, err
	}
	return foods, nil
}

func (s *adminService) AddFood(foodName string) (int, error) {
	id, err := s.food.Save(foodName)
	if err != nil {
		s.logger.Info(err.Error())
		return id, err
	}
	return id, nil
}

func (s *adminService) DeleteFood(id int) error {
	err := s.food.Delete(id)
	if err != nil {
		s.logger.Info(err.Error())
		return err
	}
	return nil
}
