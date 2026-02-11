package admin

import (
	"errors"

	"github.com/arya237/foodPilot/internal/getways"
	"github.com/arya237/foodPilot/internal/models"
)

type Notifier interface {
	Broadcast(provider models.IdProvider, msg string) error
}

type IdentityRepository interface {
	ListByProvider(provider models.IdProvider, page, pageSize int) ([]*models.Identities, error)
}

const PAGE_SIZE = 100

var (
	ErrInvalidProvider  = errors.New("invalid provider")
	ErrProviderNotFound = errors.New("provider not found")
)

type notifier struct {
	senders      map[models.IdProvider]getways.Sender
	identityRepo IdentityRepository
}

func NewNotifier(senders map[models.IdProvider]getways.Sender, repo IdentityRepository) Notifier {
	return &notifier{
		senders:      senders,
		identityRepo: repo,
	}
}

func (n notifier) Broadcast(provider models.IdProvider, msg string) error {
	if !provider.IsValid() {
		return ErrInvalidProvider
	}

	if _, ok := n.senders[provider]; !ok {
		return ErrProviderNotFound
	}

	for i := 1; ; i += 1 {
		identities, err := n.identityRepo.ListByProvider(provider, i, PAGE_SIZE)
		if err != nil {
			return err
		}
		if len(identities) == 0 {
			break
		}

		for _, identity := range identities {
			err := n.senders[provider].Send(identity.Identifier, msg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
