package repositories

import (
	"github.com/arya237/foodPilot/internal/models"
)

type UserRepo interface {
	SaveUser(username, password string) (int, error)
	GetUserById(id int) (*models.User, error)
	GetUserByUserName(username string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	DeleteUser(id int) error
	UpdateUser(new *models.User) error
}
