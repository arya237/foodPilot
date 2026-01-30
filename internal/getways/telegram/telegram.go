package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Config struct {
	API   string
	Token string
}

func New(cfg *Config) (bot *tgbotapi.BotAPI,err error) {
	if cfg.API == "" {
		bot, err = tgbotapi.NewBotAPI(cfg.Token)
	} else {
		bot, err = tgbotapi.NewBotAPIWithAPIEndpoint(cfg.Token, cfg.API)
	}

	if err != nil {
		return nil, err
	}
	
	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot, err
}
