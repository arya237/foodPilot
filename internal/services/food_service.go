package services

import (
	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/arya237/foodPilot/pkg/logger"
)

type FoodService interface {
	GetAll() ([]*models.Food, error)
	Save(foodName string) (int, error)
	Delete(foodID int) error
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
	return f.repo.GetAll()
}

func (f *foodService) Save(foodName string) (int, error) {
	id, err := f.repo.Save(foodName)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (f *foodService) Delete(foodID int) error {
	err := f.repo.Delete(foodID)
	if err != nil {
		return err
	}

	return nil
}
