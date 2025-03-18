package repository

import (
	"database/sql"
	"log"

	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/database"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/internal/models"
)

type UserRepository interface {
	GetUserById(id int) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type UserRepositoryImplementation struct {
	DB database.Database
}

func NewUserRepository(db database.Database) UserRepository {
	return &UserRepositoryImplementation{DB: db}
}

func (repository *UserRepositoryImplementation) GetUserById(id int) (*models.User, error) {
	var user models.User
	query := `SELECT id, userName, email, password, firstName, lastName, createdAt, updatedAt FROM users WHERE id = $1`
	err := repository.DB.GetDB().QueryRow(query, id).Scan(
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
		log.Println("Error Getting User By Id:", err)
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepositoryImplementation) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, userName, email, password, firstName, lastName, createdAt, updatedAt FROM users WHERE email = $1`
	// err := repository.DB.QueryRow(query, email).Scan(
	err := repository.DB.GetDB().QueryRow(query, email).Scan(
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
		log.Println("Error Getting User By Email:", err)
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepositoryImplementation) CreateUser(user *models.User) (*models.User, error) {
	query := `INSERT INTO users (userName, email, password, firstName, lastName, createdAt, updatedAt) 
              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := repository.DB.GetDB().QueryRow(
		query,
		user.Username,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.CreatedAt,
		user.UpdatedAt).Scan(&user.Id)

	if err != nil {
		log.Println("Error Creating User", err)
		return nil, err
	}

	return user, nil
}
