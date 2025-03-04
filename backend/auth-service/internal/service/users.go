package service

import (
	"time"

	"main.go/internal/models"
	"main.go/internal/repository"
)

type UserService struct {
	Repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
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
