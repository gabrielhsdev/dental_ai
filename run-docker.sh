#!/bin/bash

set -euo pipefail

# Load environment variables from .env file
load_env() {
    if [ -f .env ]; then
        source .env
    else
        printf "Error: .env file not found\n"
        exit 1
    fi
}

# Check required environment variables
check_env_vars() {
    local missing_vars=()
    local required_vars=(AUTH_SERVICE_HOST DB_SERVICE_HOST DIAGNOSTICS_SERVICE_HOST AUTH_SERVICE_REPO_URL DB_SERVICE_REPO_URL)
    for var in "${required_vars[@]}"; do
        if [ -z "${!var}" ]; then
            missing_vars+=("$var")
        fi
    done

    if [ ${#missing_vars[@]} -ne 0 ]; then
        printf "Error: The following required environment variables are missing: %s\n" "${missing_vars[*]}"
        exit 1
    fi
}

# Copy .env file to each service directory
copy_env_files() {
    for service in AUTH_SERVICE_HOST DB_SERVICE_HOST DIAGNOSTICS_SERVICE_HOST; do
        local dir="backend/${!service}"
        if [ ! -d "$dir" ]; then
            printf "Error: Directory %s does not exist\n" "$dir"
            exit 1
        fi
        cp .env "$dir/"
        if [ $? -ne 0 ]; then
            printf "Error: Failed to copy .env file to %s\n" "$dir"
            exit 1
        fi
        printf ".env file copied successfully to %s\n" "$dir"
    done
}

# Execute go commands inside AUTH_SERVICE_HOST container
prepare_auth_service() {
    local auth_service_dir="backend/${AUTH_SERVICE_HOST}"
    printf "Running Go setup commands inside %s folder\n" "$auth_service_dir"
    
    cd "$auth_service_dir" || { echo "Failed to enter $auth_service_dir directory"; exit 1; }

    if [ ! -f "go.mod" ]; then
        REPO_URL="${GIT_REPO_URL}/tree/main/backend/${AUTH_SERVICE_HOST}"
        go mod init "$REPO_URL"
    fi

    go mod tidy
    go build
    if [ $? -ne 0 ]; then
        printf "Error: Failed to run 'go build' command\n"
        exit 1
    fi

    printf "Go setup commands executed successfully for %s\n" "$AUTH_SERVICE_HOST"
}

# Run Docker commands
run_docker() {
    printf "Running docker compose down -v, deleting volumes\n"
    docker compose down -v
    if [ $? -ne 0 ]; then
        printf "Error: Failed to run 'docker compose down -v'\n"
        exit 1
    fi

    printf "Running docker compose up --build\n"
    docker compose up --build
    if [ $? -ne 0 ]; then
        printf "Error: Failed to run 'docker compose up --build'\n"
        exit 1
    fi
}

# Main script execution
main() {
    load_env
    if [ $? -ne 0 ]; then
        printf "Error: Failed to load environment variables\n"
        exit 1
    fi

    export $(grep -v '^#' .env | xargs)

    check_env_vars
    if [ $? -ne 0 ]; then
        printf "Error: Environment variable check failed\n"
        exit 1
    fi

    copy_env_files
    if [ $? -ne 0 ]; then
        printf "Error: Failed to copy .env files to service directories\n"
        exit 1
    fi

    prepare_auth_service
    if [ $? -ne 0 ]; then
        printf "Error: Failed to prepare %s service\n" "$AUTH_SERVICE_HOST"
        exit 1
    fi

    run_docker
    if [ $? -ne 0 ]; then
        printf "Error: Docker commands execution failed\n"
        exit 1
    fi
}

main