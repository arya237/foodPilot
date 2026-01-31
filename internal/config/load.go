package config

import (
	"github.com/arya237/foodPilot/internal/db/postgres"
	"github.com/arya237/foodPilot/internal/db/tempdb"
	"github.com/arya237/foodPilot/internal/getways/email"
	"github.com/arya237/foodPilot/internal/getways/telegram"
	"github.com/arya237/foodPilot/pkg/reservations/samad"
)

func New() (*Config, error) {
	config := Config{
		SamadConfig:     &samad.Config{},
		DBConfig:        &tempdb.Config{},
		MessagingConfig: &email.Config{},
		PostGresConfig:  &postgres.Config{},
		TelegramBot:     &telegram.Config{},
	}

	// Readin samad config
	config.SamadConfig.GetProgramUrl = GetEnv("GETPROGRAMURL", "")
	config.SamadConfig.GetTokenUrl = GetEnv("GETTOKENURL", "")
	config.SamadConfig.ReserveUrl = GetEnv("RESERVEURL", "")
	config.SamadConfig.GetSelfIDUrl = GetEnv("GETSELFIDURL", "")
	config.SamadConfig.AuthHeader = GetEnv("AUTHHEADER", "")

	// Reading fack db config
	config.DBConfig.AdminUsername = GetEnv("ADMIN_USERNAME", "admin")
	config.DBConfig.AdminPassword = GetEnv("ADMIN_PASSWORD", "admin")

	// Reading postgres config
	config.PostGresConfig.Host = GetEnv("Host", "localhost")
	config.PostGresConfig.Port = GetEnv("PORT", "5432")
	config.PostGresConfig.User = GetEnv("USER", "postgres")
	config.PostGresConfig.DBName = GetEnv("Database ", "postgres")
	config.PostGresConfig.Password = GetEnv("PostGresPassword", "")

	// Messenger
	config.MessagingConfig.From = GetEnv("MSG_FROM", "")
	config.MessagingConfig.Key = GetEnv("MSG_KEY", "")

	// bot
	config.TelegramBot.Token = GetEnv("BOT_TOKEN", "")
	config.TelegramBot.API = GetEnv("BOT_API", "")

	return &config, nil
}
