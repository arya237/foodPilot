package postgres

import (
	"testing"

	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/stretchr/testify/assert"
)

func TestIdentitiesGetByProvide_Notfound(t *testing.T) {
	repo := NewIdentities(db)
	_, err := repo.GetByProvide(models.TELEGRAM, "some text")
	assert.ErrorIs(t, err, repositories.ErrorNotFound)
}

func TestIdentitiesSave_BadUserID(t *testing.T) {
	repo := NewIdentities(db)

	_, err := repo.Save(&models.Identities{
		UserID:     12000, //nonexsting
		Provider:   models.TELEGRAM,
		Identifier: "some thinf",
	})
	assert.NotNil(t, err)
}

func TestIdentitiesSave_BadProvider(t *testing.T) {
	userRepo := NewUserRepo(db)
	repo := NewIdentities(db)
	user, err := userRepo.Save(&models.User{
		Username: "test user",
		Role:     models.RoleUser,
	})
	assert.Nil(t, err)

	_, err = repo.Save(&models.Identities{
		UserID:     user.Id,
		Provider:   "bad provider",
		Identifier: "some thinf",
	})
	assert.NotNil(t, err)

	userRepo.Delete(user.Id)
}

func TestIdentitiesSaveAndGet(t *testing.T) {
	userRepo := NewUserRepo(db)
	repo := NewIdentities(db)
	user, err := userRepo.Save(&models.User{
		Username: "test user",
		Role:     models.RoleUser,
	})
	defer userRepo.Delete(user.Id)

	assert.Nil(t, err)
	identity, err := repo.Save(&models.Identities{
		UserID:     user.Id,
		Provider:   models.TELEGRAM,
		Identifier: "some thing",
	})
	assert.Nil(t, err)

	found, err := repo.GetByProvide(identity.Provider, identity.Identifier)
	assert.Nil(t, err)
	assert.Equal(t, found, identity)
}
func TestListByProvide_empty(t *testing.T) {
	repo := NewIdentities(db)
	lst, err := repo.ListByProvider(models.TELEGRAM, 0, 10)
	assert.NoError(t, err)
	assert.Empty(t, lst)
}

func TestListByProvide_BadProvider(t *testing.T) {
	repo := NewIdentities(db)
	_, err := repo.ListByProvider("bad provider", 0, 10)
	assert.ErrorIs(t, err, repositories.ErrorBadArgument)
}

func TestListByProvide_BadPage(t *testing.T) {
	repo := NewIdentities(db)
	_, err := repo.ListByProvider(models.TELEGRAM, 0, -10)
	assert.ErrorIs(t, err, repositories.ErrorBadArgument)
}

func TestListByProvide_success2itr(t *testing.T) {
	userRepo := NewUserRepo(db)
	repo := NewIdentities(db)
	user1, err := userRepo.Save(&models.User{
		Username: "test user1", Role: models.RoleUser,
	})
	assert.NoError(t, err)
	defer userRepo.Delete(user1.Id)

	user2, err := userRepo.Save(&models.User{
		Username: "test user2", Role: models.RoleUser,
	})
	assert.NoError(t, err)
	defer userRepo.Delete(user2.Id)

	identity1, err := repo.Save(&models.Identities{
		UserID:     user1.Id,
		Provider:   models.TELEGRAM,
		Identifier: "some thing1",
	})
	assert.NoError(t, err)

	identity2, err := repo.Save(&models.Identities{
		UserID:     user2.Id,
		Provider:   models.TELEGRAM,
		Identifier: "some thing2",
	})
	assert.NoError(t, err)

	_, err = repo.Save(&models.Identities{
		UserID:     user2.Id,
		Provider:   models.BALE,
		Identifier: "some thing2",
	})
	assert.NoError(t, err)

	lst, err := repo.ListByProvider(models.TELEGRAM, 1, 5)
	assert.NoError(t, err)
	assert.ElementsMatch(t, lst, []*models.Identities{identity1, identity2})

	lst, err = repo.ListByProvider(models.TELEGRAM, 2, 5)
	assert.NoError(t, err)
	assert.Empty(t, lst)

}