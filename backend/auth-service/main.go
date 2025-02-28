package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/config"
)

func main() {
	// Port
	port := extractPortFlag()

	// Env
	config.LoadEnv()

	// Database
	config.LoadPostgres()

	// Router
	router := gin.Default()

	// Initialize Routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Auth Service Running!"})
	})

	// Start server on the specified port
	log.Printf("Starting server on port %s...\n", port)
	router.Run(":" + port)
}

func extractPortFlag() string {
	portFlag := flag.String("port", "8081", "Port for the auth service")
	flag.Parse()
	port := *portFlag
	return port
}
