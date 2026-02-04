package telegram

import (
	"errors"

	"github.com/arya237/foodPilot/internal/services/auth"
	tele "gopkg.in/telebot.v3"
)

type Handler struct {
	bot  *tele.Bot
	auth auth.Auth
}

func Start(bot *tele.Bot, auth auth.Auth) error {

	if bot == nil {
		return errors.New("Bot is nil")
	}
	

	bot.Handle(tele.OnText, func(c tele.Context) error  {
		return c.Send("meow")
	})

	bot.Start()
	return nil
}


