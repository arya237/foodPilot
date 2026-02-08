package bot

import (
	"errors"

	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/services/auth"
	"github.com/arya237/foodPilot/internal/services/restaurant"
	tele "gopkg.in/telebot.v3"
)

func Start(bot *tele.Bot, auth auth.Auth, restaurant restaurant.Connector, provider models.IdProvider) error {

	if bot == nil {
		return errors.New("Bot is nil")
	}
	bot.Use(AuthMiddleware(auth, provider))

	h := newHandler(restaurant)
	bot.Handle("/start", onStart)
	bot.Handle(btnAboutUs, aboutUs)
	bot.Handle(btnAutoReserve, h.onRestaurantLogin)
	bot.Handle(tele.OnText, h.onText)
	bot.Handle(btnRestaurantSetting, onRestaurantSetting)

	bot.Start()
	return nil
}

func (h *handler) onText(c tele.Context) error {
	var value *userState
	h.mu.RLock()
	if val, ok := h.cache[c.Chat().ID]; ok {
		value = &val
	}
	h.mu.RUnlock()

	if value == nil {
		return others(c)
	}

	switch value.state {
	case waitingForUsername:
		return h.onUsername(c)
	case waitingForPassword:
		return h.onPassword(c)
	}

	return others(c)
}
