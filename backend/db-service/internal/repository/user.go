package repository

import (
	"database/sql"

	"github.com/gabrielhsdev/dental_ai/backend/db-service/internal/models"
	"github.com/gabrielhsdev/dental_ai/backend/db-service/pkg/database"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetUserById(id uuid.UUID) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type UserRepositoryImplementation struct {
	DB database.Database
}

func NewUserRepository(db database.Database) UserRepository {
	return &UserRepositoryImplementation{DB: db}
}

func (repository *UserRepositoryImplementation) GetUserById(id uuid.UUID) (*models.User, error) {
	var user models.User
	query := `SELECT id, userName, email, password, firstName, lastName, createdAt, updatedAt FROM users WHERE id = $1`
	err := repository.DB.QueryRow(query, id).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepositoryImplementation) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, userName, email, password, firstName, lastName, createdAt, updatedAt FROM users WHERE email = $1`
	err := repository.DB.QueryRow(query, email).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
