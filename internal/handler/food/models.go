package food

import "github.com/arya237/foodPilot/internal/models"

type GetFoodsResponse struct {
	Foods []models.Food `json:"foods"`
}

type RateFoodsRequest struct {
	Foods map[string]int `json:"foods" binding:"required"`
}

type RateFoodsResponse struct {
	Message string `json:"message"`
}

type AutoSaveRequest struct {
	AutoSave bool `json:"autosave" binding:"required"`
}

type AutoSaveResponse struct {
	Message string `json:"message"`
}
