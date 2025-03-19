package main

import (
	"log"
	"os"

	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/config"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/database"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/handlers"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/repository"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/service"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Configurations
	config.LoadEnv()
	database, err := database.LoadDatabase("postgres")
	if err != nil {
		log.Fatalf("Error loading database: %v", err)
	}

	// Initialize Repositories
	userRepository := repository.NewUserRepository(database)

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
	serviceName := os.Getenv("AUTH_SERVICE_HOST")
	log.Printf("Starting %s on port %s", serviceName, port)
	router.Run(":" + port)
}
