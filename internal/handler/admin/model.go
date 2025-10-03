package admin

import "github.com/arya237/foodPilot/internal/models"

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