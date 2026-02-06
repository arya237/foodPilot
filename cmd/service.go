package cmd

import (
	"github.com/arya237/foodPilot/internal/getways/reservations"
	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/internal/services/auth"
)

type Service struct {
	User    services.UserService
	Admin   services.AdminService
	Reserve services.Reserve
	Auth    auth.Auth
}

func NewService(repo *Repo, samad reservations.ReserveFunctions) *Service {
	userService := services.NewUserService(repo.User, repo.Food, repo.Rate, repo.Cred, samad)
	return &Service{
		User:    userService,
		Admin:   services.NewAdminService(repo.User, repo.Food),
		Reserve: services.NewReserveService(repo.User, repo.Cred, userService, samad),
		Auth:    auth.New(repo.identities, repo.User),
	}
}
