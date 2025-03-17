package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func LoadPostgres() {
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

	log.Printf("Connecting to PostgreSQL - Host: %s, Port: %s, User: %s, DB: %s", host, port, user, dbname)

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatalf("Error: One or more database environment variables are not set")
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening PostgreSQL database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error pinging PostgreSQL database: %v", err)
	}

	log.Println("PostgreSQL database connection established successfully")
}
