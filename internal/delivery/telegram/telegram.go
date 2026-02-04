package telegram

import (
	"errors"

	"github.com/arya237/foodPilot/internal/services/auth"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	bot  *tgbotapi.BotAPI
	auth auth.Auth
}

func Start(bot *tgbotapi.BotAPI, auth auth.Auth) error {
	if bot == nil {
		return errors.New("Bot is nil")
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	h := &Handler{
		bot:  bot,
		auth: auth,
	}

	for update := range updates {
		h.HandleUpdate(update)
	}

	return nil
}

func (h *Handler) HandleUpdate(update tgbotapi.Update) {
	// if update.Message == nil {
	// 	return
	// }
	// // id := fmt.Sprintf("%d", update.Message.Chat.ID)

	// // user, err := h.auth.Login(models.TELEGRAM, id)

	// if update.Message.IsCommand() {
	// 	// handleCommand(update.Message, user)
	// 	return
	// }
}
