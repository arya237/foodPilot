package cmd

import (
	"database/sql"

	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/arya237/foodPilot/internal/repositories/postgres"
)

type Repo struct {
	User       repositories.User
	Food       repositories.Food
	Rate       repositories.Rate
	Cred       repositories.RestaurantCredentials
	identities repositories.Identities
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		User:       postgres.NewUserRepo(db),
		Food:       postgres.NewFoodRepo(db),
		Rate:       postgres.NewRateRepo(db),
		Cred:       postgres.NewResturantCred(db),
		identities: postgres.NewIdentities(db),
	}
}
