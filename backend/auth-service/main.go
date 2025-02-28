package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"main.go/config"
)

func main() {
	// Initialize Configurations
	config.LoadEnv()
	config.LoadPostgres()
	router := gin.Default()
	port := os.Getenv("AUTH_SERVICE_PORT")

	// Initialize Routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Auth Service Running!"})
	})

	// Log the port
	log.Printf("Starting server on port %s...\n", port)
	router.Run(":" + port)
}
