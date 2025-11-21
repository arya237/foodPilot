package admin

import (
	"time"

	"github.com/arya237/foodPilot/internal/http/middelware"
	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

type AdminHandler struct {
	AdminService   services.AdminService
	ReserveService services.Reserve
	Logger         logger.Logger
}

func New(adminSerivce services.AdminService, reserveService services.Reserve) *AdminHandler {
	return &AdminHandler{
		AdminService:   adminSerivce,
		ReserveService: reserveService,
		Logger:         logger.New("Admin panel logger"),
	}
}

func RegisterRoutes(group *gin.RouterGroup, adminHandler AdminHandler) {

	rate := limiter.Rate{
		Period: 3 * time.Second,
		Limit:  10,
	}

	store := memory.NewStore()
	limiter := limiter.New(store, rate)

	group.Use(middelware.LimitMiddelware(limiter), middelware.AuthMiddleware(), middelware.AdminOnly())
	group.GET("/user", adminHandler.GetUsers)
	group.POST("/user", adminHandler.AddNewUser)
	group.DELETE("/user/:userID", adminHandler.DeleteUser)
	group.PUT("/user", adminHandler.UpdateUser)

	group.GET("/food", adminHandler.GetFood)
	group.POST("/food", adminHandler.AddNewFood)
	group.DELETE("/food/:foodID", adminHandler.DeleteFood)

	group.POST("/reserve", adminHandler.ReserveFood)
}
