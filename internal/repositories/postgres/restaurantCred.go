package postgres

import (
	"database/sql"

	"github.com/arya237/foodPilot/internal/models"
	repo "github.com/arya237/foodPilot/internal/repositories"
)

type RestaurantCredentials struct {
	db *sql.DB
}

func NewResturantCred(db *sql.DB) repo.RestaurantCredentials {
	return &RestaurantCredentials{db: db}
}

func (r *RestaurantCredentials) Save(re *models.RestaurantCredentials) (*models.RestaurantCredentials, error) {
	return nil, nil
}
func (r *RestaurantCredentials) GetByUserID(id int) (*models.RestaurantCredentials, error) {
	return nil, nil
}
