package service

import (
	"github.com/gabrielhsdev/dental_ai/backend/db-service/internal/models"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/internal/repository"
	"github.com/google/uuid"
)

type PatientsServiceInterface interface {
	GetPatientsByUserId(userId uuid.UUID) ([]*models.Patient, error)
	GetPatientById(id uuid.UUID) (*models.Patient, error)
	CreatePatient(patient *models.Patient) (*models.Patient, error)
}

type PatientService struct {
	Repository repository.PatientRepository
}

func NewPatientService(repository repository.PatientRepository) PatientsServiceInterface {
	return &PatientService{Repository: repository}
}

func (service *PatientService) GetPatientsByUserId(userId uuid.UUID) ([]*models.Patient, error) {
	return service.Repository.GetPatientsByUserId(userId)
}

func (service *PatientService) GetPatientById(id uuid.UUID) (*models.Patient, error) {
	return service.Repository.GetPatientById(id)
}

func (service *PatientService) CreatePatient(patient *models.Patient) (*models.Patient, error) {
	return service.Repository.CreatePatient(patient)
}
