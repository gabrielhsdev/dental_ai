package service

import (
	"github.com/gabrielhsdev/dental_ai/backend/db-service/internal/models"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/internal/repository"
	"github.com/google/uuid"
)

type UserServiceInterface interface {
	GetUserById(id uuid.UUID) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type UserService struct {
	Repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserServiceInterface {
	return &UserService{Repository: repository}
}

func (service *UserService) GetUserById(id uuid.UUID) (*models.User, error) {
	return service.Repository.GetUserById(id)
}

func (service *UserService) GetUserByEmail(email string) (*models.User, error) {
	return service.Repository.GetUserByEmail(email)
}
