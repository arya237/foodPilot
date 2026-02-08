package bot

import (
	"fmt"

	tele "gopkg.in/telebot.v3"
)

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
	return c.Send("تامام", mainMenu(true))
}
