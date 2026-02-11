package cmd

import (
	"github.com/arya237/foodPilot/internal/getways"
	"github.com/arya237/foodPilot/internal/getways/bot"
	tele "gopkg.in/telebot.v3"
)

type getway struct {
	telegramSender getways.Sender
	baleSender     getways.Sender
}

func NewGetway(teleBot *tele.Bot, baleBot *tele.Bot) *getway {
	return &getway{
		telegramSender: bot.NewSender(teleBot),
		baleSender:     bot.NewSender(baleBot),
	}

}
