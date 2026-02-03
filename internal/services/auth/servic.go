package auth

import (
	"slices"

	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
)

var (
	trusted = []models.IdProvider{models.TELEGRAM}
)

type auth struct {
	idRepo   repositories.Identities
	userRepo repositories.User
}

func New(idRepo repositories.Identities, userRepo repositories.User) Auth {
	return &auth{
		idRepo: idRepo,
		userRepo: userRepo,
	}
}

func (a *auth) login(provider models.IdProvider, identifier string) (*models.User, error) {
	if !slices.Contains(trusted, provider) {
		return nil, ErrInvalidProvider
	}
	identity , err := a.idRepo.GetByProvide(provider, identifier)
	if err != nil {
		return nil, ErrInvalidCredintial
	}
	user, err := a.userRepo.GetById(identity.UserID)
	if err != nil {
		return nil, ErrInvalidCredintial
	}
	return user, nil
}

func (a *auth) SignUp(provider models.IdProvider, identifier string) (*models.User, error) {
	return nil, nil
}
