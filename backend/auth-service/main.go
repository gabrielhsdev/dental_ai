package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
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

	// Start server on port 8081
	router.Run(":8081")
}
