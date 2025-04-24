package routes

import (
	"github.com/gabrielhsdev/dental_ai/backend/db-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

// AuthRoutes sets up the routes for authentication
func AuthRoutes(router *gin.Engine, handler handlers.AuthHandlerInterface) {
	authGroup := router.Group("/core")
	{
		authGroup.POST("/login", handler.Login)
		authGroup.POST("/register", handler.Register)
		authGroup.POST("/me", handler.Me)
		authGroup.POST("/validate", handler.Validate)
	}
}
