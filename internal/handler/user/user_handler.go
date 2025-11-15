package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AutoSave     godoc
// @Summary     Toggle user autosave
// @Description Toggle user autosave attribute
// @Tags        User
// @Accept      json
// @Param user  body AutoSaveRequest true "Toggle info"
// @Security    BearerAuth
// @Success     200 {object} AutoSaveResponse
// @Failure     404 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /user/autosave [POST]
func (h *UserHandler) AutoSave(c *gin.Context) {

	req := AutoSaveRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request"})
		return
	}

	id, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "user not found"})
		return
	}
	userID, _ := strconv.Atoi(id.(string))

	err := h.UserService.ToggleAutoSave(userID, *req.AutoSave)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "server error"})
		return
	}

	c.JSON(http.StatusOK, AutoSaveResponse{
		Message: "Auto save updated",
	})
}

// GetRates     godoc
// @Summary     return user rates
// @Description return user rates
// @Tags        User
// @Security    BearerAuth
// @Success     200 {object} RatesResponse
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /user/rates [GET]
func (h *UserHandler) GetRates(c *gin.Context) {
	id, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "user not found"})
		return
	}

	userID, _ := strconv.Atoi(id.(string))

	userRates, err := h.UserService.ViewRating(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, RatesResponse{Rates: userRates})
}

// GetFood      godoc
// @Summary     Get foods
// @Description Return all the foods
// @Tags        User
// @Security    BearerAuth
// @Produce     json
// @Success     200 {object} GetFoodsResponse
// @Failure     500 {object} ErrorResponse
// @Router      /user/view-foods [GET]
func (h *UserHandler) GetFoods(c *gin.Context) {

	foodList, err := h.UserService.ViewFoods()
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
// @Tags        User
// @Security    BearerAuth
// @Accept      json
// @Param       rates body RateFoodsRequest true "Rates info"
// @Produce     json
// @Success     200 {object} RateFoodsResponse
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /user/rate [POST]
func (h *UserHandler) RateFoods(c *gin.Context) {
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

	message, err := h.UserService.RateFoods(userID.(string), rates.Foods)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, RateFoodsResponse{
		Message: message,
	})
}