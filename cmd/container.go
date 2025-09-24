package cmd

import (
	"github.com/arya237/foodPilot/internal/handler/auth"
	"github.com/arya237/foodPilot/internal/handler/food"
	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/arya237/foodPilot/internal/repositories/fakedb"
	"github.com/arya237/foodPilot/internal/services"
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
	UserService services.UserService
	FoodService services.FoodService
	RateService services.RateFoodService

	mutex sync.RWMutex
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) init(db *fakedb.FakeDb) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.db = db
	c.UserRepo = repositories.NewUserRepo(c.db)
	c.FoodRepo = repositories.NewFoodRepo(c.db)
	c.RateRepo = repositories.NewRateRepo(c.db)

	c.UserService = services.NewUserService(c.UserRepo)
	c.FoodService = services.NewFoodService(c.FoodRepo)
	c.RateService = services.NewRateFoodService(c.RateRepo, c.FoodRepo)
}

func (c *Container) GetFoodHandler() *food.FoodHandler {
	c.mutex.RLock()
	defer c.mutex.Unlock()
	foodHandler := food.NewFoodHandler(c.RateService, c.FoodService)
	return foodHandler
}

func (c *Container) GetLoginHandler() *auth.LoginHandler {
	c.mutex.RLock()
	defer c.mutex.Unlock()
	loginHandler := auth.NewLoginHandler(time.Hour, c.UserService)
	return loginHandler
}
