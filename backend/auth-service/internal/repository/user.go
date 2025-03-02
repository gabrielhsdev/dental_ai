package repository

import (
	"database/sql"
	"log"

	"main.go/internal/models"
)

type UserRepository interface {
	GetUserById(id int) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
}

type UserRepositoryImplementation struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImplementation{DB: db}
}

func (repository *UserRepositoryImplementation) GetUserById(id int) (*models.User, error) {
	var user models.User
	query := `SELECT id, userName, email, firstName, lastName, createdAt, updatedAt FROM users WHERE id = $1`
	err := repository.DB.QueryRow(query, id).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
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

func (repository *UserRepositoryImplementation) CreateUser(user *models.User) (*models.User, error) {
	query := `INSERT INTO users (userName, email, password, firstName, lastName, createdAt, updatedAt) 
              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := repository.DB.QueryRow(
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
