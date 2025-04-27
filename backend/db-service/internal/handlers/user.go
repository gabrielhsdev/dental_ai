package handlers

import (
	"errors"
	"net/http"

	"github.com/gabrielhsdev/dental_ai/backend/db-service/internal/service"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/pkg/jwt"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/pkg/logger"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/pkg/resources"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandlerInterface interface {
	GetUserById(context *gin.Context)
	GetUserByEmail(context *gin.Context)
}

type UserHandler struct {
	UserService     service.UserServiceInterface
	Logger          logger.LoggerInterface
	JWTManager      jwt.JWTManagerInterface
	ResponseManager response.ResponseManagerInterface
	ResourceManager resources.ResourceManagerInterface
}

func NewUserHandler(
	userService service.UserServiceInterface,
	logger logger.LoggerInterface,
	jwtManager jwt.JWTManagerInterface,
	responseManager response.ResponseManagerInterface,
	resourceManager resources.ResourceManagerInterface,
) UserHandlerInterface {
	return &UserHandler{
		UserService:     userService.(*service.UserService),
		Logger:          logger,
		JWTManager:      jwtManager,
		ResponseManager: responseManager,
		ResourceManager: resourceManager,
	}
}

func (handler *UserHandler) GetUserById(context *gin.Context) {
	idString := context.Param("id")
	action := "Get User By Id"

	if idString == "" {
		err := errors.New("invalid credentials")
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetUserResource(), nil)
		handler.ResponseManager.Send(context, http.StatusBadRequest, "Invalid Id Parameter", nil, err)
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetUserResource(), nil)
		handler.ResponseManager.Send(context, http.StatusBadRequest, "Invalid Id", nil, err)
		return
	}

	user, err := handler.UserService.GetUserById(id)
	if err != nil {
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetUserResource(), nil)
		handler.ResponseManager.Send(context, http.StatusInternalServerError, "Failed to get user", nil, err)
		return
	}
	if user == nil {
		err := errors.New("user not found")
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetUserResource(), nil)
		handler.ResponseManager.Send(context, http.StatusNotFound, "User not found", nil, nil)
		return
	}

	handler.Logger.Info(context, action, handler.ResourceManager.GetUserResource(), map[string]interface{}{"user": user})
	handler.ResponseManager.Send(context, http.StatusOK, "User found", user, nil)
}

func (handler *UserHandler) GetUserByEmail(context *gin.Context) {
	email := context.Query("email")
	action := "Get User By Email"

	if email == "" {
		err := errors.New("email parameter is required")
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetUserResource(), nil)
		handler.ResponseManager.Send(context, http.StatusBadRequest, "Email parameter is required", nil, err)
		return
	}

	user, err := handler.UserService.GetUserByEmail(email)
	if err != nil {
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetUserResource(), nil)
		handler.ResponseManager.Send(context, http.StatusInternalServerError, "Failed to get user", nil, err)
		return
	}
	if user == nil {
		err := errors.New("user not found")
		handler.Logger.Error(context, action, err, handler.ResourceManager.GetUserResource(), nil)
		handler.ResponseManager.Send(context, http.StatusNotFound, "User not found", nil, nil)
		return
	}

	handler.Logger.Info(context, action, handler.ResourceManager.GetUserResource(), map[string]interface{}{"user": user})
	handler.ResponseManager.Send(context, http.StatusOK, "User found", user, nil)
}
