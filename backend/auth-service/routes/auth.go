package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/handlers"
)

// AuthRoutes sets up the routes for authentication
func AuthRoutes(router *gin.Engine, handler *handlers.AuthHandler) {
	authGroup := router.Group("/core")
	{
		authGroup.POST("/login", handler.Login)
		authGroup.POST("/register", handler.Register)
		authGroup.POST("/me", handler.Me)
		authGroup.POST("/validate", handler.Validate)
	}
}
