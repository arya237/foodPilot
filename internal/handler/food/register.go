package food

import (
	"time"

	"github.com/arya237/foodPilot/internal/auth"
	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

type FoodHandler struct {
	RateService services.RateFoodService
	FoodService services.FoodService
	logger      logger.Logger
}

func NewFoodHandler(r services.RateFoodService, f services.FoodService) *FoodHandler {
	return &FoodHandler{
		logger:      logger.New("foodHandler"),
		FoodService: f,
		RateService: r,
	}
}

func RegisterRoutes(group *gin.RouterGroup, foodHandler *FoodHandler) {

	rate := limiter.Rate{
		Period: 3 * time.Second,
		Limit:  2,
	}

	store := memory.NewStore()
	limiter := limiter.New(store, rate)

	group.Use(auth.LimitMiddelware(limiter), auth.AuthMiddleware())
	group.GET("/list", foodHandler.GetFoods)
	group.POST("/rate", foodHandler.RateFoods)
}
