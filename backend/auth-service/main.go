package main

import (
	"log"

	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/handlers"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/repository"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/service"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/pkg/database"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/pkg/environment"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/pkg/logger"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/pkg/mode"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/routes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	/*
		TODO: - Finish up the logger and understand how to use it better
		So, now I have my logger under pkg/logger, and my auditLogs
		table w/ repository handler  model and service, I need to tie both together right ?
		So when I call my log from pkg I actually save in the
		db the proxy_set_headers from nginx.
		In order to do that, we need to inject our logger manager into our responseManager since we will have context over there
		loggerManager := logger.LoadLogger()
	*/

	modeManager := mode.NewModeManager()
	envManager := environment.NewEnvManager(modeManager)
	/* httpClient := request.NewHttpClient() This is the httpManager so we can make request between our services */

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
	logger := logger.NewLogger(zapLogger, auditLogsService)

	// Initialize Handler
	authHandler := handlers.NewAuthHandler(userService, logger)
	userHandler := handlers.NewUserHandler(userService)

	// Initialize Router
	router := gin.Default()

	// Initialize Routes
	routes.AuthRoutes(router, authHandler)
	routes.UserRoutes(router, userHandler)

	// Run Server
	port := envManager.GetAuthServicePort()
	router.Run(":" + port)
}
