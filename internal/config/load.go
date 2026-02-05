package config

import (
	"github.com/arya237/foodPilot/internal/db/postgres"
	"github.com/arya237/foodPilot/internal/db/tempdb"
	"github.com/arya237/foodPilot/internal/getways/bot"
	"github.com/arya237/foodPilot/internal/getways/email"
	"github.com/arya237/foodPilot/internal/getways/reservations/samad"
)

func New() (*Config, error) {
	config := Config{
		SamadConfig:     &samad.Config{},
		DBConfig:        &tempdb.Config{},
		MessagingConfig: &email.Config{},
		PostGresConfig:  &postgres.Config{},
		TelegramBot:     &bot.Config{},
		BaleBot:         &bot.Config{},
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
	config.PostGresConfig.Host = GetEnv("DB_HOST", "localhost")
	config.PostGresConfig.Port = GetEnv("DB_PORT", "5432")
	config.PostGresConfig.User = GetEnv("DB_USER", "postgr")
	config.PostGresConfig.DBName = GetEnv("DB_NAME", "postgres")
	config.PostGresConfig.Password = GetEnv("DB_PASSWORD", "")

	// Messenger
	config.MessagingConfig.From = GetEnv("MSG_FROM", "")
	config.MessagingConfig.Key = GetEnv("MSG_KEY", "")

	// telgram bot
	config.TelegramBot.Token = GetEnv("TELE_BOT_TOKEN", "")
	config.TelegramBot.API = GetEnv("TELE_BOT_API", "")

	// Bale bot
	config.BaleBot.Token = GetEnv("BALE_BOT_TOKEN", "")
	config.BaleBot.API = GetEnv("BALE_BOT_API", "")

	return &config, nil
}
