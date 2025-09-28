package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *AdminHandler) GetUsers(c *gin.Context){
	users , err := h.UserServise.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, GetUsersResponse{
		Users: users,
	})
}

func (h *AdminHandler) GetFood(c *gin.Context){
	foodList, err := h.FoodService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, GetFoodsResponse{
		foodList,
	})
}
