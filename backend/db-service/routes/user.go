package routes

import (
	"github.com/gabrielhsdev/dental_ai/backend/db-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, handler handlers.UserHandlerInterface) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/:id", handler.GetUserById)
		userGroup.GET("/email", handler.GetUserByEmail)
	}
}
