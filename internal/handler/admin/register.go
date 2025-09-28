package admin

import (
	"time"

	"github.com/arya237/foodPilot/internal/auth"
	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

type AdminHandler struct {
	UserServise services.UserService
	FoodService    services.FoodService
	Logger         logger.Logger
}

func New(userServise services.UserService, foodService services.FoodService) *AdminHandler{
	return &AdminHandler{
		UserServise: userServise,
		FoodService: foodService,
		Logger: logger.New("Admin panel logger"),
	}
}

func RegisterRoutes(group *gin.RouterGroup, adminHandler AdminHandler) {

	rate := limiter.Rate{
		Period: 3 * time.Second,
		Limit:  10,
	}
	
	store := memory.NewStore()
	limiter := limiter.New(store, rate)

	group.Use(auth.LimitMiddelware(limiter))

	
}