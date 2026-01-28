package ui

import (
	"net/http"
	"strconv"

	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/pkg/logger"
	"github.com/gin-gonic/gin"
)

type HomeHandler struct {
	UserService services.UserService
	logger      logger.Logger
}

func NewHomeHandler(u services.UserService) *HomeHandler {
	return &HomeHandler{
		UserService: u,
		logger:      logger.New("webHomeHandler"),
	}
}

// HandleHomePage serves the protected home page with food list
func (h *HomeHandler) HandleHomePage(c *gin.Context) {
	// Get user ID from context (set by middleware)
	userIDStr, exists := c.Get("userID")
	if !exists {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	userID, err := strconv.Atoi(userIDStr.(string))
	if err != nil {
		h.logger.Info("Error converting userID: " + err.Error())
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	// Get all foods
	foods, err := h.UserService.ViewFoods()
	if err != nil {
		data := gin.H{
			"IsAuthenticated": true,
			"Error":           "Failed to load foods: " + err.Error(),
			"Foods":           []interface{}{},
			"UserRatings":     make(map[string]int),
		}
		c.HTML(http.StatusInternalServerError, "home.html", data)
		return
	}

	// Get user's existing ratings
	userRatings, err := h.UserService.ViewRating(userID)
	if err != nil {
		// If no ratings exist, that's okay - just use empty map
		userRatings = make(map[string]int)
	}

	data := gin.H{
		"IsAuthenticated": true,
		"Foods":           foods,
		"UserRatings":     userRatings,
		"Error":           c.Query("error"),
		"Success":         c.Query("success"),
	}

	c.HTML(http.StatusOK, "home.html", data)
}

// HandleRateSubmit processes rating form submission
func (h *HomeHandler) HandleRateSubmit(c *gin.Context) {
	// Get user ID from context
	userIDStr, exists := c.Get("userID")
	if !exists {
		c.Redirect(http.StatusSeeOther, "/login?error=unauthorized")
		return
	}

	// Parse form data - foods come as foods[foodName]=rating
	foodsMap := make(map[string]int)
	
	// Parse the form first
	if err := c.Request.ParseForm(); err != nil {
		c.Redirect(http.StatusSeeOther, "/home?error=invalid+form+data")
		return
	}

	// Get all form values with prefix "foods["
	for key, values := range c.Request.PostForm {
		if len(key) > 6 && key[:6] == "foods[" && len(key) > 7 && key[len(key)-1:] == "]" {
			foodName := key[6 : len(key)-1]
			if len(values) > 0 && values[0] != "" {
				rating, err := strconv.Atoi(values[0])
				if err == nil && rating >= 0 && rating <= 100 {
					foodsMap[foodName] = rating
				}
			}
		}
	}

	if len(foodsMap) == 0 {
		c.Redirect(http.StatusSeeOther, "/home?error=no+ratings+provided")
		return
	}

	// Save ratings
	_, err := h.UserService.RateFoods(userIDStr.(string), foodsMap)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/home?error="+err.Error())
		return
	}

	c.Redirect(http.StatusSeeOther, "/home?success=Ratings+saved+successfully")
}

