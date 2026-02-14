package auth

import (
	"testing"

	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	testUser := &models.User{
		Id: 5, Username: "target user",
	}

	tests := []struct {
		tag            string
		mockIdentities repositories.Identities
		mockUser       repositories.User
		porovide       models.IdProvider
		identifier     string
		wantErr        error
		wantUser       *models.User
	}{
		{
			tag:        "invalid provider",
			porovide:   "somthing bad",
			identifier: "some thing",
			wantErr:    ErrInvalidProvider,
		},
		{
			tag:            "identifier not exist",
			mockIdentities: &mockIdentities{},
			porovide:       models.TELEGRAM,
			identifier:     "bad id",
			wantErr:        ErrUserNotFound,
		},
		{
			tag: "valid telegram login",
			mockIdentities: &mockIdentities{
				list: []*models.Identities{
					{UserID: 5, Provider: models.TELEGRAM, Identifier: "good id"},
				},
			},
			mockUser: &mockUser{
				list: []*models.User{
					testUser,
				},
			},
			porovide:   models.TELEGRAM,
			identifier: "good id",
			wantUser:   testUser,
		},
	}

	for _, tc := range tests {
		t.Run(tc.tag, func(t *testing.T) {
			service := New(tc.mockIdentities, tc.mockUser)
			user, err := service.Login(tc.porovide, tc.identifier)

			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.wantUser, user)
		})
	}
}

type mockIdentities struct {
	err  error
	list []*models.Identities
}

func (m *mockIdentities) Save(new *models.Identities) (*models.Identities, error) {
	return nil, nil
}
func (m *mockIdentities) GetByProvide(provide models.IdProvider, identifier string) (*models.Identities, error) {
	if m.err != nil {
		return nil, m.err
	}

	for _, val := range m.list {
		if val.Provider == provide && val.Identifier == identifier {
			return val, nil
		}
	}
	return nil, repositories.ErrorNotFound
}

type mockUser struct {
	id   int
	err  error
	list []*models.User
}

func (m *mockUser) Save(newUser *models.User) (*models.User, error) {
	newUser.Id = m.id
	m.id++
	return newUser, nil
}
func (m *mockUser) GetById(id int) (*models.User, error) {
	for _, val := range m.list {
		if val.Id == id {
			return val, nil
		}
	}
	return nil, nil
}
func (m *mockUser) GetByUserName(username string) (*models.User, error) {
	return nil, nil
}
func (m *mockUser) GetAll() ([]*models.User, error) {
	return nil, nil
}
func (m *mockUser) Delete(id int) error {
	return nil
}
func (m *mockUser) Update(new *models.User) error {
	return nil
}
func (i *mockIdentities) ListByProvider(provider models.IdProvider, page, pageSize int) ([]*models.Identities, error) {
	//TODO
	return nil, nil
}
