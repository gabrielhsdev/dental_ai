package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/config"
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
		config.SendResponse(context, http.StatusBadRequest, "Invalid Input", nil, err)
	}

	// Find User Via Email
	storedUser, err := handler.Userservice.GetUserByEmail(user.Email)
	if err != nil || storedUser == nil {
		config.SendResponse(context, http.StatusUnauthorized, "Invalid Username Or Password", nil, err)
		return
	}

	// Check Password
	if storedUser.Password != user.Password {
		config.SendResponse(context, http.StatusUnauthorized, "Invalid Username Or Password", nil, nil)
		return
	}

	// Generate JWT
	token, err := utils.GenerateJWT(storedUser.Id, storedUser.Username)
	if err != nil {
		config.SendResponse(context, http.StatusInternalServerError, "Failed to generate token", nil, err)
		return
	}

	// Returns JWT
	config.SendResponse(context, http.StatusOK, "Login Success", gin.H{"token": token}, nil)
}

func (handler *AuthHandler) Register(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		config.SendResponse(context, http.StatusBadRequest, "Invalid Input", nil, err)
		return
	}

	registeredUser, err := handler.Userservice.RegisterUser(&user)
	if err != nil {
		config.SendResponse(context, http.StatusInternalServerError, "Failed to register user", nil, err)
		return
	}

	config.SendResponse(context, http.StatusOK, "User Registered", registeredUser, nil)
}

func (handler *AuthHandler) Me(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		config.SendResponse(context, http.StatusUnauthorized, "Unauthorized, Token Not Found", nil, nil)
		return
	}

	claims, err := utils.ValidateJWT(token)
	if err != nil {
		config.SendResponse(context, http.StatusUnauthorized, "Unauthorized", nil, err)
		return
	}

	userId := int(claims["sub"].(float64))
	user, err := handler.Userservice.GetUserById(userId)
	if err != nil {
		config.SendResponse(context, http.StatusNotFound, "User not found", nil, err)
		return
	}

	config.SendResponse(context, http.StatusOK, "User Found", user, nil)
}

func (handler *AuthHandler) Validate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		config.SendResponse(context, http.StatusUnauthorized, "Unauthorized, Token Not Found", nil, nil)
		return
	}

	claims, err := utils.ValidateJWT(token)
	if err != nil {
		config.SendResponse(context, http.StatusUnauthorized, "Unauthorized", nil, err)
		return
	}

	config.SendResponse(context, http.StatusOK, "Token Valid", claims, nil)
}
