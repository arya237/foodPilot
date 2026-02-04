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

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}

	return nil
}

func HandleUpdate(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	if update.Message.IsCommand() {
		// handleCommand(update.Message, user)
		return
	}
}
