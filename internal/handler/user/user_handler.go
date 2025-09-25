package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) AutoSave(c *gin.Context) {

	req := AutoSaveRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	id, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	userID, _ := strconv.Atoi(id.(string))

	err := h.UserService.ToggleAutoSave(userID, *req.AutoSave)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error!"})
		return
	}

	c.JSON(http.StatusOK, AutoSaveResponse{
		Message: "Auto save updated",
	})
}
