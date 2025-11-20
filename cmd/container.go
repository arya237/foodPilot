package cmd

import (
	"sync"
	"time"

	_ "github.com/arya237/foodPilot/docs"
	"github.com/arya237/foodPilot/internal/config"
	"github.com/arya237/foodPilot/internal/db/tempdb"
	"github.com/arya237/foodPilot/internal/handler/admin"
	"github.com/arya237/foodPilot/internal/handler/auth"
	"github.com/arya237/foodPilot/internal/handler/user"
	"github.com/arya237/foodPilot/internal/repositories"
	"github.com/arya237/foodPilot/internal/repositories/memory"
	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/pkg/reservations"
	"github.com/arya237/foodPilot/pkg/reservations/samad"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Container struct {
	db *tempdb.FakeDb

	//repositories
	UserRepo repositories.User
	FoodRepo repositories.Food
	RateRepo repositories.Rate

	//service
	UserService    services.UserService
	AdminService   services.AdminService
	Samad          reservations.RequiredFunctions
	ReserveService services.Reserve

	mutex sync.RWMutex
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) SetUp(db *tempdb.FakeDb, conf *samad.Config) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.db = db
	c.UserRepo = memory.NewUserRepo(c.db)
	c.FoodRepo = memory.NewFoodRepo(c.db)
	c.RateRepo = memory.NewRateRepo(c.db)

	c.UserService = services.NewUserService(c.UserRepo, c.FoodRepo, c.RateRepo, conf)

	c.AdminService = services.NewAdminService(c.UserRepo, c.FoodRepo)

	c.Samad = samad.NewSamad(conf)
	c.ReserveService = services.NewReserveService(c.UserRepo, c.UserService, c.Samad)
}


func (c *Container) GetLoginHandler() *auth.AuthHandler {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	loginHandler := auth.NewHandler(time.Hour, c.UserService)
	return loginHandler
}

func (c *Container) GetUserHandler() *user.UserHandler {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	userHandler := user.NewUserHandler(c.UserService)
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
func Run() (error) {
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
		return err
	}

	db := tempdb.NewDb(conf.DBConfig)
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

	return engine.Run(":8080")
}
