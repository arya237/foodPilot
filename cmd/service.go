package cmd

import (
	"github.com/arya237/foodPilot/internal/getways"
	"github.com/arya237/foodPilot/internal/getways/reservations"
	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/internal/services/admin"
	"github.com/arya237/foodPilot/internal/services/auth"
	"github.com/arya237/foodPilot/internal/services/restaurant"
)

type Service struct {
	User       services.UserService
	Admin      services.AdminService
	Reserve    services.Reserve
	Auth       auth.Auth
	Restaurant restaurant.Connector
	Notifier   admin.Notifier
}

func NewService(repo *Repo, samad reservations.ReserveFunctions, getway *getway) *Service {
	userService := services.NewUserService(repo.User, repo.Food, repo.Rate, repo.Cred, samad)
	return &Service{
		User:       userService,
		Admin:      services.NewAdminService(repo.User, repo.Food),
		Reserve:    services.NewReserveService(repo.User, repo.Cred, userService, samad),
		Auth:       auth.New(repo.identities, repo.User),
		Restaurant: restaurant.New(repo.Cred, repo.User, samad),
		Notifier: admin.NewNotifier(map[models.IdProvider]getways.Sender{
			models.TELEGRAM: getway.telegramSender,
			models.BALE:     getway.baleSender,
		}, repo.identities),
	}
}
