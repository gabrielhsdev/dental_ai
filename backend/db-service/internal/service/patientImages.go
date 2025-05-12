package service

import (
	"github.com/gabrielhsdev/dental_ai/backend/db-service/internal/models"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/internal/repository"
	"github.com/google/uuid"
)

type PatientImagesServiceInterface interface {
	Create(patientImage *models.PatientImages) (*models.PatientImages, error)
	GetByPatientId(patientId uuid.UUID) ([]*models.PatientImages, error)
	GetById(id uuid.UUID) (*models.PatientImages, error)
}

type PatientImagesService struct {
	Repository repository.PatientImagesRepository
}

func NewPatientImagesService(repository repository.PatientImagesRepository) PatientImagesServiceInterface {
	return &PatientImagesService{Repository: repository}
}

func (service *PatientImagesService) Create(patientImage *models.PatientImages) (*models.PatientImages, error) {
	return service.Repository.Create(patientImage)
}

func (service *PatientImagesService) GetByPatientId(patientId uuid.UUID) ([]*models.PatientImages, error) {
	return service.Repository.GetByPatientId(patientId)
}

func (service *PatientImagesService) GetById(id uuid.UUID) (*models.PatientImages, error) {
	return service.Repository.GetById(id)
}
