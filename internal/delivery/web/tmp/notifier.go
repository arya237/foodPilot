package tmp

import (
	"net/http"

	"github.com/arya237/foodPilot/internal/models"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}
type broadcastRequest struct {
	Provider models.IdProvider
	Msg      string
}
type broadcastResponse struct {
	Response string
}

// Broadcast    godoc
// @Summary     Broadcast message to users
// @Description Send a broadcast message to all users of a specific provider
// @Tags        tmp
// @Accept      json
// @Produce     json
// @Param       request body broadcastRequest true "Broadcast message details"
// @Success     200 {object} broadcastResponse
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /api/tmp/broadcast [POST]
func (h *handler) broadcast(c *gin.Context) {
	var req broadcastRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	if err := h.notifier.Broadcast(req.Provider, req.Msg); err != nil {

		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, broadcastResponse{
		Response: "send",
	})
}
