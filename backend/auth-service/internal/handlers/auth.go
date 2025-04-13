package handlers

import (
	"errors"
	"net/http"

	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/models"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/service"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/pkg/logger"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandlerInterface interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
	Me(context *gin.Context)
	Validate(context *gin.Context)
}

type AuthHandler struct {
	UserService service.UserServiceInterface
	Logger      logger.LoggerInterface
}

func NewAuthHandler(userService service.UserServiceInterface, logger logger.LoggerInterface) AuthHandlerInterface {
	return &AuthHandler{
		UserService: userService.(*service.UserService),
		Logger:      logger,
	}
}

func (handler *AuthHandler) Login(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		handler.Logger.Error(context, "Login", err, "user", nil)
		utils.SendResponse(context, http.StatusBadRequest, "Invalid Input", nil, err)
	}

	// Find User Via Email
	storedUser, err := handler.UserService.GetUserByEmail(user.Email)
	if err != nil || storedUser == nil {
		handler.Logger.Error(context, "Login", err, "user", nil)
		utils.SendResponse(context, http.StatusUnauthorized, "Invalid Username Or Password", nil, err)
		return
	}

	// Check Password
	if storedUser.Password != user.Password {
		err := errors.New("invalid credentials")
		handler.Logger.Error(context, "Login", err, "user", nil)
		utils.SendResponse(context, http.StatusUnauthorized, "Invalid Username Or Password", nil, nil)
		return
	}

	// Generate JWT
	token, err := utils.GenerateJWT(storedUser.Id, storedUser.Username)
	if err != nil {
		handler.Logger.Error(context, "Login", err, "user", nil)
		utils.SendResponse(context, http.StatusInternalServerError, "Failed to generate token", nil, err)
		return
	}

	// Returns JWT
	handler.Logger.Info(context, "Login", "Login successful", map[string]interface{}{"user": storedUser.Email})
	utils.SendResponse(context, http.StatusOK, "Login Success", gin.H{"token": token}, nil)
}

func (handler *AuthHandler) Register(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		utils.SendResponse(context, http.StatusBadRequest, "Invalid Input", nil, err)
		return
	}

	registeredUser, err := handler.UserService.RegisterUser(&user)
	if err != nil {
		utils.SendResponse(context, http.StatusInternalServerError, "Failed to register user", nil, err)
		return
	}

	utils.SendResponse(context, http.StatusOK, "User Registered", registeredUser, nil)
}

func (handler *AuthHandler) Me(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		utils.SendResponse(context, http.StatusUnauthorized, "Unauthorized, Token Not Found", nil, nil)
		return
	}

	claims, err := utils.ValidateJWT(token)
	if err != nil {
		utils.SendResponse(context, http.StatusUnauthorized, "Unauthorized", nil, err)
		return
	}

	userId := int(claims["sub"].(float64))
	user, err := handler.UserService.GetUserById(userId)
	if err != nil {
		utils.SendResponse(context, http.StatusNotFound, "User not found", nil, err)
		return
	}

	utils.SendResponse(context, http.StatusOK, "User Found", user, nil)
}

func (handler *AuthHandler) Validate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		utils.SendResponse(context, http.StatusUnauthorized, "Unauthorized, Token Not Found", nil, nil)
		return
	}

	claims, err := utils.ValidateJWT(token)
	if err != nil {
		utils.SendResponse(context, http.StatusUnauthorized, "Unauthorized", nil, err)
		return
	}

	utils.SendResponse(context, http.StatusOK, "Token Valid", claims, nil)
}
