package delivery

import (
	"log"
	"time"

	"github.com/arya237/foodPilot/internal/delivery/bot"
	"github.com/arya237/foodPilot/internal/delivery/web"
	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/internal/services/admin"
	"github.com/arya237/foodPilot/internal/services/auth"
	"github.com/arya237/foodPilot/internal/services/restaurant"
	tele "gopkg.in/telebot.v3"
)

const (
	DELIVERY_OPTIONS = 3
	TOKEN_EXP        = 24 * time.Hour
)

type NeededServises struct {
	User       services.UserService
	Admin      services.AdminService
	Resrve     services.Reserve
	Auth       auth.Auth
	Restaurant restaurant.Connector
	Notifier   admin.Notifier
}

func Start(services *NeededServises, teleBot *tele.Bot, baleBot *tele.Bot) error {

	ch := make(chan any)

	go func() {
		err := web.Start(TOKEN_EXP, services.User, services.Admin, services.Resrve, services.Notifier)
		log.Println("delivery web:", err)
		ch <- true
	}()

	go func() {
		err := bot.Start(teleBot, services.Auth, services.Restaurant, models.TELEGRAM)
		log.Println("delivery telegram:", err)
		ch <- true
	}()

	go func() {
		err := bot.Start(baleBot, services.Auth, services.Restaurant, models.BALE)
		log.Println("delivery Bale:", err)
		ch <- true
	}()

	for range DELIVERY_OPTIONS {
		<-ch
	}

	return nil
}
