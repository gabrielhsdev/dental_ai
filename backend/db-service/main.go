package main

import (
	"log"

	"github.com/gabrielhsdev/dental_ai/backend/db-service/internal/handlers"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/internal/repository"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/internal/service"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/pkg/database"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/pkg/environment"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/pkg/headers"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/pkg/jwt"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/pkg/logger"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/pkg/mode"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/pkg/resources"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/pkg/response"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/routes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	modeManager := mode.NewModeManager()
	envManager := environment.NewEnvManager(modeManager)
	headersHandler := headers.NewHeadersManager()
	jwtManager := jwt.NewJWTManager(envManager.GetJWTSecretKey())
	responseManager := response.NewResponseManager()
	resourceManager := resources.NewResourceManager()

	// Initialize Database
	database, err := database.LoadDatabase("postgres", modeManager, envManager)
	if err != nil {
		log.Fatalf("Error loading database: %v", err)
	}

	// Initialize Repositories
	userRepository := repository.NewUserRepository(database)
	auditLogsRepository := repository.NewAuditLogRepository(database)
	patientRepository := repository.NewPatientRepository(database)
	patientImagesRepository := repository.NewPatientImagesRepository(database)

	// Initialize Service
	userService := service.NewUserService(userRepository)
	auditLogsService := service.NewAuditLogsService(auditLogsRepository)
	patientService := service.NewPatientService(patientRepository)
	patientImagesService := service.NewPatientImagesService(patientImagesRepository)

	// Initialize logger
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapLogger, err := config.Build(zap.AddCaller())
	if err != nil {
		panic("Failed to initialize logger")
	}
	loggerManager := logger.NewLogger(zapLogger, auditLogsService, headersHandler, resourceManager)

	// Initialize Handler
	userHandler := handlers.NewUserHandler(userService, loggerManager, jwtManager, responseManager, resourceManager)
	patientHandler := handlers.NewPatientHandler(patientService, loggerManager, jwtManager, responseManager, resourceManager)
	patientImagesHandler := handlers.NewPatientImagesHandler(patientImagesService, loggerManager, jwtManager, responseManager, resourceManager)

	// Initialize Router
	router := gin.Default()

	// Initialize Routes
	routes.UserRoutes(router, userHandler)
	routes.PatientRoutes(router, patientHandler)
	routes.PatientImagesRoutes(router, patientImagesHandler)

	// Run Server
	router.Run(":" + envManager.GetDBServicePort())
}
