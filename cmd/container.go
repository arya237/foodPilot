package cmd

import (
	"database/sql"
	"sync"
	"time"

	_ "github.com/arya237/foodPilot/docs"
	"github.com/arya237/foodPilot/internal/config"
	db_postgres "github.com/arya237/foodPilot/internal/db/postgres"
	"github.com/arya237/foodPilot/internal/repositories"

	//"github.com/arya237/foodPilot/internal/repositories/memory"
	repo_postgres "github.com/arya237/foodPilot/internal/repositories/postgres"
	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/internal/web"
	"github.com/arya237/foodPilot/pkg/reservations"
	"github.com/arya237/foodPilot/pkg/reservations/samad"
)

type Container struct {
	//db *tempdb.FakeDb
	db *sql.DB
	//repositories
	UserRepo repositories.User
	FoodRepo repositories.Food
	RateRepo repositories.Rate

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
	//c.UserRepo = memory.NewUserRepo(c.db)
	//c.FoodRepo = memory.NewFoodRepo(c.db)
	//c.RateRepo = memory.NewRateRepo(c.db)

	c.UserRepo = repo_postgres.NewUserRepo(c.db)
	c.FoodRepo = repo_postgres.NewFoodRepo(c.db)
	c.RateRepo = repo_postgres.NewRateRepo(c.db)

	c.UserService = services.NewUserService(c.UserRepo, c.FoodRepo, c.RateRepo, conf)

	c.AdminService = services.NewAdminService(c.UserRepo, c.FoodRepo)

	c.Samad = samad.NewSamad(conf)
	c.ReserveService = services.NewReserveService(c.UserRepo, c.UserService, c.Samad)
}

func Run() error {

	conf, err := config.New()
	if err != nil {
		return err
	}

	//db := tempdb.NewDb(conf.DBConfig)
	db := db_postgres.NewDB(conf.PostGresConfig)
	container := NewContainer()
	//container.SetUp(db, conf.SamadConfig)
	container.SetUp(db, conf.SamadConfig)
	return web.Start(time.Hour, container.UserService,
		container.AdminService, container.ReserveService)
}
