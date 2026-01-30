package ui

import (
	"net/http"
	"time"

	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/internal/delivery/web/middelware"
	"github.com/arya237/foodPilot/pkg/logger"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	TokenExpiry time.Duration
	UserService services.UserService
	logger      logger.Logger
}

func NewAuthHandler(expiry time.Duration, u services.UserService) *AuthHandler {
	return &AuthHandler{
		TokenExpiry: expiry,
		UserService: u,
		logger:      logger.New("webAuthHandler"),
	}
}

// HandleLoginPage serves the login page
func (h *AuthHandler) HandleLoginPage(c *gin.Context) {
	data := gin.H{
		"IsAuthenticated": false,
		"Error":            c.Query("error"),
	}
	c.HTML(http.StatusOK, "login.html", data)
}

// HandleLoginSubmit processes login form submission
func (h *AuthHandler) HandleLoginSubmit(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		data := gin.H{
			"IsAuthenticated": false,
			"Error":            "Username and password are required",
		}
		c.HTML(http.StatusBadRequest, "login.html", data)
		return
	}

	// Login with user service
	user, err := h.UserService.Login(username, password)
	h.logger.Warn("Login with user service", logger.Field{Key: "username", Value: username})
	if err != nil {
		data := gin.H{
			"IsAuthenticated": false,
			"Error":            err.Error(),
		}
		c.HTML(http.StatusBadRequest, "login.html", data)
		return
	}

	// Generate JWT token
	jwtToken, err := middelware.GenerateJWT(user, h.TokenExpiry)
	if err != nil {
		data := gin.H{
			"IsAuthenticated": false,
			"Error":            "Could not generate token",
		}
		c.HTML(http.StatusInternalServerError, "login.html", data)
		return
	}

	// Set JWT in HTTP-only cookie
	c.SetCookie("auth_token", jwtToken, int(h.TokenExpiry.Seconds()), "/", "", false, true)

	// Redirect to home page
	c.Redirect(http.StatusSeeOther, "/home")
}

// HandleSignupPage serves the signup page
func (h *AuthHandler) HandleSignupPage(c *gin.Context) {
	data := gin.H{
		"IsAuthenticated": false,
		"Error":            c.Query("error"),
	}
	c.HTML(http.StatusOK, "signup.html", data)
}

// HandleSignupSubmit processes signup form submission
func (h *AuthHandler) HandleSignupSubmit(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		data := gin.H{
			"IsAuthenticated": false,
			"Error":            "Username and password are required",
		}
		c.HTML(http.StatusBadRequest, "signup.html", data)
		return
	}

	// Register user
	user, err := h.UserService.SignUp(username, password)
	if err != nil {
		data := gin.H{
			"IsAuthenticated": false,
			"Error":            err.Error(),
		}
		c.HTML(http.StatusBadRequest, "signup.html", data)
		return
	}

	// Generate JWT token
	jwtToken, err := middelware.GenerateJWT(user, h.TokenExpiry)
	if err != nil {
		data := gin.H{
			"IsAuthenticated": false,
			"Error":            "Could not generate token",
		}
		c.HTML(http.StatusInternalServerError, "signup.html", data)
		return
	}

	// Set JWT in HTTP-only cookie
	c.SetCookie("auth_token", jwtToken, int(h.TokenExpiry.Seconds()), "/", "", false, true)

	// Redirect to home page
	c.Redirect(http.StatusSeeOther, "/home")
}

// HandleLogout handles user logout
func (h *AuthHandler) HandleLogout(c *gin.Context) {
	// Clear the auth cookie
	c.SetCookie("auth_token", "", -1, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/login")
}

