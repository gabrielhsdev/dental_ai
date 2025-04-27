package handlers

import (
	"errors"
	"net/http"

	"github.com/gabrielhsdev/dental_ai/backend/auth-service/internal/models"
	"github.com/gabrielhsdev/dental_ai/backend/auth-service/internal/service"
	"github.com/gabrielhsdev/dental_ai/backend/auth-service/pkg/jwt"
	"github.com/gabrielhsdev/dental_ai/backend/auth-service/pkg/logger"
	"github.com/gabrielhsdev/dental_ai/backend/auth-service/pkg/resources"
	"github.com/gabrielhsdev/dental_ai/backend/auth-service/pkg/response"
	"github.com/gin-gonic/gin"
)

type AuthHandlerInterface interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
	Me(context *gin.Context)
	Validate(context *gin.Context)
}

type AuthHandler struct {
	UserService     service.UserServiceInterface
	Logger          logger.LoggerInterface
	JWTManager      jwt.JWTManagerInterface
	ResponseManager response.ResponseManagerInterface
	ResourceManager resources.ResourceManagerInterface
}

func NewAuthHandler(
	userService service.UserServiceInterface,
	logger logger.LoggerInterface,
	jwtManager jwt.JWTManagerInterface,
	responseManager response.ResponseManagerInterface,
	resourceManager resources.ResourceManagerInterface,
) AuthHandlerInterface {
	return &AuthHandler{
		UserService:     userService.(*service.UserService),
		Logger:          logger,
		JWTManager:      jwtManager,
		ResponseManager: responseManager,
		ResourceManager: resourceManager,
	}
}

func (handler *AuthHandler) Login(context *gin.Context) {
	var user models.User
	var action string = "Login"

	if err := context.ShouldBindJSON(&user); err != nil {
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetAuthenticationResource(), nil)
		handler.ResponseManager.Send(context, http.StatusBadRequest, "Invalid Input", nil, err)
	}

	storedUser, err := handler.UserService.GetUserByEmail(user.Email)
	if err != nil || storedUser == nil {
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetAuthenticationResource(), nil)
		handler.ResponseManager.Send(context, http.StatusUnauthorized, "Invalid Username Or Password", nil, err)
		return
	}

	if storedUser.Password != user.Password {
		err := errors.New("invalid credentials")
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetAuthenticationResource(), nil)
		handler.ResponseManager.Send(context, http.StatusUnauthorized, "Invalid Username Or Password", nil, nil)
		return
	}

	token, err := handler.JWTManager.Generate(storedUser.Id, storedUser.Password)
	if err != nil {
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetAuthenticationResource(), nil)
		handler.ResponseManager.Send(context, http.StatusInternalServerError, "Failed to generate token", nil, err)
		return
	}

	handler.Logger.Info(context, action, handler.ResourceManager.GetAuthenticationResource(), map[string]interface{}{"user": storedUser.Email})
	handler.ResponseManager.Send(context, http.StatusOK, "Login Success", gin.H{"token": token}, nil)
}

func (handler *AuthHandler) Register(context *gin.Context) {
	var user models.User
	var action string = "Register"

	if err := context.ShouldBindJSON(&user); err != nil {
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetAuthenticationResource(), nil)
		handler.ResponseManager.Send(context, http.StatusBadRequest, "Invalid Input", nil, err)
		return
	}

	registeredUser, err := handler.UserService.RegisterUser(&user)
	if err != nil {
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetAuthenticationResource(), nil)
		handler.ResponseManager.Send(context, http.StatusInternalServerError, "Failed to register user", nil, err)
		return
	}

	handler.Logger.Info(context, action, handler.ResourceManager.GetAuthenticationResource(), map[string]interface{}{"user": registeredUser.Email})
	handler.ResponseManager.Send(context, http.StatusOK, "User Registered", registeredUser, nil)
}

func (handler *AuthHandler) Me(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		err := errors.New("unauthorized, token not found")
		handler.Logger.Error(context, "Me", err, handler.ResourceManager.GetAuthenticationResource(), nil)
		handler.ResponseManager.Send(context, http.StatusUnauthorized, "Unauthorized, Token Not Found", nil, nil)
		return
	}

	claims, err := handler.JWTManager.Validate(token)
	if err != nil {
		handler.Logger.Error(context, "Me", err, handler.ResourceManager.GetAuthenticationResource(), nil)
		handler.ResponseManager.Send(context, http.StatusUnauthorized, "Unauthorized", nil, err)
		return
	}

	userId := claims.Sub
	user, err := handler.UserService.GetUserById(userId)
	if err != nil {
		handler.Logger.Error(context, "Me", err, handler.ResourceManager.GetAuthenticationResource(), nil)
		handler.ResponseManager.Send(context, http.StatusNotFound, "User not found", nil, err)
		return
	}

	handler.Logger.Info(context, "Me", handler.ResourceManager.GetAuthenticationResource(), map[string]interface{}{"user": user.Email})
	handler.ResponseManager.Send(context, http.StatusOK, "User Found", user, nil)
}

func (handler *AuthHandler) Validate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		err := errors.New("unauthorized, token not found")
		handler.Logger.Error(context, "Validate", err, handler.ResourceManager.GetAuthenticationResource(), nil)
		handler.ResponseManager.Send(context, http.StatusUnauthorized, "Unauthorized, Token Not Found", nil, nil)
		return
	}

	claims, err := handler.JWTManager.Validate(token)
	if err != nil {
		handler.Logger.Error(context, "Validate", err, handler.ResourceManager.GetAuthenticationResource(), nil)
		handler.ResponseManager.Send(context, http.StatusUnauthorized, "Unauthorized", nil, err)
		return
	}

	handler.Logger.Info(context, "Validate", handler.ResourceManager.GetAuthenticationResource(), map[string]interface{}{"claims": claims})
	handler.ResponseManager.Send(context, http.StatusOK, "Token Valid", claims, nil)
}
