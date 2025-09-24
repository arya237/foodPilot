package food

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *FoodHandler) GetFoods(c *gin.Context) {

	foodList, err := h.FoodService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, GetFoodsResponse{
		foodList,
	})
}

func (h *FoodHandler) RateFoods(c *gin.Context) {
	rates := RateFoodsRequest{}
	if err := c.ShouldBindJSON(&rates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userNotFound"})
		return
	}

	message, err := h.RateService.SaveRate(userID.(string), rates.Foods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, RateFoodsResponse{
		Message: message,
	})
}
