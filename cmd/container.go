package cmd

import (
	"database/sql"
	"log"
	"sync"
	"time"

	_ "github.com/arya237/foodPilot/docs"
	"github.com/arya237/foodPilot/internal/config"
	db_postgres "github.com/arya237/foodPilot/internal/db/postgres"
	"github.com/arya237/foodPilot/internal/delivery"
	"github.com/arya237/foodPilot/internal/getways/telegram"
	"github.com/arya237/foodPilot/internal/services/auth"
	tele "gopkg.in/telebot.v3"

	"github.com/arya237/foodPilot/internal/getways/reservations"
	"github.com/arya237/foodPilot/internal/getways/reservations/samad"
	"github.com/arya237/foodPilot/internal/services"
)

type Container struct {
	//service
	UserService    services.UserService
	AdminService   services.AdminService
	ReserveService services.Reserve
	AuthService    auth.Auth

	//getways
	Samad reservations.ReserveFunctions

	mutex sync.RWMutex
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) SetUp(db *sql.DB, conf *samad.Config) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	repo := NewRepo(db)

	c.UserService = services.NewUserService(repo.User, repo.Food, repo.Rate, repo.Cred, conf)
	c.AdminService = services.NewAdminService(repo.User, repo.Food)
	c.AuthService = auth.New(repo.identities, repo.User)

	c.Samad = samad.NewSamad(conf)
	c.ReserveService = services.NewReserveService(repo.User, repo.Cred, c.UserService, c.Samad)
}

func Run() error {

	conf, err := config.New()
	if err != nil {
		return err
	}

	db := db_postgres.NewDB(conf.PostGresConfig)
	if db == nil {
		log.Println("db is nil ...")
	}
	container := NewContainer()

	container.SetUp(db, conf.SamadConfig)

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

	return delivery.Start(time.Hour, container.UserService,
		container.AdminService, container.ReserveService, bot, container.AuthService)
}
