package main

import (
	"github.com/arya237/foodPilot/cmd"
	"github.com/arya237/foodPilot/internal/config"
	"github.com/arya237/foodPilot/internal/handler/auth"
	"github.com/arya237/foodPilot/internal/handler/food"
	"github.com/arya237/foodPilot/internal/repositories/fakedb"
	"github.com/gin-gonic/gin"
	"log"
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

	authGroup := engine.Group("/auth")
	foodGroup := engine.Group("/food")

	auth.RegisterRoutes(authGroup, authHandlers)
	food.RegisterRoutes(foodGroup, foodHandlers)

	log.Print(engine.Run("localhost:8080"))
}
