package routes

import (
	"github.com/gabrielhsdev/dental_ai/backend/db-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func PatientImagesRoutes(router *gin.Engine, handler handlers.PatientImagesHandlerInterface) {
	userGroup := router.Group("/patientsImages")
	{
		userGroup.POST("/", handler.Create)
		userGroup.GET("/:id", handler.GetById)
		userGroup.GET("/patient/:patientId", handler.GetByPatientId)
		userGroup.GET("/user/:userId", handler.GetByUserId)
	}
}
