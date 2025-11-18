package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	mu sync.Mutex
	instance *sql.DB
)

func NewDB(config Config) *sql.DB {
	if (instance != nil){
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Failed to open database: %v", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		log.Printf("Failed to ping database: %v", err)
		return nil
	}

	instance = db
	log.Println("Database connected successfully")

	return instance
}

func Close() error {
	lastInstance := instance
	instance = nil

	if lastInstance != nil {
		return lastInstance.Close()
	}
	return nil
}
