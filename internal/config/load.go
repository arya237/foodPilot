package config

import "github.com/arya237/foodPilot/pkg/food_reserve/samad"

func New() (*Config, error) {
	config := Config{
		SamadConfig: &samad.Config{},
	}

	config.SamadConfig.Username = GetEnv("username", "")
	config.SamadConfig.Password = GetEnv("password", "")
	config.SamadConfig.GetProgramUrl = GetEnv("GetProgramUrl", "")
	config.SamadConfig.GetTokenUrl = GetEnv("GetTokenUrl", "")
	config.SamadConfig.ReserveUrl = GetEnv("ReserveUrl", "")

	return &config, nil
}
