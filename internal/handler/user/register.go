package user

import (
	"time"

	"github.com/arya237/foodPilot/internal/auth"
	"github.com/arya237/foodPilot/internal/services"
	"github.com/arya237/foodPilot/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

type UserHandler struct {
	UserService services.UserService
	RateService services.RateFoodService
	Logger      logger.Logger
}

func NewUserHandler(u services.UserService, r services.RateFoodService) *UserHandler {
	return &UserHandler{
		UserService: u,
		RateService: r,
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
	group.GET("/rates", userHandler.GetRates)
}
