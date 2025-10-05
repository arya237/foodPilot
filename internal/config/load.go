package config

import (
	"github.com/arya237/foodPilot/internal/repositories/fakedb"
	"github.com/arya237/foodPilot/pkg/reservations/samad"
)

func New() (*Config, error) {
	config := Config{
		SamadConfig: &samad.Config{},
		DBConfig: &fakedb.Config{},
	}

	// Readin samad config
	config.SamadConfig.GetProgramUrl = GetEnv("GETPROGRAMURL", "")
	config.SamadConfig.GetTokenUrl = GetEnv("GETTOKENURL", "")
	config.SamadConfig.ReserveUrl = GetEnv("RESERVEURL", "")
	config.SamadConfig.GetSelfIDUrl = GetEnv("GETSELFIDURL", "")

	// Reading fack db config
	config.DBConfig.Username = GetEnv("ADMIN_USERNAME", "admin")
	config.DBConfig.Password = GetEnv("ADMIN_PASSWORD", "admin")

	return &config, nil
}
