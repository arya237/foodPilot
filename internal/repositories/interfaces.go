package repositories

import (
	"errors"

	"github.com/arya237/foodPilot/internal/models"
)

type User interface {
	Save(newUser *models.User) (*models.User, error)
	GetById(id int) (*models.User, error)
	GetByUserName(username string) (*models.User, error)
	GetAll() ([]*models.User, error)
	Delete(id int) error
	Update(new *models.User) error
}

type Rate interface {
	Save(userID, foodID, score int) error
	GetByUser(userID int) ([]*models.Rate, error)
	Delete(userID, foodID int) error
	Update(userID int, new *models.Rate) error
}

type Food interface {
	Save(name string) (int, error)
	GetById(id int) (*models.Food, error)
	GetAll() ([]*models.Food, error)
	Delete(id int) error
	Update(new *models.Food) error
}

type RestaurantCredentials interface {
	Save(r *models.RestaurantCredentials)(*models.RestaurantCredentials, error)
	GetByUserID(id int)(*models.RestaurantCredentials, error)
}

type Identities interface {
	Save(new *models.Identities)(*models.IdProvider, error)
	GetByProvide(provide models.IdProvider, identifier string) (*models.IdProvider, error)
	Delete(id int)
}
// --------------- Errors -------------------------------
var (
	ErrorDuplicate = errors.New("duplicate")
	ErrorNotFound  = errors.New("there is no entity in database")
)
