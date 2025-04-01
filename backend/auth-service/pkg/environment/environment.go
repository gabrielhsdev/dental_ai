package environment

import "os"

type EnvManagerInterface interface {
	GetDBHost() string
	GetDBHostDebug() string
	GetDBPort() string
	GetDBUser() string
	GetDBPassword() string
	GetDBName() string
}

// Use the exact names as the .env file for now
type EnvManager struct {
	db_host       string
	db_host_debug string
	db_port       string
	db_user       string
	db_password   string
	db_name       string
}

func NewEnvManager() EnvManagerInterface {
	return &EnvManager{
		db_host:       getEnv("DB_HOST"),
		db_host_debug: getEnv("DB_HOST_DEBUG"),
		db_port:       getEnv("DB_PORT"),
		db_user:       getEnv("DB_USER"),
		db_password:   getEnv("DB_PASSWORD"),
		db_name:       getEnv("DB_NAME"),
	}
}

func (envManager *EnvManager) GetDBHost() string {
	return envManager.db_host
}

func (envManager *EnvManager) GetDBHostDebug() string {
	return envManager.db_host_debug
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

// Aux function to get env variables
func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("Environment variable " + key + " not set")
	}
	return value
}
