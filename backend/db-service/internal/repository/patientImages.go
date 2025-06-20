package repository

import (
	"github.com/gabrielhsdev/dental_ai/backend/db-service/internal/models"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/pkg/database"
	"github.com/google/uuid"
)

type PatientImagesRepository interface {
	Create(patientImage *models.PatientImages) (*models.PatientImages, error)
	GetByPatientId(patientId uuid.UUID) ([]*models.PatientImages, error)
	GetById(id uuid.UUID) (*models.PatientImages, error)
	GetByUserId(userId uuid.UUID) ([]*models.PatientImages, error)
}

type PatientImagesRepositoryImplementation struct {
	DB database.Database
}

func NewPatientImagesRepository(db database.Database) PatientImagesRepository {
	return &PatientImagesRepositoryImplementation{DB: db}
}

func (repository *PatientImagesRepositoryImplementation) Create(patientImage *models.PatientImages) (*models.PatientImages, error) {
	query := `INSERT INTO patient_images (patientId, imageData, fileType, description, uploadedAt, createdAt, updatedAt, inferenceData)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err := repository.DB.QueryRow(query,
		patientImage.PatientId,
		patientImage.ImageData,
		patientImage.FileType,
		patientImage.Description,
		patientImage.UploadedAt,
		patientImage.CreatedAt,
		patientImage.UpdatedAt,
		patientImage.InferenceData).Scan(&patientImage.Id)
	if err != nil {
		return nil, err
	}
	return patientImage, nil
}

func (repository *PatientImagesRepositoryImplementation) GetByPatientId(patientId uuid.UUID) ([]*models.PatientImages, error) {
	query := `SELECT id, patientId, imageData, fileType, description, uploadedAt, createdAt, updatedAt, inferenceData FROM patient_images WHERE patientId = $1`
	rows, err := repository.DB.Query(query, patientId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patientImages []*models.PatientImages
	for rows.Next() {
		var patientImage models.PatientImages
		err := rows.Scan(
			&patientImage.Id,
			&patientImage.PatientId,
			&patientImage.ImageData,
			&patientImage.FileType,
			&patientImage.Description,
			&patientImage.UploadedAt,
			&patientImage.CreatedAt,
			&patientImage.UpdatedAt,
			&patientImage.InferenceData,
		)
		if err != nil {
			return nil, err
		}
		patientImages = append(patientImages, &patientImage)
	}
	return patientImages, nil
}

func (repository *PatientImagesRepositoryImplementation) GetById(id uuid.UUID) (*models.PatientImages, error) {
	query := `SELECT id, patientId, imageData, fileType, description, uploadedAt, createdAt, updatedAt, inferenceData FROM patient_images WHERE id = $1`
	var patientImage models.PatientImages
	err := repository.DB.QueryRow(query, id).Scan(
		&patientImage.Id,
		&patientImage.PatientId,
		&patientImage.ImageData,
		&patientImage.FileType,
		&patientImage.Description,
		&patientImage.UploadedAt,
		&patientImage.CreatedAt,
		&patientImage.UpdatedAt,
		&patientImage.InferenceData,
	)
	if err != nil {
		return nil, err
	}
	return &patientImage, nil
}

func (repository *PatientImagesRepositoryImplementation) GetByUserId(userId uuid.UUID) ([]*models.PatientImages, error) {
	query := `
	SELECT 
		pi.id,
		pi.patientId,
		pi.imageData,
		pi.fileType,
		pi.description,
		pi.uploadedAt,
		pi.createdAt,
		pi.updatedAt,
		pi.inferenceData
	FROM patient_images pi
	INNER JOIN patients p ON pi.patientId = p.id
	INNER JOIN users u ON p.userId = u.id
	WHERE u.id = $1`
	rows, err := repository.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patientImages []*models.PatientImages
	for rows.Next() {
		var patientImage models.PatientImages
		err := rows.Scan(
			&patientImage.Id,
			&patientImage.PatientId,
			&patientImage.ImageData,
			&patientImage.FileType,
			&patientImage.Description,
			&patientImage.UploadedAt,
			&patientImage.CreatedAt,
			&patientImage.UpdatedAt,
			&patientImage.InferenceData,
			// Add fields for patient first and last name if needed
		)
		if err != nil {
			return nil, err
		}
		patientImages = append(patientImages, &patientImage)
	}
	return patientImages, nil
}
