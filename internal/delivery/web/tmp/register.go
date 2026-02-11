package tmp

import (
	"github.com/arya237/foodPilot/internal/services/admin"
	"github.com/gin-gonic/gin"
)

type handler struct {
	notifier admin.Notifier
}

func New(notifier admin.Notifier) *handler {
	return &handler{
		notifier: notifier,
	}
}

func RegisterRoutes(group *gin.RouterGroup, handler *handler) {
	group.POST("/broadcast", handler.broadcast)
}
