package restaurant

import (
	"errors"
	"testing"

	"github.com/arya237/foodPilot/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) GetById(id int) (*models.User, error) {
	args := m.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	user, _ := args.Get(0).(*models.User)
	return user, args.Error(1)
}

func (m *MockUserRepo) Update(new *models.User) error {
	args := m.Called(new)
	return args.Error(0)
}

// --------------------------------------------------
type MockCredRepo struct {
	mock.Mock
}

func (m *MockCredRepo) GetByUserID(id int) (*models.RestaurantCredentials, error) {
	args := m.Called(id)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	cred, _ := args.Get(0).(*models.RestaurantCredentials)
	return cred, args.Error(1)
}

func (m *MockCredRepo) Save(c *models.RestaurantCredentials) (*models.RestaurantCredentials, error) {
	args := m.Called(c)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	cred, _ := args.Get(0).(*models.RestaurantCredentials)
	return cred, args.Error(1)
}

// --------------------------------------------------
type MockRestaurant struct {
	mock.Mock
}

func (m *MockRestaurant) GetAccessToken(username, password string) (string, error) {
	args := m.Called(username, password)
	if args.Error(1) != nil {
		return "", args.Error(1)
	}
	return args.String(0), args.Error(1)
}

// --------------------------------------------------
func TestConnect(t *testing.T) {
	testsCases := []struct {
		tag      string
		setup    func(u *MockUserRepo, c *MockCredRepo, r *MockRestaurant)
		userID   int
		username string
		password string
		wantErr  error
	}{
		{
			tag:      "non exsting user",
			userID:   5,
			username: "username",
			password: "password",
			setup: func(u *MockUserRepo, c *MockCredRepo, r *MockRestaurant) {
				u.On("GetById", 5).
					Return(nil, errors.New("db error")).Once()
			},
			wantErr: ErrUserNotFound,
		},
		{
			tag:      "user has connection",
			userID:   1,
			username: "username",
			password: "password",
			setup: func(u *MockUserRepo, c *MockCredRepo, r *MockRestaurant) {
				u.On("GetById", 1).
					Return(&models.User{Id: 1}, nil).
					Once()

				c.On("GetByUserID", 1).
					Return(&models.RestaurantCredentials{}, nil).
					Once()
			},
			wantErr: ErrAlreadyConnected,
		},
		{
			tag:      "invalid credentials",
			userID:   1,
			username: "username",
			password: "password",
			setup: func(u *MockUserRepo, c *MockCredRepo, r *MockRestaurant) {
				u.On("GetById", 1).
					Return(&models.User{}, nil).
					Once()

				c.On("GetByUserID", 1).
					Return(nil, nil).
					Once()

				r.On("GetAccessToken", "username", "password").
					Return("", errors.New("auth error")).
					Once()
			},
			wantErr: ErrAuthFailed,
		},
		{
			tag:      "succses full",
			userID:   1,
			username: "u",
			password: "p",
			setup: func(u *MockUserRepo, c *MockCredRepo, r *MockRestaurant) {
				u.On("GetById", 1).
					Return(&models.User{}, nil).
					Once()

				c.On("GetByUserID", 1).
					Return(nil, nil).
					Once()

				r.On("GetAccessToken", "u", "p").
					Return("token", nil).
					Once()

				c.On("Save", mock.AnythingOfType("*models.RestaurantCredentials")).
					Return(&models.RestaurantCredentials{}, nil).
					Once()

				u.On("Update", mock.AnythingOfType("*models.User")).
					Return(nil).
					Once()
			},
		},
	}

	for _, tc := range testsCases {
		t.Run(tc.tag, func(t *testing.T) {

			userRepo := new(MockUserRepo)
			credRepo := new(MockCredRepo)
			restaurant := new(MockRestaurant)

			if tc.setup != nil {
				tc.setup(userRepo, credRepo, restaurant)
			}

			c := New(credRepo, userRepo, restaurant)
			err := c.Connect(tc.userID, tc.username, tc.password)

			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
			} else {
				assert.NoError(t, err)
			}
			userRepo.AssertExpectations(t)
			credRepo.AssertExpectations(t)
			restaurant.AssertExpectations(t)
		})
	}
}
