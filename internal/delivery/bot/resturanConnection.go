package bot

import (
	tele "gopkg.in/telebot.v3"
)

func (h *handler) onRestaurantLogin(c tele.Context) error {
	h.mu.Lock()
	h.cache[c.Chat().ID] = userState{
		userID: c.Chat().ID,
		state:  waitingForUsername,
	}
	h.mu.Unlock()

	return c.Send("لطفا نام کاربری خود در سماد را وارد کنید")
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

	return c.Send("لطفا پسورد خود در سماد را وارد کنید.")
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
		return c.Send("اتصال برقرار نشد! این خطا می تواند حاصل اختلالات شبکه یا رمز نا معتبر باشد.")
	}

	h.cache[c.Chat().ID] = userState{
		userID: c.Chat().ID,
		state:  idel,
	}
	return c.Send("اتصال برقرار شد. از منوی اصلی در قسمت تنظیمات  رزور خودکار می توانید اقدام به مدیریت حساب خود کنید.", mainMenu(true))
}
