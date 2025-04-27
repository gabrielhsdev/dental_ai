package routes

import (
	"github.com/gabrielhsdev/dental_ai/backend/auth-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, handler handlers.AuthHandlerInterface) {
	authGroup := router.Group("/")
	{
		authGroup.POST("/login", handler.Login)
		authGroup.POST("/register", handler.Register)
		authGroup.POST("/me", handler.Me)
		authGroup.POST("/validate", handler.Validate)
	}
}
