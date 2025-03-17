package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/service"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/utils"
	"github.com/gin-gonic/gin"
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
		utils.SendResponse(context, http.StatusBadRequest, "Invalid Id", nil, err)
		return
	}

	user, err := handler.Service.GetUserById(id)
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
