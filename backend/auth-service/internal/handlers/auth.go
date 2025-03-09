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

func (handler *AuthHandler) Register(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	registeredUser, err := handler.Userservice.RegisterUser(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	context.JSON(http.StatusOK, registeredUser)
}

func (handler *AuthHandler) Me(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, Token Not Found"})
		return
	}

	claims, err := utils.ValidateJWT(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": err.Error()})
		return
	}

	userId := int(claims["sub"].(float64))
	user, err := handler.Userservice.GetUserById(userId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	context.JSON(http.StatusOK, user)
}

func (handler *AuthHandler) Validate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	claims, err := utils.ValidateJWT(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	context.JSON(http.StatusOK, claims)
}
