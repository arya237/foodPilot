package postgres

import (
	"database/sql"
	"testing"

	"github.com/arya237/foodPilot/internal/db/postgres"
)

var (
	db *sql.DB
)
func TestMain(m *testing.M) {
	db = postgres.NewDB(&postgres.Config{
		Host:     "localhost",
		Port:     "5000",
		DBName:   "testDB",
		User:     "testuser",
		Password: "testpass",
	})
	
	m.Run()
}
