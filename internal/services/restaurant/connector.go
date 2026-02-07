package restaurant

import (
	"errors"

	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
)

type Connector interface {
	Connect(userID int, username, password string) error
}

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrAlreadyConnected = errors.New("you already have a connection")
	ErrAuthFailed       = errors.New("authentication failed")
)

type UserStorage interface {
	GetById(id int) (*models.User, error)
}

type resturant interface {
	GetAccessToken(studentNumber string, password string) (string, error)
}

type connector struct {
	userStorage UserStorage
	credStorage repositories.RestaurantCredentials
	resturant   interface {
		GetAccessToken(studentNumber string, password string) (string, error)
	}
}

func New(credStorage repositories.RestaurantCredentials, user UserStorage, resturant resturant) Connector {
	return &connector{
		credStorage: credStorage,
		userStorage: user,
		resturant:   resturant,
	}
}

func (c *connector) Connect(userID int, username, password string) error {
	_, err := c.userStorage.GetById(userID)
	if err != nil {
		return ErrUserNotFound
	}

	cred, err := c.credStorage.GetByUserID(userID)
	if cred != nil && err == nil {
		return ErrAlreadyConnected
	}

	token, err := c.resturant.GetAccessToken(username, password)
	if err != nil {
		return ErrAuthFailed
	}

	_, err = c.credStorage.Save(&models.RestaurantCredentials{
		UserID:      userID,
		Username:    username,
		Password:    password,
		AccessToken: token,
		AutoSave:    true,
	})

	if err != nil {
		return err
	}

	return nil
}
