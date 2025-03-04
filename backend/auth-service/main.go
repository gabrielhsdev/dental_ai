package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/internal/handlers"
	"main.go/internal/repository"
	"main.go/internal/service"
	"main.go/routes"
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
