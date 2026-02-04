package postgres

import (
	"testing"

	"github.com/arya237/foodPilot/internal/db/postgres"
	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/stretchr/testify/assert"
)

func TestGetResturantCred_EmptyRecode(t *testing.T) {
	db := postgres.NewDB(&postgres.Config{
		Host:     "localhost",
		Port:     "5000",
		DBName:   "testDB",
		User:     "testuser",
		Password: "testpass",
	})

	repo := NewResturantCred(db)

	_, err := repo.GetByUserID(1)
	assert.ErrorIs(t, err, repositories.ErrorNotFound)
}

func TestSaveResturantCred_InvalidUserRefrence(t *testing.T) {
	db := postgres.NewDB(&postgres.Config{
		Host:     "localhost",
		Port:     "5000",
		DBName:   "testDB",
		User:     "testuser",
		Password: "testpass",
	})

	repo := NewResturantCred(db)

	_, err := repo.Save(&models.RestaurantCredentials{
		UserID: 8000, //nonexsiting
		Username: "user name",
		Password: "password",
		Token: "token",
	})

	assert.NotNil(t, err)
}

func TestSaveAndGetResturantCred_basic(t *testing.T) {
	db := postgres.NewDB(&postgres.Config{
		Host:     "localhost",
		Port:     "5000",
		DBName:   "testDB",
		User:     "testuser",
		Password: "testpass",
	})

	userRepo := NewUserRepo(db)
	repo := NewResturantCred(db)

	user, err := userRepo.Save(&models.User{
		Username: "SaveAndGetResturantCred",
		Role: models.RoleUser,
		AutoSave: false,
	})
	assert.Nil(t, err)
	defer func ()  {
		err = userRepo.Delete(user.Id)
		assert.Nil(t, err)
	}()

	cred, err := repo.Save(&models.RestaurantCredentials{
		UserID: user.Id, 
		Username: "user name",
		Password: "password",
		Token: "token",
	})
	assert.Nil(t, err)

	getCred, err := repo.GetByUserID(user.Id)
	assert.Nil(t, err)
	assert.Equal(t, getCred, cred)
}