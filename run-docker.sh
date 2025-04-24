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
    local required_vars=(AUTH_SERVICE_FOLDER DB_SERVICE_FOLDER DIAGNOSTICS_SERVICE_FOLDER GIT_REPO_URL)
    for var in "${required_vars[@]}"; do
        if [ -z "${!var:-}" ]; then
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
    for service in AUTH_SERVICE_FOLDER DB_SERVICE_FOLDER DIAGNOSTICS_SERVICE_FOLDER; do
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

# Prepare auth service
prepare_auth_service() {
    local auth_service_dir="backend/${AUTH_SERVICE_FOLDER}"
    printf "Running Go setup commands inside %s folder\n" "$auth_service_dir"
    
    cd "$auth_service_dir" || { echo "Failed to enter $auth_service_dir directory"; exit 1; }

    if [ ! -f "go.mod" ]; then
        go mod init "${GIT_REPO_URL}/backend/${AUTH_SERVICE_HOST}"
    fi

    go mod tidy
    if [ $? -ne 0 ]; then
        printf "Error: Failed to run 'go mod tidy' command\n"
        exit 1
    fi 

    go build
    if [ $? -ne 0 ]; then
        printf "Error: Failed to run 'go build' command\n"
        exit 1
    fi

    cd ../..

    printf "Go setup commands executed successfully in %s folder\n" "$auth_service_dir"
}

# Prepare db service
prepare_db_service() {
    local db_service_dir="backend/${DB_SERVICE_FOLDER}"
    printf "Running Go setup commands inside %s folder\n" "$db_service_dir"

    cd "$db_service_dir" || { echo "Failed to enter $db_service_dir directory"; exit 1; }

    if [ ! -f "go.mod" ]; then
        go mod init "${GIT_REPO_URL}/backend/${DB_SERVICE_HOST}"
    fi

    go mod tidy
    if [ $? -ne 0 ]; then
        printf "Error: Failed to run 'go mod tidy' command\n"
        exit 1
    fi 

    go build
    if [ $? -ne 0 ]; then
        printf "Error: Failed to run 'go build' command\n"
        exit 1
    fi

    cd ../..

    printf "Go setup commands executed successfully in %s folder\n" "$db_service_dir" 
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

# Delete our /data folder on our current directory
delete_data_folder() {
    if [ -d "./data" ]; then
        printf "Deleting ./data folder\n"
        rm -rf ./data
        if [ $? -ne 0 ]; then
            printf "Error: Failed to delete ./data folder\n"
            exit 1
        fi
    else
        printf "./data folder does not exist, skipping deletion\n"
    fi
}

# Main script execution
main() {
    printf "Starting script execution...\n"
    delete_data_folder
    if [ $? -ne 0 ]; then
        printf "Error: Failed to delete data folder\n"
        exit 1
    fi

    load_env
    if [ $? -ne 0 ]; then
        printf "Error: Failed to load environment variables\n"
        exit 1
    fi

    export $(grep -v '^#' .env | xargs)
    if [ $? -ne 0 ]; then
        printf "Error: Failed to export environment variables\n"
        exit 1
    fi

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
        printf "Error: Failed to complete function prepare_auth_service\n"
        exit 1
    fi

    prepare_db_service
    if [ $? -ne 0 ]; then
        printf "Error: Failed to complete function prepare_db_service\n"
        exit 1
    fi

    run_docker
    if [ $? -ne 0 ]; then
        printf "Error: Docker commands execution failed\n"
        exit 1
    fi
}

main