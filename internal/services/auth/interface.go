package auth

import (
	"errors"

	"github.com/arya237/foodPilot/internal/models"
)

type Auth interface {
	login(provider models.IdProvider, identifier string) (*models.User, error)
	SignUp(provider models.IdProvider, identifier string) (*models.User, error)
}

var (
	ErrInvalidProvider = errors.New("this provider is not trusted")
	ErrInvalidCredintial = errors.New("credential is invalid")
)
