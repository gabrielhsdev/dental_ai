package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/handlers"
)

func UserRoutes(router *gin.Engine, handler *handlers.UserHandler) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/:id", handler.GetUserById)
	}
}
