package config

import (
	"github.com/arya237/foodPilot/internal/db/postgres"
	"github.com/arya237/foodPilot/internal/db/tempdb"
	"github.com/arya237/foodPilot/internal/getways/telegram"
	"github.com/arya237/foodPilot/internal/getways/email"
	"github.com/arya237/foodPilot/internal/getways/reservations/samad"
)

type Config struct {
	SamadConfig     *samad.Config
	MessagingConfig *email.Config
	DBConfig        *tempdb.Config
	PostGresConfig  *postgres.Config
	TelegramBot     *telegram.Config
}
