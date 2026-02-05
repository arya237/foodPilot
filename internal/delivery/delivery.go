package delivery

import (
	"log"
	"time"

	"github.com/arya237/foodPilot/internal/delivery/bot"
	"github.com/arya237/foodPilot/internal/delivery/web"
	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/internal/services/auth"
	tele "gopkg.in/telebot.v3"
)

const DELIVERY_OPTIONS = 2

func Start(tokenEpereTime time.Duration, userService services.UserService,
	adminService services.AdminService, resrveService services.Reserve, teleBot *tele.Bot, auth auth.Auth) error {

	ch := make(chan any)

	go func() {
		err := web.Start(tokenEpereTime, userService, adminService, resrveService)
		log.Println(err)
		ch <- true
	}()

	go func() {
		err := bot.Start(teleBot, auth)
		log.Println(err)
		ch <- true
	}()

	for range DELIVERY_OPTIONS {
		<-ch
	}

	return nil
}
