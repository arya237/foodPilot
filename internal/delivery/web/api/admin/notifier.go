package admin

import (
	"net/http"

	"github.com/arya237/foodPilot/internal/models"
	"github.com/gin-gonic/gin"
)

type broadcastRequest struct {
	provider models.IdProvider
	msg      string
}
type broadcastResponse struct {
	response string
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
func (h *AdminHandler) broadcast(c *gin.Context) {
	var req broadcastRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if err := h.notifier.Broadcast(req.provider, req.msg); err != nil {

		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, broadcastResponse{
		response: "send",
	})
}
