package food

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetFood      godoc
// @Summary     Get foods
// @Description Return all the foods
// @Tags        Food
// @Security    BearerAuth
// @Produce     json
// @Success     200 {object} GetFoodsResponse
// @Failure     500 {object} ErrorResponse
// @Router      /food/list [GET]
func (h *FoodHandler) GetFoods(c *gin.Context) {

	foodList, err := h.FoodService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, GetFoodsResponse{
		foodList,
	})
}

// RateFood     godoc
// @Summary     Rates foods
// @Description Rates all the foods
// @Tags        Food
// @Security    BearerAuth
// @Accept      json
// @Param       rates body RateFoodsRequest true "Rates info"
// @Produce     json
// @Success     200 {object} RateFoodsResponse
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /food/rate [POST]
func (h *FoodHandler) RateFoods(c *gin.Context) {
	rates := RateFoodsRequest{}
	if err := c.ShouldBindJSON(&rates); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request"})
		return
	}

	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "userNotFound"})
		return
	}

	message, err := h.RateService.SaveRate(userID.(string), rates.Foods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, RateFoodsResponse{
		Message: message,
	})
}
