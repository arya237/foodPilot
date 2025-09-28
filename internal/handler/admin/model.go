package admin

import "github.com/arya237/foodPilot/internal/models"

type GetFoodsResponse struct {
	Foods []*models.Food `json:"foods"`
}

type GetUsersResponse struct {
	Users []*models.User `json:"users"`
}