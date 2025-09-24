package food

import (
	"github.com/arya237/foodPilot/internal/auth"
	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/pkg/logger"
	"github.com/gin-gonic/gin"
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

	group.Use(auth.AuthMiddleware())

	group.GET("/list", foodHandler.GetFoods)
	group.POST("/rate", foodHandler.RateFoods)
}
