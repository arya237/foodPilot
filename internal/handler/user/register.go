package user

import (
	"github.com/arya237/foodPilot/internal/auth"
	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/pkg/logger"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService services.UserService
	Logger      logger.Logger
}

func NewUserHandler(u services.UserService) *UserHandler {
	return &UserHandler{
		UserService: u,
		Logger:      logger.New("userHandler"),
	}
}

func RegisterRoutes(group *gin.RouterGroup, userHandler *UserHandler) {

	group.Use(auth.AuthMiddleware())
	group.POST("/autosave", userHandler.AutoSave)
}
