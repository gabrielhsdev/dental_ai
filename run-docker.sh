#!/bin/bash

if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

if [[ -z "$AUTH_SERVICE_HOST" || -z "$DB_SERVICE_HOST" || -z "$DIAGNOSTICS_SERVICE_HOST" ]]; then
    echo "Error: One or more required environment variables are missing."
    exit 1
fi

mkdir -p backend/"$AUTH_SERVICE_HOST"
mkdir -p backend/"$DB_SERVICE_HOST"
mkdir -p backend/"$DIAGNOSTICS_SERVICE_HOST"

cp .env backend/"$AUTH_SERVICE_HOST"/
cp .env backend/"$DB_SERVICE_HOST"/
cp .env backend/"$DIAGNOSTICS_SERVICE_HOST"/

echo ".env files copied successfully to:"
echo "- backend/$AUTH_SERVICE_HOST/"
echo "- backend/$DB_SERVICE_HOST/"
echo "- backend/$DIAGNOSTICS_SERVICE_HOST/"

echo "Running docker compose up --build"
docker compose up --build
