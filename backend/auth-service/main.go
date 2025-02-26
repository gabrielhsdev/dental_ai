package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Port comes from the command line arguments
	portFlag := flag.String("port", "8081", "Port for the auth service")
	flag.Parse()
	port := *portFlag

	// Create a new Gin router
	router := gin.Default()

	// Health check route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Auth Service Running!"})
	})

	// Login route (dummy response)
	router.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "User logged in"})
	})

	// Register route (dummy response)
	router.POST("/register", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "User registered"})
	})

	// Start server on the specified port
	log.Printf("Starting server on port %s...\n", port)
	router.Run(":" + port)
}
