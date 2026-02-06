package postgres

import (
	"testing"

	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/stretchr/testify/assert"
)

func TestCreate_BadRole(t *testing.T) {
	repo := NewUserRepo(db)
	_, err := repo.Save(&models.User{
		Username: "test user 1",
		Role:     "bad role",
	})
	assert.NotNil(t, err)
}

func TestCreate_ValidInput(t *testing.T) {
	repo := NewUserRepo(db)
	user := &models.User{
		Username: "test user 2",
		Role:     models.RoleUser,
	}

	getUser, err := repo.Save(user)
	assert.Nil(t, err)
	assert.Equal(t, user, getUser)
}

func TestGetByID_NotFound(t *testing.T) {
	repo := NewUserRepo(db)

	NonExsitingID := 500000
	_, err := repo.GetById(NonExsitingID)
	assert.ErrorIs(t, err, repositories.ErrorNotFound)
}

func TestGetByUserName_NotFound(t *testing.T) {
	repo := NewUserRepo(db)

	NonExsitingUsername := "a non existing user"
	_, err := repo.GetByUserName(NonExsitingUsername)
	assert.ErrorIs(t, err, repositories.ErrorNotFound)
}

func TestCreateAndGetByID(t *testing.T) {
	repo := NewUserRepo(db)
	user := &models.User{
		Username: "test user 4",
		Role:     models.RoleUser,
	}

	getUser, err := repo.Save(user)
	assert.Nil(t, err)

	findUser, err := repo.GetById(getUser.Id)
	assert.Nil(t, err)
	assert.Equal(t, user, findUser)
}

func TestCreateAndGetByUserName(t *testing.T) {
	repo := NewUserRepo(db)
	user := &models.User{
		Username: "test user 5",
		Role:     models.RoleUser,
	}

	getUser, err := repo.Save(user)
	assert.Nil(t, err)

	findUser, err := repo.GetByUserName(getUser.Username)
	assert.Nil(t, err)
	assert.Equal(t, user, findUser)
}

func TestDelete_NonExsitingUser(t *testing.T) {
	repo := NewUserRepo(db)

	NonExsitingID := 555555
	err := repo.Delete(NonExsitingID)
	assert.ErrorIs(t, err, repositories.ErrorNotFound)
}

func TestCreateDelete(t *testing.T) {
	repo := NewUserRepo(db)
	user := &models.User{
		Username: "test user 6",
		Role:     models.RoleUser,
	}
	getUser, err := repo.Save(user)
	assert.Nil(t, err)

	err = repo.Delete(getUser.Id)
	assert.Nil(t, err)
}

