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
	connection := false
	if con, ok := c.Get("connection").(bool); ok {
		connection = con
	}
	
	return c.Send("به فود پایلوت خوش آمدید", mainMenu(connection))
}

func aboutUs(c tele.Context) error {
	return c.Send("ما خیلی خفنیم")
}
