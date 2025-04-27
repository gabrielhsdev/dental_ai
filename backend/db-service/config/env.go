package config

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

// LoadEnv loads the .env on the same folder
func LoadEnv() {
	rootPath, err := filepath.Abs(".")
	if err != nil {
		log.Fatal("Failed to determine current directory:", err)
	}

	envPath := filepath.Join(rootPath, ".env")
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading .env file from %s: %v", envPath, err)
	}

	log.Println(".env loaded successfully from", envPath)
}
