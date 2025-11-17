package config

import (
	"github.com/arya237/foodPilot/internal/db"
	"github.com/arya237/foodPilot/pkg/messaging"
	"github.com/arya237/foodPilot/pkg/reservations/samad"
)

func New() (*Config, error) {
	config := Config{
		SamadConfig:     &samad.Config{},
		DBConfig:        &db.Config{},
		MessagingConfig: &messaging.Config{},
	}

	// Readin samad config
	config.SamadConfig.GetProgramUrl = GetEnv("GETPROGRAMURL", "")
	config.SamadConfig.GetTokenUrl = GetEnv("GETTOKENURL", "")
	config.SamadConfig.ReserveUrl = GetEnv("RESERVEURL", "")
	config.SamadConfig.GetSelfIDUrl = GetEnv("GETSELFIDURL", "")

	// Reading fack db config
	config.DBConfig.AdminUsername = GetEnv("ADMIN_USERNAME", "admin")
	config.DBConfig.AdminPassword = GetEnv("ADMIN_PASSWORD", "admin")

	// Messenger
	config.MessagingConfig.From = GetEnv("MSG_FROM", "")
	config.MessagingConfig.Key = GetEnv("MSG_KEY", "")

	return &config, nil
}
