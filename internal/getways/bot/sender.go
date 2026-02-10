package bot

import (
	"strconv"

	"github.com/arya237/foodPilot/internal/getways"
	tele "gopkg.in/telebot.v3"
)

type botMessenger struct {
	bot *tele.Bot
}

func NewSender(bot *tele.Bot) getways.Sender {
	return &botMessenger{
		bot: bot,
	}
}
func (b *botMessenger) Send(to, message string) error {
	id, err := strconv.Atoi(to)
	if err != nil {
		return err
	}
	_, err = b.bot.Send(tele.ChatID(id), message)
	return err
}
