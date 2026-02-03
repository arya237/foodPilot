package postgres

import (
	"testing"

	"github.com/arya237/foodPilot/internal/db/postgres"
	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/stretchr/testify/assert"
)

func TestIdentitiesGetByProvide_Notfound(t *testing.T) {
	db := postgres.NewDB(&postgres.Config{
		Host:     "localhost",
		Port:     "5000",
		DBName:   "testDB",
		User:     "testuser",
		Password: "testpass",
	})

	repo := NewIdentities(db)
	_, err := repo.GetByProvide(models.TELEGRAM, "some text")
	assert.ErrorIs(t, err, repositories.ErrorNotFound)
}

func TestIdentitiesSave_BadUserID(t *testing.T) {
	db := postgres.NewDB(&postgres.Config{
		Host:     "localhost",
		Port:     "5000",
		DBName:   "testDB",
		User:     "testuser",
		Password: "testpass",
	})

	repo := NewIdentities(db)

	_, err := repo.Save(&models.Identities{
		UserID: 12000, //nonexsting
		Provider: models.TELEGRAM,
		Identifier: "some thinf",
	})
	assert.NotNil(t, err)
}

func TestIdentitiesSave_BadProvider(t *testing.T) {
	db := postgres.NewDB(&postgres.Config{
		Host:     "localhost",
		Port:     "5000",
		DBName:   "testDB",
		User:     "testuser",
		Password: "testpass",
	})

	userRepo := NewUserRepo(db)
	repo := NewIdentities(db)
	user, err := userRepo.Save(&models.User{
		Username: "test user",
		Role: models.RoleUser,
	})
	assert.Nil(t, err)

	_, err = repo.Save(&models.Identities{
		UserID: user.Id,
		Provider: "bad provider",
		Identifier: "some thinf",
	})
	assert.NotNil(t, err)

	userRepo.Delete(user.Id)
}

func TestIdentitiesSaveAndGet(t *testing.T) {
	db := postgres.NewDB(&postgres.Config{
		Host:     "localhost",
		Port:     "5000",
		DBName:   "testDB",
		User:     "testuser",
		Password: "testpass",
	})

	userRepo := NewUserRepo(db)
	repo := NewIdentities(db)
	user, err := userRepo.Save(&models.User{
		Username: "test user",
		Role: models.RoleUser,
	})
	defer userRepo.Delete(user.Id)

	assert.Nil(t, err)
	identity, err := repo.Save(&models.Identities{
		UserID: user.Id,
		Provider: models.TELEGRAM,
		Identifier: "some thing",
	})
	assert.Nil(t, err)

	found, err := repo.GetByProvide(identity.Provider, identity.Identifier)
	assert.Nil(t, err)
	assert.Equal(t, found, identity)
}
