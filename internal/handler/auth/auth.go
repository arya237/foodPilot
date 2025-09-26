package auth

import (
	"net/http"
	"time"

	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/pkg/logger"

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
}

// ***************** methodes *********************************//

// Login        godoc
// @Summary     login a user
// @Description Login a user to app and generate code for it
// @Tags        auth
// @Accept      json
// @Param       login body LoginRequest true "Login info"
// @Produce     json
// @Success     200 {object} LoginResponse
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /auth/login [POST]
func (h *LoginHandler) HandleLogin(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request"})
		return
	}

	ID, token, err := h.UserService.Login(req.Username, req.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	jwtToken, err := auth.GenerateJWT(ID, token, h.TokenExpiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: jwtToken})
}
