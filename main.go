package main

import (
	"log"

	"github.com/arya237/foodPilot/cmd"
	"github.com/arya237/foodPilot/internal/config"
	"github.com/arya237/foodPilot/internal/handler/auth"
	"github.com/arya237/foodPilot/internal/handler/food"
	"github.com/arya237/foodPilot/internal/handler/user"
	"github.com/arya237/foodPilot/internal/repositories/fakedb"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()
	db := fakedb.NewDb()
	conf, err := config.New()
	if err != nil {
		log.Print(err.Error())
	}

	container := cmd.NewContainer()
	container.SetUp(db, conf.SamadConfig)

	foodHandlers := container.GetFoodHandler()
	authHandlers := container.GetLoginHandler()
	userHandler := container.GetUserHandler()

	authGroup := engine.Group("/auth")
	foodGroup := engine.Group("/food")
	userGroup := engine.Group("/user")

	auth.RegisterRoutes(authGroup, authHandlers)
	food.RegisterRoutes(foodGroup, foodHandlers)
	user.RegisterRoutes(userGroup, userHandler)

	log.Print(engine.Run("localhost:8080"))
}
