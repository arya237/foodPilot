package admin

import (
	"errors"
	"log"
	"testing"

	"github.com/arya237/foodPilot/internal/getways"
	"github.com/arya237/foodPilot/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestBrodcast(t *testing.T) {
	dbErr := errors.New("db error")
	sendErr := errors.New("sender error")

	testsCases := []struct {
		tag      string
		senders  map[models.IdProvider]getways.Sender
		mockRepo IdentityRepository
		prodiver models.IdProvider
		msg      string
		wantErr  error
	}{
		{
			tag:      "invalid provider",
			prodiver: "bad provider",
			msg:      "hello",
			wantErr:  ErrInvalidProvider,
		},
		{
			tag: "prodiver not found",
			senders: map[models.IdProvider]getways.Sender{
				models.BALE: &mockSender{},
			},
			prodiver: models.TELEGRAM,
			msg:      "hello",
			wantErr:  ErrProviderNotFound,
		},
		{
			tag: "db error",
			senders: map[models.IdProvider]getways.Sender{
				models.TELEGRAM: &mockSender{
					err: dbErr,
				},
			},
			mockRepo: &mockIdentity{err: dbErr},
			prodiver: models.TELEGRAM,
			msg:      "hello",
			wantErr:  dbErr,
		},
		{
			tag: "sender error",
			senders: map[models.IdProvider]getways.Sender{
				models.TELEGRAM: &mockSender{
					err: sendErr,
				},
			},
			mockRepo: &mockIdentity{
				cnt: 2,
				slice: []*models.Identities{
					{ID: 5, Identifier: "moew"},
				},
			},
			prodiver: models.TELEGRAM,
			msg:      "hello",
			wantErr:  sendErr,
		},
		{
			tag: "success",
			senders: map[models.IdProvider]getways.Sender{
				models.TELEGRAM: &mockSender{},
			},
			mockRepo: &mockIdentity{
				cnt: 1,
				slice: []*models.Identities{
					{ID: 5, Identifier: "moew"},
				},
			},
			prodiver: models.TELEGRAM,
			msg:      "hello",
			wantErr: nil,
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.tag, func(t *testing.T) {
			n := NewNotifier(tc.senders, tc.mockRepo)
			err := n.Broadcast(tc.prodiver, tc.msg)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
			} else {
				assert.NoError(t, err)
			}

		})
	}
}

// ================= helpers ===========
type mockSender struct {
	sentMsg string
	err     error
}

func (m *mockSender) Send(to, msg string) error {
	m.sentMsg = msg
	return m.err
}

type mockIdentity struct {
	cnt   int
	slice []*models.Identities
	err   error
}

func (m *mockIdentity) ListByProvider(provider models.IdProvider, page, pageSize int) ([]*models.Identities, error) {
	log.Println("inside mock", m.cnt)
	if m.cnt <= 0 {
		return nil, m.err
	}
	m.cnt -= 1 
	log.Println("inside mock2", m.cnt)
	return m.slice, m.err
}
