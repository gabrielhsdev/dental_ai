package handlers

import (
	"log"
	"net/http"

	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/service"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandlerInterface interface {
	GetUserById(context *gin.Context)
}

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService service.UserServiceInterface) UserHandlerInterface {
	return &UserHandler{
		UserService: userService.(*service.UserService),
	}
}

func (handler *UserHandler) GetUserById(context *gin.Context) {
	idString := context.Param("id")
	if idString == "" {
		utils.SendResponse(context, http.StatusBadRequest, "Id is required", nil, nil)
		return
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		utils.SendResponse(context, http.StatusBadRequest, "Invalid Id", nil, err)
		return
	}

	user, err := handler.UserService.GetUserById(id)
	if err != nil {
		log.Println(err)
		utils.SendResponse(context, http.StatusInternalServerError, "Failed to get user", nil, err)
		return
	}
	if user == nil {
		utils.SendResponse(context, http.StatusNotFound, "User not found", nil, nil)
		return
	}

	utils.SendResponse(context, http.StatusOK, "User found", user, nil)
}
