package main

import (
	"log"

	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/handlers"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/repository"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/service"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/pkg/database"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/pkg/environment"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/pkg/headers"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/pkg/jwt"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/pkg/logger"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/pkg/mode"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/routes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	modeManager := mode.NewModeManager()
	envManager := environment.NewEnvManager(modeManager)
	headersHandler := headers.NewHeadersManager()
	jwtManager := jwt.NewJWTManager(envManager.GetJWTSecretKey())

	// Initialize Database
	database, err := database.LoadDatabase("postgres", modeManager, envManager)
	if err != nil {
		log.Fatalf("Error loading database: %v", err)
	}

	// Initialize Repositories
	userRepository := repository.NewUserRepository(database)
	auditLogsRepository := repository.NewAuditLogRepository(database)

	// Initialize Service
	userService := service.NewUserService(userRepository)
	auditLogsService := service.NewAuditLogsService(auditLogsRepository)

	// Initialize logger
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapLogger, err := config.Build(zap.AddCaller())
	if err != nil {
		panic("Failed to initialize logger")
	}
	loggerManager := logger.NewLogger(zapLogger, auditLogsService, headersHandler)

	// Initialize Handler
	authHandler := handlers.NewAuthHandler(userService, loggerManager, jwtManager)
	userHandler := handlers.NewUserHandler(userService)

	// Initialize Router
	router := gin.Default()

	// Initialize Routes
	routes.AuthRoutes(router, authHandler)
	routes.UserRoutes(router, userHandler)

	// Run Server
	router.Run(":" + envManager.GetAuthServicePort())
}
