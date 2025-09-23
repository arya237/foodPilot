package repositories

import (
	"github.com/arya237/foodPilot/internal/models"
)

type FoodRepo interface {
	SaveFood(name string) (int, error)
	GetFoodById(id int) (*models.Food, error)
	GetAllFood() ([]*models.Food, error)
	DeleteFood(id int) error
	UpdateFood(new *models.Food) error
}