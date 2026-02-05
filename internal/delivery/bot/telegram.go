package bot

import (
	"errors"

	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/services/auth"
	tele "gopkg.in/telebot.v3"
)

var (
	btnAboutUs     = "درباره ما"
	btnAutoReserve = "رزور خودکار"
)

func Start(bot *tele.Bot, auth auth.Auth, provider models.IdProvider) error {

	if bot == nil {
		return errors.New("Bot is nil")
	}
	bot.Use(AuthMiddleware(auth, provider))

	bot.Handle("/start", onStart)
	bot.Handle(tele.OnText, others)
	bot.Handle(btnAboutUs, aboutUs)
	bot.Handle(btnAutoReserve, autoRserve)

	bot.Start()
	return nil
}
