package config

import (
	"github.com/arya237/foodPilot/pkg/messaging"
	"github.com/arya237/foodPilot/pkg/reservations/samad"
)

type Config struct {
	SamadConfig    *samad.Config
	MessagingConfig *messaging.Config
}
