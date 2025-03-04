package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/internal/handlers"
)

// AuthRoutes sets up the routes for authentication
func AuthRoutes(router *gin.Engine, handler *handlers.AuthHandler) {
	authGroup := router.Group("/core")
	{
		authGroup.POST("/login", handler.Login)
	}
}
