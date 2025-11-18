package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

var (
	once     sync.Once
	instance *DB
)

func NewDB(config Config) *DB {
	once.Do(func() {
		connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.Host, config.Port, config.User, config.Password, config.DBName)

		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Printf("Failed to open database: %v", err)
			return
		}

		if err := db.Ping(); err != nil {
			log.Printf("Failed to ping database: %v", err)
			return
		}

		instance = &DB{db: db}
		log.Println("Database connected successfully")
	})

	return instance
}

func (d *DB) Close() error {
	if d.db != nil {
		return d.db.Close()
	}
	return nil
}
