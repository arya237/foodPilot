package ui

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/arya237/foodPilot/internal/delivery/web/middelware"
	"github.com/arya237/foodPilot/internal/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine, tokenExpiry time.Duration, userService services.UserService) error {
	// Load and set HTML templates
	// Parse all templates - base.html must be parsed first so blocks can be overridden
	tmpl, err := template.ParseGlob(filepath.Join("internal", "delivery", "web", "ui", "templates", "*.html"))
	if err != nil {
		return err
	}
	engine.SetHTMLTemplate(tmpl)

	// Create handlers
	authHandler := NewAuthHandler(tokenExpiry, userService)
	homeHandler := NewHomeHandler(userService)

	// Static files
	engine.Static("/statics", "./statics")

	// Public routes (no auth required)
	engine.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/login")
	})
	engine.GET("/login", authHandler.HandleLoginPage)
	engine.POST("/login", authHandler.HandleLoginSubmit)
	engine.GET("/signup", authHandler.HandleSignupPage)
	engine.POST("/signup", authHandler.HandleSignupSubmit)
	engine.GET("/logout", authHandler.HandleLogout)

	// Protected routes (require auth)
	protected := engine.Group("")
	protected.Use(middelware.WebAuthMiddleware())
	{
		protected.GET("/home", homeHandler.HandleHomePage)
		protected.POST("/rate", homeHandler.HandleRateSubmit)
	}

	return nil
}
