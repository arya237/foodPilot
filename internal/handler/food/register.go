package food

import (
	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/pkg/logger"
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

//func RegisterRoutes(group *gin.RouterGroup) {
//	h := NewFoodHandler()
//	group.Use(auth.AuthMiddleware())
//
//	group.GET("/list", h.GetFoods)
//	group.POST("/rate", h.RateFoods)
//	group.POST("/autosave", h.AutoSave)
//}
