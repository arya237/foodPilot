package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AutoSave     godoc
// @Summary     Toggle user autosave
// @Description Toggle user autosave attribute
// @Tags        user
// @Accept      json
// @Param user  body AutoSaveRequest true "Toggle info"
// @Security    BearerAuth
// @Success     200 {object} AutoSaveResponse
// @Failure     404 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /user/autosave [POST]
func (h *UserHandler) AutoSave(c *gin.Context) {

	req := AutoSaveRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request"})
		return
	}

	id, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "user not found"})
		return
	}
	userID, _ := strconv.Atoi(id.(string))

	err := h.UserService.ToggleAutoSave(userID, *req.AutoSave)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "server error"})
		return
	}

	c.JSON(http.StatusOK, AutoSaveResponse{
		Message: "Auto save updated",
	})
}
