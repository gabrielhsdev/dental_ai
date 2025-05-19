package routes

import (
	"github.com/gabrielhsdev/dental_ai/backend/db-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func PatientRoutes(router *gin.Engine, handler handlers.PatientHandlerInterface) {
	userGroup := router.Group("/patients")
	{
		userGroup.GET("/:id", handler.GetPatientById)
		userGroup.POST("/", handler.CreatePatient)
		userGroup.GET("/user", handler.GetPatientsByUserId)
	}
}
