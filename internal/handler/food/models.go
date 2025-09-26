package food

import "github.com/arya237/foodPilot/internal/models"

// GetFoodsResponse is the response containing a list of available foods
type GetFoodsResponse struct {
	Foods []*models.Food `json:"foods"`
}

// RateFoodsRequest is the request body for rating multiple foods
type RateFoodsRequest struct {
	Foods map[string]int `json:"foods" binding:"required"`
}

// RateFoodsResponse is the response returned after successfully rating foods
type RateFoodsResponse struct {
	Message string `json:"message"`
}

// ErrorResponse is returned when an error occurs
type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}

// MessageResponse is returned when successful and message occurs
type MessageResponse struct {
	Message string `json:"message" example:"message"`
}