package bot

import (
	"sync"

	"github.com/arya237/foodPilot/internal/services/restaurant"
	tele "gopkg.in/telebot.v3"
)

type handler struct {
	restaurant restaurant.Connector
	mu         sync.RWMutex
	cache      map[int64]userState
}

func newHandler(restaurant restaurant.Connector) *handler {
	return &handler{
		cache:      make(map[int64]userState),
		restaurant: restaurant,
	}
}

func onStart(c tele.Context) error {
	keyboard := &tele.ReplyMarkup{
		ResizeKeyboard: true,
	}

	btnAutoReserve := keyboard.Text(btnAutoReserve)
	btnAboutUs := keyboard.Text(btnAboutUs)

	keyboard.Reply(
		keyboard.Row(btnAutoReserve),
		keyboard.Row(btnAboutUs),
	)
	return c.Send("به فود پایلوت خوش آمدید", keyboard)
}

func aboutUs(c tele.Context) error {
	return c.Send("ما خیلی خفنیم")
}
