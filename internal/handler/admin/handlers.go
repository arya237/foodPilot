package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetFood      godoc
// @Summary     Get users
// @Description Return all the users
// @Tags        Admin
// @Security    BearerAuth
// @Produce     json
// @Success     200 {object} GetUsersResponse
// @Failure     500 {object} ErrorResponse
// @Router      /admin/users [GET]
func (h *AdminHandler) GetUsers(c *gin.Context){
	users , err := h.UserServise.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, GetUsersResponse{
		Users: users,
	})
}

// GetFood      godoc
// @Summary     Get food
// @Description Return all the foods
// @Tags        Admin
// @Security    BearerAuth
// @Produce     json
// @Success     200 {object} GetFoodsResponse
// @Failure     500 {object} ErrorResponse
// @Router      /admin/foods [GET]
func (h *AdminHandler) GetFood(c *gin.Context){
	foodList, err := h.FoodService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, GetFoodsResponse{
		foodList,
	})
}
