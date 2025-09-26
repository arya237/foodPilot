package food

import (
	"github.com/arya237/foodPilot/internal/auth"
	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"time"
)

type FoodHandler struct {
	RateService    services.RateFoodService
	FoodService    services.FoodService
	ReserveService services.Reserve
	logger         logger.Logger
}

func NewFoodHandler(r services.RateFoodService, f services.FoodService, reserve services.Reserve) *FoodHandler {
	return &FoodHandler{
		logger:         logger.New("foodHandler"),
		FoodService:    f,
		RateService:    r,
		ReserveService: reserve,
	}
}

func RegisterRoutes(group *gin.RouterGroup, foodHandler *FoodHandler) {

	rate := limiter.Rate{
		Period: 3 * time.Second,
		Limit:  2,
	}

	store := memory.NewStore()
	limiter := limiter.New(store, rate)

	group.Use(auth.LimitMiddelware(limiter))

	group.GET("/list", auth.AuthMiddleware(), foodHandler.GetFoods)
	group.POST("/rate", auth.AuthMiddleware(), foodHandler.RateFoods)
	group.POST("/reserve", foodHandler.reserveFood)
}
