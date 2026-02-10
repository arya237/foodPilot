package config

import (
	"github.com/arya237/foodPilot/internal/infrastructure/postgres"
	"github.com/arya237/foodPilot/internal/infrastructure/tempdb"
	"github.com/arya237/foodPilot/internal/infrastructure/bot"
	"github.com/arya237/foodPilot/internal/getways/email"
	"github.com/arya237/foodPilot/internal/getways/reservations/samad"
)

type Config struct {
	SamadConfig     *samad.Config
	MessagingConfig *email.Config
	DBConfig        *tempdb.Config
	PostGresConfig  *postgres.Config
	TelegramBot     *bot.Config
	BaleBot         *bot.Config
}
