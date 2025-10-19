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
// @Router      /admin/users [GET]

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

//func (h *AdminHandler) UpdateUser(c *gin.Context) {
//	var arrived UpdateUserRequest
//	if err := c.ShouldBindJSON(&arrived); err != nil {
//		c.JSON(http.StatusBadRequest, ErrorResponse{
//			Error: err.Error(),
//		})
//	}
//
//	message, err := h.UserServise.Update(arrived.Username, arrived.Password, arrived.Updated)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, ErrorResponse{
//			Error: err.Error(),
//		})
//	}
//
//	c.JSON(http.StatusAccepted, UpdateUserResponse{
//		Error:   "",
//		Message: message,
//	})
//}

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
// @Router      /admin/reserve [POST]
func (h *AdminHandler) ReserveFood(c *gin.Context) {
	results, err := h.ReserveService.ReserveFood()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ReserveFoodResponse{Results: results})
}
