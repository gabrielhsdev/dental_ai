package handlers

import (
	"log"
	"net/http"

	"github.com/gabrielhsdev/dental_ai/backend/auth-service/internal/service"
	"github.com/gabrielhsdev/dental_ai/backend/auth-service/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandlerInterface interface {
	GetUserById(context *gin.Context)
}

type UserHandler struct {
	UserService     *service.UserService
	ResponseManager response.ResponseManagerInterface
}

func NewUserHandler(userService service.UserServiceInterface, responseManager response.ResponseManagerInterface) UserHandlerInterface {
	return &UserHandler{
		UserService:     userService.(*service.UserService),
		ResponseManager: responseManager,
	}
}

func (handler *UserHandler) GetUserById(context *gin.Context) {
	idString := context.Param("id")
	if idString == "" {
		handler.ResponseManager.Send(context, http.StatusBadRequest, "Id is required", nil, nil)
		return
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		handler.ResponseManager.Send(context, http.StatusBadRequest, "Invalid Id", nil, err)
		return
	}

	user, err := handler.UserService.GetUserById(id)
	if err != nil {
		log.Println(err)
		handler.ResponseManager.Send(context, http.StatusInternalServerError, "Failed to get user", nil, err)
		return
	}
	if user == nil {
		handler.ResponseManager.Send(context, http.StatusNotFound, "User not found", nil, nil)
		return
	}

	handler.ResponseManager.Send(context, http.StatusOK, "User found", user, nil)
}
