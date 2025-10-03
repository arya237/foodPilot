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
func (h *AdminHandler) GetUsers(c *gin.Context) {
	users, err := h.UserServise.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
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
func (h *AdminHandler) GetFood(c *gin.Context) {
	foodList, err := h.FoodService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
	}

	c.JSON(http.StatusOK, GetFoodsResponse{
		foodList,
	})
}

func (h *AdminHandler) AddNewFood(c *gin.Context) {
	var arrived AddNewFoodRequest

	if err := c.BindJSON(&arrived); err != nil {
		c.JSON(http.StatusBadRequest, AddNewFoodResponse{Error: err.Error(), Message: ""})
		return
	}

	message, err := h.FoodService.Save(arrived.FoodName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, AddNewFoodResponse{Error: err.Error(), Message: message})
		return
	}

	c.JSON(http.StatusCreated, AddNewFoodResponse{Error: "", Message: message})
}

func (h *AdminHandler) DeleteFood(c *gin.Context) {
	var arrived DeleteFoodRequest
	if err := c.BindJSON(&arrived); err != nil {
		c.JSON(http.StatusBadRequest, DeleteFoodResponse{Error: err.Error(), Message: ""})
		return
	}

	message, err := h.FoodService.Delete(arrived.FoodID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, DeleteFoodResponse{Error: err.Error(), Message: message})
		return
	}

	c.JSON(http.StatusOK, DeleteFoodResponse{Error: "", Message: message})
}

// ReserveFood  godoc
// @Summary     Reserve food
// @Description Reserve food for all users
// @Tags        Admin
// @Security    BearerAuth
// @Produce     json
// @Success     200 {object} MessageResponse
// @Failure     500 {object} ErrorResponse
// @Router      /admin/reserve [POST]
func (h *AdminHandler) ReserveFood(c *gin.Context) {
	message, err := h.ReserveService.ReserveFood()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, MessageResponse{Message: message})
}
