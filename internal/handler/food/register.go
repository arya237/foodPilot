package food

import (
	"github.com/arya237/foodPilot/internal/auth"
	"github.com/gin-gonic/gin"
)

type FoodHandler struct {
	//...
}

func NewFoodHandler() *FoodHandler {
	return &FoodHandler{}
}

func RegisterRoutes(group *gin.RouterGroup) {
	h := NewFoodHandler()
	group.Use(auth.AuthMiddleware())
	
	group.GET("/list", h.GetFoods)
	group.POST("/rate", h.RateFoods)
	group.POST("/autosave", h.AutoSave)
}
