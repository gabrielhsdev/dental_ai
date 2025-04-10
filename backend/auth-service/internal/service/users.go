package service

import (
	"time"

	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/models"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/repository"
)

type UserServiceInterface interface {
	GetUserById(id int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	RegisterUser(user *models.User) (*models.User, error)
}

type UserService struct {
	Repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserServiceInterface {
	return &UserService{Repository: repository}
}

func (service *UserService) GetUserById(id int) (*models.User, error) {
	return service.Repository.GetUserById(id)
}

func (service *UserService) GetUserByEmail(email string) (*models.User, error) {
	return service.Repository.GetUserByEmail(email)
}

func (service *UserService) RegisterUser(user *models.User) (*models.User, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return service.Repository.CreateUser(user)
}
