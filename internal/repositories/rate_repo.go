package repositories

import (
	"github.com/arya237/foodPilot/internal/models"
)

type RateRepo interface {
	SaveRate(user_id, food_id, score int) (int, error)
	GetRateByUser(user_id int) ([]*models.Rate, error)
	DeleteRate(user_id, rate_id int) error
	UpdateRate(user_id int, new *models.Rate) error
}