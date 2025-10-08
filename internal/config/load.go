package config

import (
	"github.com/arya237/foodPilot/pkg/messaging"
	"github.com/arya237/foodPilot/pkg/reservations/samad"
)

func New() (*Config, error) {
	config := Config{
		SamadConfig: &samad.Config{},
		MessagingConfig: &messaging.Config{},
	}

	// SAMAD
	config.SamadConfig.Username = GetEnv("username", "")
	config.SamadConfig.Password = GetEnv("password", "")
	config.SamadConfig.GetProgramUrl = GetEnv("GetProgramUrl", "")
	config.SamadConfig.GetTokenUrl = GetEnv("GetTokenUrl", "")
	config.SamadConfig.ReserveUrl = GetEnv("ReserveUrl", "")

	// Messenger
	config.MessagingConfig.From = GetEnv("MSG_FROM", "") 
	config.MessagingConfig.Key = GetEnv("MSG_KEY", "")

	return &config, nil
}
