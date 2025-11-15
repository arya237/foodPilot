package cmd

import (
	"sync"
	"time"

	_ "github.com/arya237/foodPilot/docs"
	"github.com/arya237/foodPilot/internal/config"
	"github.com/arya237/foodPilot/internal/db"
	"github.com/arya237/foodPilot/internal/handler/admin"
	"github.com/arya237/foodPilot/internal/handler/auth"
	"github.com/arya237/foodPilot/internal/handler/user"
	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/pkg/reservations"
	"github.com/arya237/foodPilot/pkg/reservations/samad"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Container struct {
	db *db.FakeDb

	//repositories
	UserRepo repositories.User
	FoodRepo repositories.Food
	RateRepo repositories.Rate

	//service
	UserService    services.UserService
	RateService    services.RateFoodService
	AdminService   services.AdminService
	Samad          reservations.RequiredFunctions
	ReserveService services.Reserve

	mutex sync.RWMutex
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) SetUp(db *db.FakeDb, conf *samad.Config) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.db = db
	c.UserRepo = repositories.NewUserRepo(c.db)
	c.FoodRepo = repositories.NewFoodRepo(c.db)
	c.RateRepo = repositories.NewRateRepo(c.db)

	c.UserService = services.NewUserService(c.UserRepo, c.FoodRepo, c.RateRepo, conf)
	c.RateService = services.NewRateFoodService(c.RateRepo, c.FoodRepo)
	c.AdminService = services.NewAdminService(c.UserRepo, c.FoodRepo)

	c.Samad = samad.NewSamad(conf)
	c.ReserveService = services.NewReserveService(c.UserRepo, c.RateService, c.Samad)
}


func (c *Container) GetLoginHandler() *auth.LoginHandler {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	loginHandler := auth.NewLoginHandler(time.Hour, c.UserService)
	return loginHandler
}

func (c *Container) GetUserHandler() *user.UserHandler {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	userHandler := user.NewUserHandler(c.UserService, c.RateService)
	return userHandler
}

func (c *Container) GetAdminHandler() *admin.AdminHandler {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	handler := admin.New(c.AdminService, c.ReserveService)
	return handler
}

// @title                      FoodPilot
// @description                Auto food reserve
// @termsOfService             http://swagger.io/terms/
// @contact.name               FoodPilot Dev Team
// @contact.url                https://github.com/arya237/foodPilot
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
// @description                Type `Bearer ` followed by your JWT token. example: "Bearer abcde12345"
func NewApp() (*gin.Engine, error) {
	engine := gin.Default()
	swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.DocExpansion("none"),
	)

	corsConfig := cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}

	engine.Use(cors.New(corsConfig))

	engine.GET("/swagger/*any", swaggerHandler)

	conf, err := config.New()
	if err != nil {
		return nil, err
	}

	db := db.NewDb(conf.DBConfig)
	container := NewContainer()
	container.SetUp(db, conf.SamadConfig)

	authHandlers := container.GetLoginHandler()
	userHandler := container.GetUserHandler()
	adminHandler := container.GetAdminHandler()

	authGroup := engine.Group("/auth")
	userGroup := engine.Group("/user")
	adminGroup := engine.Group("/admin")

	auth.RegisterRoutes(authGroup, authHandlers)
	user.RegisterRoutes(userGroup, userHandler)
	admin.RegisterRoutes(adminGroup, *adminHandler)

	return engine, nil
}
