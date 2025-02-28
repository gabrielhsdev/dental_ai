package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/config"
)

func main() {
	// Initialize the server
	port := config.ExtractPortFlag()
	config.LoadEnv()
	config.LoadPostgres()
	router := gin.Default()

	// Initialize Routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Auth Service Running!"})
	})

	// Log the port
	log.Printf("Starting server on port %s...\n", port)
	router.Run(":" + port)
}
