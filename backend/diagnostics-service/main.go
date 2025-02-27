package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Port comes from the command line arguments
	portFlag := flag.String("port", "8083", "Port for the diagnostics service")
	flag.Parse()
	port := *portFlag

	// Create a new Gin router
	router := gin.Default()

	// Health check route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Diagnostics Service Running!"})
	})

	// Start server on the specified port
	log.Printf("Starting server on port %s...\n", port)
	router.Run(":" + port)
}
