package auth

import (
	"errors"
	"slices"

	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
)

var (
	trusted = []models.IdProvider{models.TELEGRAM, models.BALE}
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

func (a *auth) Login(provider models.IdProvider, identifier string) (*models.User, error) {
	if !isTrustedProvider(provider) {
		return nil, ErrInvalidProvider
	}
	identity , err := a.idRepo.GetByProvide(provider, identifier)
	if err != nil {
		if errors.Is(err, repositories.ErrorNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, ErrInvalidCredintial
	}
	user, err := a.userRepo.GetById(identity.UserID)
	if err != nil {
		return nil, ErrInvalidCredintial
	}
	return user, nil
}

func (a *auth) SignUp(provider models.IdProvider, identifier string, user *models.User) (*models.User, error) {
	if !isTrustedProvider(provider) {
		return nil, ErrInvalidProvider
	}

	newUser, err := a.userRepo.Save(user)
	if err != nil {
		return nil, err
	}
	_, err = a.idRepo.Save(&models.Identities{
		Provider: provider,
		Identifier: identifier,
		UserID: newUser.Id,
	})

	if err != nil {
		return nil, err
	}
	return newUser, err
}

func isTrustedProvider(provider models.IdProvider) bool {
	return slices.Contains(trusted, provider)
}
