package admin

import (
	"github.com/arya237/foodPilot/internal/models"
	"github.com/arya237/foodPilot/internal/services"
)

// GetFoodsResponse is the response containing a list of available foods
type GetFoodsResponse struct {
	Foods []*models.Food `json:"foods"`
}

// GetUsersResponse is the response containing a list of registered users
type GetUsersResponse struct {
	Users []*models.User `json:"users"`
}

// ErrorResponse is returned when an error occurs
type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}

// MessageResponse is returned when successful and message occurs
type MessageResponse struct {
	Message string `json:"message" example:"message"`
}

type ReserveFoodResponse struct {
	Results []*services.UserReserveResult `json:"results"`
}

type AddNewFoodRequest struct {
	FoodName string `json:"food"`
}

type DeleteFoodRequest struct {
	FoodID int `json:"foodId"`
}

type AddNewUserRequest struct {
	Username string          `json:"username"`
	Password string          `json:"password"`
	Role     models.UserRole `json:"role"`
}

type AddNewUserResponse struct {
	ID int `json:"ID"`
}

type DeleteUserRequest struct {
	UserID int `json:"userID"`
}

type UpdateUserRequest struct {
	Id       int             `json:"userid"`
	Username string          `json:"username"`
	Password string          `json:"password"`
	Autosave bool            `json:"autosave"`
	Role     models.UserRole `json:"role"`
	Token    string          `json:"token"`
}
