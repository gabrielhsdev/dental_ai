package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/pkg/environment"
	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/pkg/mode"
	_ "github.com/lib/pq"
)

type Database interface {
	CloseDB() error
	QueryRow(query string, args ...interface{}) *sql.Row
}

type SQLDatabase struct {
	DB *sql.DB
}

func (database *SQLDatabase) CloseDB() error {
	return database.DB.Close()
}

func (database *SQLDatabase) QueryRow(query string, args ...interface{}) *sql.Row {
	/*
	* We can do some formatting of the query row here if we change from postgres to another database
	* For Example: MySQL uses ? instead of $1 for query parameters
	* We can also log the query here for debugging purposes
	 */
	return database.DB.QueryRow(query, args...)
}

func LoadDatabase(dbType string, modeManager mode.ModeManagerInterface, envManager environment.EnvManagerInterface) (Database, error) {
	host := envManager.GetDBHost()
	if modeManager.IsDevelopment() {
		host = envManager.GetDBHostDevelopment()
	}
	port := envManager.GetDBPort()
	user := envManager.GetDBUser()
	password := envManager.GetDBPassword()
	dbname := envManager.GetDBName()

	log.Printf("Connecting database - Host: %s, Port: %s, User: %s, DB: %s", host, port, user, dbname)
	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatalf("Error: One or more database environment variables are not set")
	}

	var connStr string
	var driverName string

	switch dbType {
	case "postgres":
		driverName = "postgres"
		connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	default:
		return nil, fmt.Errorf("error: Unsupported database type")
	}

	db, err := sql.Open(driverName, connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	log.Println("Database connection established successfully")
	return &SQLDatabase{DB: db}, nil
}
