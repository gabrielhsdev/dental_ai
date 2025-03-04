package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/internal/handlers"
)

func UserRoutes(router *gin.Engine, handler *handlers.UserHandler) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/:id", handler.GetUserById)        // PROTECTED
		userGroup.POST("/register", handler.RegisterUser) // NOT PROTECTED
	}
}
