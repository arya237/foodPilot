package auth

import (
	"net/http"
	"time"

	"github.com/arya237/foodPilot/internal/auth"
	"github.com/gin-gonic/gin"
)


type LoginHandler struct {
	TokenExpiry time.Duration
}

func NewLoginHandler(expiry time.Duration) *LoginHandler {
	return &LoginHandler{TokenExpiry: expiry}
}

func RegisterRoutes(group *gin.RouterGroup) {
	h := NewLoginHandler(time.Hour)
	group.POST("/login", h.HandleLogin)
}
// ***************** methodes *********************************//

func (h *LoginHandler) HandleLogin(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	
	// TODO:get use info ....	


	token, err := auth.GenerateJWT("some id", "jwt token", h.TokenExpiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		Token: token,
	})
}
