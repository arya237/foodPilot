package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func Start(bot *tgbotapi.BotAPI) error{
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}

	return nil
}