package admin

import (
	"fmt"
	"net/http"
	"strconv"

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
// @Router      /api/admin/user [GET]
func (h *AdminHandler) GetUsers(c *gin.Context) {
	users, err := h.AdminService.GetUsers()
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
// @Summary     Add new user
// @Description Register new user
// @Tags        Admin
// @Security    BearerAuth
// @Accept      json
// @Param       newUser body AddNewUserRequest true "User info"
// @Produce     json
// @Success     201 {object} AddNewUserResponse
// @Failure     500 {object} ErrorResponse
// @Failure     400 {object} ErrorResponse
// @Router      /api/admin/user [POST]
func (h *AdminHandler) AddNewUser(c *gin.Context) {
	var arrived AddNewUserRequest
	if err := c.ShouldBindJSON(&arrived); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	userID, err := h.AdminService.AddUser(arrived.Username, arrived.Password, arrived.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, AddNewUserResponse{
		ID: userID,
	})
}

// GetFood      godoc
// @Summary     Delete user
// @Description Delete user
// @Tags        Admin
// @Param       userID path int true "User ID"
// @Security    BearerAuth
// @Produce     json
// @Success     200 {object} MessageResponse
// @Failure     500 {object} ErrorResponse
// @Failure     400 {object} ErrorResponse
// @Router      /api/admin/user/{userID} [DELETE]
func (h *AdminHandler) DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	err = h.AdminService.DeleteUser(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, MessageResponse{
		Message: "user deleted",
	})
}

// GetFood      godoc
// @Summary     Update user
// @Description Update user
// @Tags        Admin
// @Security    BearerAuth
// @Accept      json
// @Param       userInfo body UpdateUserRequest true "user info"
// @Produce     json
// @Success     202 {object} GetFoodsResponse
// @Failure     500 {object} ErrorResponse
// @Router      /api/admin/user [PUT]
func (h *AdminHandler) UpdateUser(c *gin.Context) {
	var arrived UpdateUserRequest
	if err := c.ShouldBindJSON(&arrived); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	err := h.AdminService.UpdateUser(arrived.Id, arrived.Username, arrived.Password, arrived.Autosave,arrived.Role, arrived.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, MessageResponse{
		Message: "User updated successfully",
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
// @Router      /api/admin/food [GET]
func (h *AdminHandler) GetFood(c *gin.Context) {
	foodList, err := h.AdminService.GetFoods()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, GetFoodsResponse{
		foodList,
	})
}

// GetFood      godoc
// @Summary     Add new food
// @Description Register new food
// @Tags        Admin
// @Security    BearerAuth
// @Accept      json
// @Param       newFood body AddNewFoodRequest true "Food info"
// @Produce     json
// @Success     201 {object} MessageResponse
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /api/admin/food [POST]
func (h *AdminHandler) AddNewFood(c *gin.Context) {
	var arrived AddNewFoodRequest

	if err := c.BindJSON(&arrived); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	id, err := h.AdminService.AddFood(arrived.FoodName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, MessageResponse{
		Message: fmt.Sprintf("food with id %v created", id),
	})
}

// GetFood      godoc
// @Summary     Delete food
// @Description Delete food
// @Tags        Admin
// @Param       foodID path int true "Food ID"
// @Security    BearerAuth
// @Produce     json
// @Success     200 {object} MessageResponse
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /api/admin/food/{foodID} [DELETE]
func (h *AdminHandler) DeleteFood(c *gin.Context) {
	foodID, err := strconv.Atoi(c.Param("foodID"))
	if err != nil {
		c.JSON(http.StatusBadRequest,  ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	err = h.AdminService.DeleteFood(foodID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, MessageResponse{
		Message: "food deleted",
	})
}

// ReserveFood  godoc
// @Summary     Reserve food
// @Description Reserve food for all users
// @Tags        Admin
// @Security    BearerAuth
// @Produce     json
// @Success     200 {object} ReserveFoodResponse
// @Failure     500 {object} ErrorResponse
// @Router      /api/admin/reserve [POST]
func (h *AdminHandler) ReserveFood(c *gin.Context) {
	results, err := h.ReserveService.ReserveFood()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ReserveFoodResponse{Results: results})
}
