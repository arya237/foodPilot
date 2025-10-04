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

	// Get user to retrieve role
	user, err := h.UserService.GetByUserName(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "could not retrieve user information"})
		return
	}

	jwtToken, err := auth.GenerateJWT(ID, token, user.Role, h.TokenExpiry)
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
// @Router      /auth/signup [POST]
func (h *LoginHandler) HandleSignUp(c *gin.Context) {
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

	// Generate access token for the new user
	userID, samadToken, err := h.UserService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "user registered but could not generate token"})
		return
	}

	// Generate JWT with user role
	jwtToken, err := auth.GenerateJWT(userID, samadToken, user.Role, h.TokenExpiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "could not generate token"})
		return
	}

	c.JSON(http.StatusCreated, SignUpResponse{
		Message: "User registered successfully",
		Token:   jwtToken,
	})
}
