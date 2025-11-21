package web

import (
	"time"

	"github.com/arya237/foodPilot/internal/web/api/admin"
	"github.com/arya237/foodPilot/internal/web/api/auth"
	"github.com/arya237/foodPilot/internal/web/api/user"
	"github.com/arya237/foodPilot/internal/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title                      FoodPilot
// @description                Auto food reserve
// @termsOfService             http://swagger.io/terms/
// @contact.name               FoodPilot Dev Team
// @contact.url                https://github.com/arya237/foodPilot
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
// @description                Type `Bearer ` followed by your JWT token. example: "Bearer abcde12345"
func Start(tokenEpereTime time.Duration, userService services.UserService,
	adminService services.AdminService, resrveService services.Reserve) error{

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

	authHandler := auth.NewHandler(tokenEpereTime , userService)
	userHandler := user.NewUserHandler(userService)
	adminHandler := admin.New(adminService, resrveService)
	
	authGroup := engine.Group("/auth")
	userGroup := engine.Group("/user")
	adminGroup := engine.Group("/admin")

	auth.RegisterRoutes(authGroup, authHandler)
	user.RegisterRoutes(userGroup, userHandler)
	admin.RegisterRoutes(adminGroup, *adminHandler) // Fuck you arya

	return engine.Run(":8080")
}
