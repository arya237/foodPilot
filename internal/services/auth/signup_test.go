package auth

import (
	"testing"

	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	testUser := &models.User{
		Id: 5, Username: "target user",
	}

	tests := []struct {
		tag            string
		mockIdentities repositories.Identities
		mockUser       repositories.User
		porovide       models.IdProvider
		user           *models.User
		identifier     string
		wantErr        error
		wantUser       *models.User
	}{
		{
			tag:        "invalid provider",
			porovide:   "bad provide",
			identifier: "identifier",
			wantErr:    ErrInvalidProvider,
		},
		{
			tag: "valid telegram",
			porovide: models.TELEGRAM,
			user: testUser,
			identifier: "good identity",
			wantUser: testUser,

			mockUser: &mockUser{id: 5},
			mockIdentities: &mockIdentities{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.tag, func(t *testing.T) {
			service := New(tc.mockIdentities, tc.mockUser)
			user, err := service.SignUp(tc.porovide, tc.identifier, tc.user)

			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.wantUser, user)
		})
	}
}
