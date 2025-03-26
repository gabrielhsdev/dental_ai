package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
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

func LoadDatabase(dbType string) (Database, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	host := os.Getenv("DB_HOST")
	if len(os.Args) > 1 && os.Args[1] == "debug" {
		host = os.Getenv("DB_HOST_DEBUG")
	}
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

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
