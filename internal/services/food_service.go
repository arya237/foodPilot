package services

import (
	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/arya237/foodPilot/pkg/logger"
)

type FoodService interface {
	GetAll() ([]*models.Food, error)
	Save(string) (string, error)
	Delete(id int) (string, error)
}

type foodService struct {
	repo   repositories.Food
	logger logger.Logger
}

func NewFoodService(repo repositories.Food) FoodService {
	return &foodService{
		repo:   repo,
		logger: logger.New("food_service"),
	}
}

func (f *foodService) GetAll() ([]*models.Food, error) {
	return f.repo.GetAllFood()
}

func (f *foodService) Save(foodName string) (string, error) {
	if _, err := f.repo.SaveFood(foodName); err != nil {
		f.logger.Info(err.Error())
		return "", err
	}

	return "food saved successfully", nil
}

func (f *foodService) Delete(id int) (string, error) {
	if err := f.repo.DeleteFood(id); err != nil {
		f.logger.Info(err.Error())
		return "", err
	}
	return "food deleted successfully", nil
}
