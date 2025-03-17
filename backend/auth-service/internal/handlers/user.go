package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/config"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/service"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (handler *UserHandler) GetUserById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		config.SendResponse(context, http.StatusBadRequest, "Invalid Id", nil, err)
		return
	}

	user, err := handler.Service.GetUserById(id)
	if err != nil {
		log.Println(err)
		config.SendResponse(context, http.StatusInternalServerError, "Failed to get user", nil, err)
		return
	}
	if user == nil {
		config.SendResponse(context, http.StatusNotFound, "User not found", nil, nil)
		return
	}

	config.SendResponse(context, http.StatusOK, "User found", user, nil)
}
