package config

import "github.com/arya237/foodPilot/pkg/reservations/samad"

func New() (*Config, error) {
	config := Config{
		SamadConfig: &samad.Config{},
	}

	config.SamadConfig.GetProgramUrl = GetEnv("GETPROGRAMURL", "")
	config.SamadConfig.GetTokenUrl = GetEnv("GETTOKENURL", "")
	config.SamadConfig.ReserveUrl = GetEnv("RESERVEURL", "")

	return &config, nil
}
