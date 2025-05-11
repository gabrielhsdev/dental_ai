package repository

import (
	"database/sql"

	"github.com/gabrielhsdev/dental_ai/backend/db-service/internal/models"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/pkg/database"
	"github.com/google/uuid"
)

type PatientRepository interface {
	GetPatientsByUserId(userId uuid.UUID) ([]*models.Patient, error)
	GetPatientById(id uuid.UUID) (*models.Patient, error)
	CreatePatient(patient *models.Patient) (*models.Patient, error)
}

type PatientRepositoryImplementation struct {
	DB database.Database
}

func NewPatientRepository(db database.Database) PatientRepository {
	return &PatientRepositoryImplementation{DB: db}
}

func (repository *PatientRepositoryImplementation) GetPatientsByUserId(userId uuid.UUID) ([]*models.Patient, error) {
	var patients []*models.Patient
	query := `SELECT id, userId, firstName, lastName, dateOfBirth, gender, phoneNumber, email, notes, createdAt, updatedAt FROM patients WHERE userId = $1`
	rows, err := repository.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var patient models.Patient
		err := rows.Scan(
			&patient.Id,
			&patient.UserId,
			&patient.FirstName,
			&patient.LastName,
			&patient.DateOfBirth,
			&patient.Gender,
			&patient.PhoneNumber,
			&patient.Email,
			&patient.Notes,
			&patient.CreatedAt,
			&patient.UpdatedAt)
		if err != nil {
			return nil, err
		}
		patients = append(patients, &patient)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return patients, nil
}

func (repository *PatientRepositoryImplementation) GetPatientById(id uuid.UUID) (*models.Patient, error) {
	var patient models.Patient
	query := `SELECT id, userId, firstName, lastName, dateOfBirth, gender, phoneNumber, email, notes, createdAt, updatedAt FROM patients WHERE id = $1`
	err := repository.DB.QueryRow(query, id).Scan(
		&patient.Id,
		&patient.UserId,
		&patient.FirstName,
		&patient.LastName,
		&patient.DateOfBirth,
		&patient.Gender,
		&patient.PhoneNumber,
		&patient.Email,
		&patient.Notes,
		&patient.CreatedAt,
		&patient.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &patient, nil
}

func (repository *PatientRepositoryImplementation) CreatePatient(patient *models.Patient) (*models.Patient, error) {
	query := `INSERT INTO patients (userId, firstName, lastName, dateOfBirth, gender, phoneNumber, email, notes, createdAt, updatedAt)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
	err := repository.DB.QueryRow(query,
		patient.UserId,
		patient.FirstName,
		patient.LastName,
		patient.DateOfBirth,
		patient.Gender,
		patient.PhoneNumber,
		patient.Email,
		patient.Notes,
		patient.CreatedAt,
		patient.UpdatedAt).Scan(&patient.Id)
	if err != nil {
		return nil, err
	}
	return patient, nil
}
