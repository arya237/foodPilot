package bot

import (
	"errors"
	"fmt"
	"sync"

	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/services/auth"
	"github.com/arya237/foodPilot/internal/services/restaurant"
	tele "gopkg.in/telebot.v3"
)

var (
	btnAboutUs     = "درباره ما"
	btnAutoReserve = "رزور خودکار"
)

type state int

const (
	startOfStates state = iota
	idel
	startResturantLogin
	waitingForUsername
	waitingForPassword
	endOfState
)

type userState struct {
	userID   int64
	state    state
	username string
}

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
		return h.onWrong(c)
	}

	switch value.state {
	case waitingForUsername:
		return h.onUsername(c)
	case waitingForPassword:
		return h.onPassword(c)
	}

	return h.onWrong(c)
}
func (h *handler) onRestaurantLogin(c tele.Context) error {
	h.mu.Lock()
	h.cache[c.Chat().ID] = userState{
		userID: c.Chat().ID,
		state:  waitingForUsername,
	}
	h.mu.Unlock()

	return c.Send("نام کاربری بده")
}

func (h *handler) onUsername(c tele.Context) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	username := c.Message().Text

	h.cache[c.Chat().ID] = userState{
		userID:   c.Chat().ID,
		state:    waitingForPassword,
		username: username,
	}

	return c.Send("پسورد بده")
}

func (h *handler) onPassword(c tele.Context) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	id, ok := c.Get("id").(int)
	if !ok {
		c.Send("توکنت گم شده برو بخ توسعه دهنده ها بگو")
	}
	err := h.restaurant.Connect(id, h.cache[c.Chat().ID].username, c.Message().Text)
	if err != nil {
		str := fmt.Sprintf("[%s],[%s] -> [%s]", h.cache[c.Chat().ID].username, c.Message().Text, err.Error())
		return c.Send(str)
	}
	
	h.cache[c.Chat().ID] = userState{
		userID: c.Chat().ID,
		state:  idel,
	}
	return c.Send("تامام")
}

func (h *handler) onWrong(c tele.Context) error {

	h.mu.RLock()
	defer h.mu.RUnlock()

	val, ok := h.cache[c.Chat().ID]
	if !ok {
		return c.Send("چیزی نداریم")
	}

	return c.Send(fmt.Sprintf("%d", val.state))
}
