package user

import (
	"github.com/arya237/foodPilot/internal/auth"
	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"time"
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

	rate := limiter.Rate{
		Period: 3 * time.Second,
		Limit:  2,
	}

	store := memory.NewStore()
	limiter := limiter.New(store, rate)

	group.Use(auth.AuthMiddleware(), auth.LimitMiddelware(limiter))
	group.POST("/autosave", userHandler.AutoSave)
}
