package cmd

import (
	"log"
	"time"

	_ "github.com/arya237/foodPilot/docs"
	"github.com/arya237/foodPilot/internal/config"
	db_postgres "github.com/arya237/foodPilot/internal/db/postgres"
	"github.com/arya237/foodPilot/internal/delivery"
	"github.com/arya237/foodPilot/internal/getways/telegram"
	tele "gopkg.in/telebot.v3"

	"github.com/arya237/foodPilot/internal/getways/reservations/samad"
)

func Run() error {

	conf, err := config.New()
	if err != nil {
		return err
	}

	db := db_postgres.NewDB(conf.PostGresConfig)
	if db == nil {
		log.Fatal("db is nil ...")
	}

	samd := samad.NewSamad(conf.SamadConfig)
	repo := NewRepo(db)

	services := NewService(repo, samd)

	connectionTries := 5
	var bot *tele.Bot
	for i := range connectionTries {
		getBot, err := telegram.New(conf.TelegramBot)
		if err == nil {
			bot = getBot
			break
		}
		log.Printf("Try[%d]:%s\n", i, err.Error())
	}

	return delivery.Start(time.Hour, services.User, services.Admin, services.Reserve, bot, services.Auth)
}
