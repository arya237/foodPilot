package api

import (
	"time"

	adminUscase "github.com/arya237/foodPilot/internal/services/admin"
	"github.com/arya237/foodPilot/internal/delivery/web/api/admin"
	"github.com/arya237/foodPilot/internal/delivery/web/api/auth"
	"github.com/arya237/foodPilot/internal/delivery/web/api/user"
	"github.com/arya237/foodPilot/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(group *gin.RouterGroup, tokenEpereTime time.Duration,
	userService services.UserService, adminService services.AdminService,
	resrveService services.Reserve, notifier adminUscase.Notifier) *gin.RouterGroup {

	authHandler := auth.NewHandler(tokenEpereTime, userService)
	userHandler := user.NewUserHandler(userService)
	adminHandler := admin.New(adminService, resrveService, notifier)

	authGroup := group.Group("/auth")
	userGroup := group.Group("/user")
	adminGroup := group.Group("/admin")

	auth.RegisterRoutes(authGroup, authHandler)
	user.RegisterRoutes(userGroup, userHandler)
	admin.RegisterRoutes(adminGroup, *adminHandler) // Fuck you arya

	return group
}
