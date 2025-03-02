package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/internal/models"
	"main.go/internal/service"
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
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id"})
		return
	}

	user, err := handler.Service.GetUserById(id)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}
	if user == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	context.JSON(http.StatusOK, user)
}

func (handler *UserHandler) RegisterUser(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	registeredUser, err := handler.Service.RegisterUser(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	context.JSON(http.StatusOK, registeredUser)
}
