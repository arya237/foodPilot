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
	"github.com/arya237/foodPilot/internal/repositories"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	repo_postgres "github.com/arya237/foodPilot/internal/repositories/postgres"
	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/internal/getways/reservations"
	"github.com/arya237/foodPilot/internal/getways/reservations/samad"
)

type Container struct {
	
	db *sql.DB
	//repositories
	UserRepo repositories.User
	FoodRepo repositories.Food
	RateRepo repositories.Rate
	CredRepo repositories.RestaurantCredentials

	//service
	UserService    services.UserService
	AdminService   services.AdminService
	Samad          reservations.ReserveFunctions
	ReserveService services.Reserve

	mutex sync.RWMutex
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) SetUp(db *sql.DB, conf *samad.Config) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.db = db

	c.UserRepo = repo_postgres.NewUserRepo(c.db)
	c.FoodRepo = repo_postgres.NewFoodRepo(c.db)
	c.RateRepo = repo_postgres.NewRateRepo(c.db)
	c.CredRepo = repo_postgres.NewResturantCred(c.db)

	c.UserService = services.NewUserService(c.UserRepo, c.FoodRepo, c.RateRepo, c.CredRepo, conf)

	c.AdminService = services.NewAdminService(c.UserRepo, c.FoodRepo)

	c.Samad = samad.NewSamad(conf)
	c.ReserveService = services.NewReserveService(c.UserRepo, c.CredRepo, c.UserService, c.Samad)
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
	var bot *tgbotapi.BotAPI
	for i := range connectionTries {
		getBot, err := telegram.New(conf.TelegramBot)
		if err == nil {
			bot = getBot
			break
		}
		log.Printf("Try[%d]:%s\n", i, err.Error())
	}

	return delivery.Start(time.Hour, container.UserService,
		container.AdminService, container.ReserveService, bot)
}
