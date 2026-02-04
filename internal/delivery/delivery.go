package delivery

import (
	"log"
	"time"

	"github.com/arya237/foodPilot/internal/delivery/telegram"
	"github.com/arya237/foodPilot/internal/delivery/web"
	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/internal/services/auth"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const DELIVERY_OPTIONS = 2

func Start(tokenEpereTime time.Duration, userService services.UserService,
	adminService services.AdminService, resrveService services.Reserve, bot *tgbotapi.BotAPI, auth auth.Auth) error {

	ch := make(chan any)

	go func() {
		err := web.Start(tokenEpereTime, userService, adminService, resrveService)
		log.Println(err)
		ch <- true
	}()

	go func() {
		err := telegram.Start(bot, auth)
		log.Println(err)
		ch <- true
	}()

	for range DELIVERY_OPTIONS {
		<-ch
	}

	return nil
}
