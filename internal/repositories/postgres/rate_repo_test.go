package postgres

import ( 
	"testing"

	"github.com/arya237/foodPilot/internal/db/postgres"
	"github.com/stretchr/testify/assert"
)

func TestSaveAndGet(t *testing.T) {
	db := postgres.NewDB(&postgres.Config{
		Host:     "localhost",
		Port:     "5000",
		DBName:   "testDB",
		User:     "testuser",
		Password: "testpass",
	})
	repo := NewRateRepo(db)
	t.Run("Create rate", func(t *testing.T) {
		_, err := repo.GetByUser(5)
		assert.Nil(t, err)
	})
}
