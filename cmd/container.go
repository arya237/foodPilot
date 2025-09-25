package cmd

import (
	"github.com/arya237/foodPilot/internal/handler/auth"
	"github.com/arya237/foodPilot/internal/handler/food"
	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/arya237/foodPilot/internal/repositories/fakedb"
	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/pkg/reservations"
	"github.com/arya237/foodPilot/pkg/reservations/samad"
	"sync"
	"time"
)

type Container struct {
	db *fakedb.FakeDb

	//repositories
	UserRepo repositories.User
	FoodRepo repositories.Food
	RateRepo repositories.Rate

	//service
	UserService    services.UserService
	FoodService    services.FoodService
	RateService    services.RateFoodService
	Samad          reservations.RequiredFunctions
	ReserveService services.Reserve

	mutex sync.RWMutex
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) SetUp(db *fakedb.FakeDb, conf *samad.Config) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.db = db
	c.UserRepo = repositories.NewUserRepo(c.db)
	c.FoodRepo = repositories.NewFoodRepo(c.db)
	c.RateRepo = repositories.NewRateRepo(c.db)

	c.UserService = services.NewUserService(c.UserRepo, conf)
	c.FoodService = services.NewFoodService(c.FoodRepo)
	c.RateService = services.NewRateFoodService(c.RateRepo, c.FoodRepo)
	c.Samad = samad.NewSamad(conf)
	c.ReserveService = services.NewReserveService(c.FoodService, c.UserService, c.RateService, c.Samad)
}

func (c *Container) GetFoodHandler() *food.FoodHandler {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	foodHandler := food.NewFoodHandler(c.RateService, c.FoodService, c.ReserveService)
	return foodHandler
}

func (c *Container) GetLoginHandler() *auth.LoginHandler {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	loginHandler := auth.NewLoginHandler(time.Hour, c.UserService)
	return loginHandler
}
