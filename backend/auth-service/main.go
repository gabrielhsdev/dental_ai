package main

import (
	"log"
	"os"

	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/config"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/handlers"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/repository"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/service"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Configurations
	config.LoadEnv()
	config.LoadPostgres()

	// Initialize Repositories
	userRepository := repository.NewUserRepository(config.DB)

	// Initialize Service
	userService := service.NewUserService(userRepository)

	// Initialize Handler
	authHandler := handlers.NewAuthHandler(userService)
	userHandler := handlers.NewUserHandler(userService)

	// Initialize Router
	router := gin.Default()

	// Initialize Routes
	routes.AuthRoutes(router, authHandler)
	routes.UserRoutes(router, userHandler)

	// Run Server
	port := os.Getenv("AUTH_SERVICE_PORT")
	log.Printf("Starting server on port %s...\n", port)
	router.Run(":" + port)
}
