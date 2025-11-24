package user

import "github.com/arya237/foodPilot/internal/models"

// AutoSaveRequest is the request body for enabling or disabling autosave
type AutoSaveRequest struct {
	AutoSave *bool `json:"autosave" binding:"required"`
}

// AutoSaveResponse is the response returned after updating the autosave setting
type AutoSaveResponse struct {
	Message string `json:"message"`
}

// RatesResponse is the response containing food ratings
type RatesResponse struct {
	Rates map[string]int `json:"rates" example:"foodName1:93,foodName2:74,foodName3:80"`
}

// ErrorResponse is returned when an error occurs
type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}

// GetFoodsResponse is the response containing a list of available foods
type GetFoodsResponse struct {
	Foods []*models.Food `json:"foods"`
}

// RateFoodsRequest is the request body for rating multiple foods
type RateFoodsRequest struct {
	Foods map[string]int `json:"foods" binding:"required" example:"foodName1:93,foodName2:74,foodName3:80"`
}

// RateFoodsResponse is the response returned after successfully rating foods
type RateFoodsResponse struct {
	Message string `json:"message"`
}
