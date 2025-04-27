package environment

import (
	"log"
	"os"

	"github.com/gabrielhsdev/dental_ai/backend/db-service/pkg/mode"
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
	GetAuthServiceUpstream() string
	GetAuthServicePath() string
	GetAuthServiceFolder() string

	GetDBServicePort() string
	GetDBServiceHost() string
	GetDBServiceUpstream() string
	GetDBServicePath() string
	GetDBServiceFolder() string

	GetDiagnosticsServicePort() string
	GetDiagnosticsServiceHost() string
	GetDiagnosticsServiceUpstream() string
	GetDiagnosticsServicePath() string
	GetDiagnosticsServiceFolder() string

	GetJWTSecretKey() string

	GetAPIGatewayURLLocal() string
	GetAPIGatewayURLDocker() string
	GetAuthProxyLocation() string
	GetGitRepoURL() string

	GetPGAdminHost() string
	GetPGAdminPort() string
	GetPGAdminDefaultEmail() string
	GetPGAdminDefaultPassword() string
}

type EnvManager struct {
	db_host             string
	db_host_development string
	db_port             string
	db_user             string
	db_password         string
	db_name             string

	auth_service_port     string
	auth_service_host     string
	auth_service_upstream string
	auth_service_path     string
	auth_service_folder   string

	db_service_port     string
	db_service_host     string
	db_service_upstream string
	db_service_path     string
	db_service_folder   string

	diagnostics_service_port     string
	diagnostics_service_host     string
	diagnostics_service_upstream string
	diagnostics_service_path     string
	diagnostics_service_folder   string

	jwt_secret_key string

	api_gateway_url_local  string
	api_gateway_url_docker string
	auth_proxy_location    string
	git_repo_url           string

	pgadmin_host             string
	pgadmin_port             string
	pgadmin_default_email    string
	pgadmin_default_password string
}

func NewEnvManager(modeManager mode.ModeManagerInterface) EnvManagerInterface {
	loadEnvFiles(modeManager)
	return &EnvManager{
		db_host:             getEnv("DB_HOST"),
		db_host_development: getEnv("DB_HOST_DEVELOPMENT"),
		db_port:             getEnv("DB_PORT"),
		db_user:             getEnv("DB_USER"),
		db_password:         getEnv("DB_PASSWORD"),
		db_name:             getEnv("DB_NAME"),

		auth_service_port:     getEnv("AUTH_SERVICE_PORT"),
		auth_service_host:     getEnv("AUTH_SERVICE_HOST"),
		auth_service_upstream: getEnv("AUTH_SERVICE_UPSTREAM"),
		auth_service_path:     getEnv("AUTH_SERVICE_PATH"),
		auth_service_folder:   getEnv("AUTH_SERVICE_FOLDER"),

		db_service_port:     getEnv("DB_SERVICE_PORT"),
		db_service_host:     getEnv("DB_SERVICE_HOST"),
		db_service_upstream: getEnv("DB_SERVICE_UPSTREAM"),
		db_service_path:     getEnv("DB_SERVICE_PATH"),
		db_service_folder:   getEnv("DB_SERVICE_FOLDER"),

		diagnostics_service_port:     getEnv("DIAGNOSTICS_SERVICE_PORT"),
		diagnostics_service_host:     getEnv("DIAGNOSTICS_SERVICE_HOST"),
		diagnostics_service_upstream: getEnv("DIAGNOSTICS_SERVICE_UPSTREAM"),
		diagnostics_service_path:     getEnv("DIAGNOSTICS_SERVICE_PATH"),
		diagnostics_service_folder:   getEnv("DIAGNOSTICS_SERVICE_FOLDER"),

		jwt_secret_key: getEnv("JWT_SECRET_KEY"),

		api_gateway_url_local:  getEnv("API_GATEWAY_URL_LOCAL"),
		api_gateway_url_docker: getEnv("API_GATEWAY_URL_DOCKER"),
		auth_proxy_location:    getEnv("AUTH_PROXY_LOCATION"),
		git_repo_url:           getEnv("GIT_REPO_URL"),

		pgadmin_host:             getEnv("PGADMIN_HOST"),
		pgadmin_port:             getEnv("PGADMIN_PORT"),
		pgadmin_default_email:    getEnv("PGADMIN_DEFAULT_EMAIL"),
		pgadmin_default_password: getEnv("PGADMIN_DEFAULT_PASSWORD"),
	}
}

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

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("Environment variable " + key + " not set")
	}
	return value
}

// === Getter implementations ===

func (e *EnvManager) GetDBHost() string            { return e.db_host }
func (e *EnvManager) GetDBHostDevelopment() string { return e.db_host_development }
func (e *EnvManager) GetDBPort() string            { return e.db_port }
func (e *EnvManager) GetDBUser() string            { return e.db_user }
func (e *EnvManager) GetDBPassword() string        { return e.db_password }
func (e *EnvManager) GetDBName() string            { return e.db_name }

func (e *EnvManager) GetAuthServicePort() string     { return e.auth_service_port }
func (e *EnvManager) GetAuthServiceHost() string     { return e.auth_service_host }
func (e *EnvManager) GetAuthServiceUpstream() string { return e.auth_service_upstream }
func (e *EnvManager) GetAuthServicePath() string     { return e.auth_service_path }
func (e *EnvManager) GetAuthServiceFolder() string   { return e.auth_service_folder }

func (e *EnvManager) GetDBServicePort() string     { return e.db_service_port }
func (e *EnvManager) GetDBServiceHost() string     { return e.db_service_host }
func (e *EnvManager) GetDBServiceUpstream() string { return e.db_service_upstream }
func (e *EnvManager) GetDBServicePath() string     { return e.db_service_path }
func (e *EnvManager) GetDBServiceFolder() string   { return e.db_service_folder }

func (e *EnvManager) GetDiagnosticsServicePort() string     { return e.diagnostics_service_port }
func (e *EnvManager) GetDiagnosticsServiceHost() string     { return e.diagnostics_service_host }
func (e *EnvManager) GetDiagnosticsServiceUpstream() string { return e.diagnostics_service_upstream }
func (e *EnvManager) GetDiagnosticsServicePath() string     { return e.diagnostics_service_path }
func (e *EnvManager) GetDiagnosticsServiceFolder() string   { return e.diagnostics_service_folder }

func (e *EnvManager) GetJWTSecretKey() string { return e.jwt_secret_key }

func (e *EnvManager) GetAPIGatewayURLLocal() string  { return e.api_gateway_url_local }
func (e *EnvManager) GetAPIGatewayURLDocker() string { return e.api_gateway_url_docker }
func (e *EnvManager) GetAuthProxyLocation() string   { return e.auth_proxy_location }
func (e *EnvManager) GetGitRepoURL() string          { return e.git_repo_url }

func (e *EnvManager) GetPGAdminHost() string            { return e.pgadmin_host }
func (e *EnvManager) GetPGAdminPort() string            { return e.pgadmin_port }
func (e *EnvManager) GetPGAdminDefaultEmail() string    { return e.pgadmin_default_email }
func (e *EnvManager) GetPGAdminDefaultPassword() string { return e.pgadmin_default_password }
