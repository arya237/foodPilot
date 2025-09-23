package food

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *FoodHandler) GetFoods(c *gin.Context) {
	//TODO:
	c.JSON(http.StatusOK, GetFoodsResponse{
		Foods: nil,
	})
}

func (h *FoodHandler) RateFoods(c *gin.Context) {
	rates := RateFoodsRequest{}
	if err := c.ShouldBindJSON(&rates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// if err := h.service.RateFoods(rates); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, RateFoodsResponse{
		Message: "ratings saved",
	})
}

func (h *FoodHandler) AutoSave(c *gin.Context) {

	req := AutoSaveRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	//h.service.ToggleAutoSave(body.AutoSave)

	c.JSON(http.StatusOK, AutoSaveResponse{
		Message: "Auto save updateed",
	})
}
