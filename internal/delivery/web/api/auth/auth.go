package auth

import (
	"net/http"
	"time"

	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/pkg/logger"

	"github.com/arya237/foodPilot/internal/delivery/web/middelware"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	TokenExpiry time.Duration
	UserService services.UserService
	logger      logger.Logger
}

func NewHandler(expiry time.Duration, u services.UserService) *AuthHandler {
	return &AuthHandler{
		TokenExpiry: expiry,
		UserService: u,
		logger:      logger.New("loginHandler"),
	}
}

func RegisterRoutes(group *gin.RouterGroup, loginHandler *AuthHandler) {
	group.POST("/login", loginHandler.HandleLogin)
	group.POST("/signup", loginHandler.HandleSignUp)
}

// ***************** methodes *********************************//

// Login        godoc
// @Summary     login a user
// @Description Login a user to app and generate code for it
// @Tags        Auth
// @Accept      json
// @Param       login body LoginRequest true "Login info"
// @Produce     json
// @Success     200 {object} LoginResponse
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /api/auth/login [POST]
func (h *AuthHandler) HandleLogin(c *gin.Context) {
	// Get request  information
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request"})
		return
	}

	//login with user
	user, err := h.UserService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	// generate token for this user
	jwtToken, err := middelware.GenerateJWT(user, h.TokenExpiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: jwtToken})
}

// SignUp        godoc
// @Summary     signup a new user
// @Description Register a new user and generate token for it
// @Tags        Auth
// @Accept      json
// @Param       signup body SignUpRequest true "Signup info"
// @Produce     json
// @Success     201 {object} SignUpResponse
// @Failure     400 {object} ErrorResponse
// @Failure     500 {object} ErrorResponse
// @Router      /api/auth/signup [POST]
func (h *AuthHandler) HandleSignUp(c *gin.Context) {
	// Bind request info
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request"})
		return
	}

	// Register user
	user, err := h.UserService.SignUp(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	// generate token for this user
	jwtToken, err := middelware.GenerateJWT(user, h.TokenExpiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "could not generate token"})
		return
	}

	c.JSON(http.StatusCreated, SignUpResponse{
		Message: "User registered successfully",
		Token:   jwtToken,
	})
}
