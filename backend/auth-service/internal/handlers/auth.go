package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/internal/models"
	"main.go/internal/service"
	"main.go/utils"
)

type AuthHandler struct {
	Userservice *service.UserService
}

func NewAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{Userservice: userService}
}

func (handler *AuthHandler) Login(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
	}

	// Find User Via Email
	storedUser, err := handler.Userservice.GetUserByEmail(user.Email)
	if err != nil || storedUser == nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Username Or Password"})
		return
	}

	// Check Password
	if storedUser.Password != user.Password {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Username Or Password"})
		return
	}

	// Generate JWT
	token, err := utils.GenerateJWT(storedUser.Id, storedUser.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Returns JWT
	context.JSON(http.StatusOK, gin.H{"token": token})
}
