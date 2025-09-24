package auth

import (
	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/pkg/logger"
	"net/http"
	"time"

	"github.com/arya237/foodPilot/internal/auth"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	TokenExpiry time.Duration
	UserService services.UserService
	logger      logger.Logger
}

func NewLoginHandler(expiry time.Duration, u services.UserService) *LoginHandler {
	return &LoginHandler{
		TokenExpiry: expiry,
		UserService: u,
		logger:      logger.New("loginHandler"),
	}
}

func RegisterRoutes(group *gin.RouterGroup, loginHandler *LoginHandler) {
	group.POST("/login", loginHandler.HandleLogin)
	group.PUT("/autosave",auth.AuthMiddleware() , loginHandler.AutoSave)

}

// ***************** methodes *********************************//

func (h *LoginHandler) HandleLogin(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	ID, token, err := h.UserService.Login(req.Username, req.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwtToken, err := auth.GenerateJWT(ID, token, h.TokenExpiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"loginResponse": LoginResponse{
		Token: jwtToken,
	}, "message": "login successful"})
}

func (h *LoginHandler) AutoSave(c *gin.Context) {

	req := AutoSaveRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	id, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
	}

	h.UserService.ToggleAutoSave(id.(int), req.AutoSave)

	c.JSON(http.StatusOK, AutoSaveResponse{
		Message: "Auto save updated",
	})
}