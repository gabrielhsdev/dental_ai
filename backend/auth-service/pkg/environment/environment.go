package environment

import (
	"log"
	"os"

	"github.com/gabrielhsdev/dental_ai/tree/main/backend/auth-service/pkg/mode"
	"github.com/joho/godotenv"
)

type EnvManagerInterface interface {
	GetDBHost() string
	GetDBHostDevelopment() string
	GetDBPort() string
	GetDBUser() string
	GetDBPassword() string
	GetDBName() string
	GetAuthServicePort() string
	GetAuthServiceHost() string
	GetJWTSecretKey() string
	GetAPIGatewayURLLocal() string
	GetAPIGatewayURLDocker() string
}

// Use the exact names as the .env file for now
type EnvManager struct {
	db_host                string
	db_host_development    string
	db_port                string
	db_user                string
	db_password            string
	db_name                string
	auth_service_port      string
	auth_service_host      string
	jwt_secret_key         string
	api_gateway_url_local  string
	api_gateway_url_docker string
}

func NewEnvManager(modeManager mode.ModeManagerInterface) EnvManagerInterface {
	loadEnvFiles(modeManager)
	return &EnvManager{
		db_host:                getEnv("DB_HOST"),
		db_host_development:    getEnv("DB_HOST_DEVELOPMENT"),
		db_port:                getEnv("DB_PORT"),
		db_user:                getEnv("DB_USER"),
		db_password:            getEnv("DB_PASSWORD"),
		db_name:                getEnv("DB_NAME"),
		auth_service_port:      getEnv("AUTH_SERVICE_PORT"),
		auth_service_host:      getEnv("AUTH_SERVICE_HOST"),
		jwt_secret_key:         getEnv("JWT_SECRET_KEY"),
		api_gateway_url_local:  getEnv("API_GATEWAY_URL_LOCAL"),
		api_gateway_url_docker: getEnv("API_GATEWAY_URL_DOCKER"),
	}
}

// Load environment files, prioritizing `.env` if it exists
func loadEnvFiles(modeManager mode.ModeManagerInterface) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	if modeManager.IsDevelopment() {
		err = godotenv.Overload(".env.development")
		if err != nil {
			log.Println("No .env.dev file found, skipping override")
		}
	}
}

func (envManager *EnvManager) GetAPIGatewayURLLocal() string {
	return envManager.api_gateway_url_local
}

func (envManager *EnvManager) GetAPIGatewayURLDocker() string {
	return envManager.api_gateway_url_docker
}

func (envManager *EnvManager) GetJWTSecretKey() string {
	return envManager.jwt_secret_key
}

func (envManager *EnvManager) GetAuthServicePort() string {
	return envManager.auth_service_port
}

func (envManager *EnvManager) GetAuthServiceHost() string {
	return envManager.auth_service_host
}

func (envManager *EnvManager) GetDBHost() string {
	return envManager.db_host
}

func (envManager *EnvManager) GetDBHostDevelopment() string {
	return envManager.db_host_development
}

func (envManager *EnvManager) GetDBPort() string {
	return envManager.db_port
}

func (envManager *EnvManager) GetDBUser() string {
	return envManager.db_user
}

func (envManager *EnvManager) GetDBPassword() string {
	return envManager.db_password
}

func (envManager *EnvManager) GetDBName() string {
	return envManager.db_name
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("Environment variable " + key + " not set")
	}
	return value
}
