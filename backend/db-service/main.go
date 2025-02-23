package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Health check route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "DB Service Running!"})
	})

	// Simple GET example (returning dummy data)
	router.GET("/animals", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"animals": []string{"Dog", "Cat", "Rabbit"},
		})
	})

	// Start server on port 8082
	router.Run(":8082")
}
