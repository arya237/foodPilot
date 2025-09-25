package services

import (
	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/arya237/foodPilot/pkg/logger"
)

type FoodService interface {
	GetAll() ([]*models.Food, error)
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
