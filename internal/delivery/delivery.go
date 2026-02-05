package delivery

import (
	"log"
	"time"

	"github.com/arya237/foodPilot/internal/delivery/bot"
	"github.com/arya237/foodPilot/internal/delivery/web"
	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/internal/services/auth"
	tele "gopkg.in/telebot.v3"
)

const (
	DELIVERY_OPTIONS = 3
	TOKEN_EXP = 24 * time.Hour 
)

type NeededServises struct {
	User   services.UserService
	Admin  services.AdminService
	Resrve services.Reserve
	Auth   auth.Auth
}

func Start(services *NeededServises, teleBot *tele.Bot, baleBot *tele.Bot) error {

	ch := make(chan any)

	go func() {
		err := web.Start(TOKEN_EXP, services.User, services.Admin, services.Resrve)
		log.Println("delivery web:", err)
		ch <- true
	}()

	go func() {
		err := bot.Start(teleBot, services.Auth, models.TELEGRAM)
		log.Println("delivery telegram:", err)
		ch <- true
	}()

	go func() {
		err := bot.Start(baleBot, services.Auth, models.BALE)
		log.Println("delivery Bale:", err)
		ch <- true
	}()

	for range DELIVERY_OPTIONS {
		<-ch
	}

	return nil
}
