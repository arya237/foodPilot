package handler

import (
	"github.com/arya237/foodPilot/internal/handler/auth"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()
	setup(r)
	return r
}

func setup(r *gin.Engine) {
	auth.RegisterRoutes(r.Group("/auth"))
}