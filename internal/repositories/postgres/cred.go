package postgres

import (
	"database/sql"
	"log"

	"github.com/arya237/foodPilot/internal/models"
	repo "github.com/arya237/foodPilot/internal/repositories"
)

type RestaurantCredentials struct {
	db *sql.DB
}

func NewResturantCred(db *sql.DB) repo.RestaurantCredentials {
	if db == nil {
		log.Fatal("db is nil")
	}
	return &RestaurantCredentials{db: db}
}

func (r *RestaurantCredentials) Save(new *models.RestaurantCredentials) (*models.RestaurantCredentials, error) {
	query := `
		INSERT INTO restaurant_credentials (user_id, username, password, token) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id
	`

	var id int
	err := r.db.QueryRow(
		query,
		new.UserID,
		new.Username,
		new.Password,
		new.AccessToken,
	).Scan(&id)

	if err != nil {
		return nil, err
	}
	new.Id = id
	return new, nil
}

func (r *RestaurantCredentials) GetByUserID(id int) (*models.RestaurantCredentials, error) {
	query := `
		SELECT id, user_id, username, password, token 
		FROM restaurant_credentials
		WHERE user_id = $1
	`
	var cred models.RestaurantCredentials
	err := r.db.QueryRow(query, id).Scan(
		&cred.Id,
		&cred.UserID,
		&cred.Username,
		&cred.Password,
		&cred.AccessToken,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repo.ErrorNotFound
		}
		return nil, err
	}
	return &cred, nil
}
