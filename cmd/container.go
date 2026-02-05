package cmd

import (
	"log"

	_ "github.com/arya237/foodPilot/docs"
	"github.com/arya237/foodPilot/internal/config"
	db_postgres "github.com/arya237/foodPilot/internal/db/postgres"
	"github.com/arya237/foodPilot/internal/delivery"
	"github.com/arya237/foodPilot/internal/getways/bot"
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
	teleBot := CreateBot(connectionTries, conf.TelegramBot)
	

	return delivery.Start(&delivery.NeededServises{
		User:   services.User,
		Admin:  services.Admin,
		Resrve: services.Reserve,
		Auth:   services.Auth,
	}, teleBot)
}

func CreateBot(connectionTries int, cfg *bot.Config)(*tele.Bot) {
	for i := range connectionTries {
		getBot, err := bot.New(cfg)
		if err == nil {
			return getBot
		}
		log.Printf("Try[%d]:%s\n", i, err.Error())
	}
	return nil
}