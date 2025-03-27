# this file is used so we can double check if the environment variables are set
# and if the file was created with the variables replaced, we cannot use envsubst because
# it kept replacing nginx variables, so we had to do it manually

REQUIRED_ENV_VARS=(
    "AUTH_SERVICE_HOST"
    "AUTH_SERVICE_PORT"
    "DB_SERVICE_HOST"
    "DB_SERVICE_PORT"
    "DIAGNOSTICS_SERVICE_HOST"
    "DIAGNOSTICS_SERVICE_PORT"
    "PGADMIN_HOST"
    "PGADMIN_PORT"
)

for var in ${REQUIRED_ENV_VARS[@]}; do
    if [ -z "${!var}" ]; then
        echo "Error: $var is not set"
        exit 1
    fi
done

# Manually replace all the variables that have the ${VAR} format
sed -e "s|\${AUTH_SERVICE_HOST}|$AUTH_SERVICE_HOST|g" \
    -e "s|\${AUTH_SERVICE_PORT}|$AUTH_SERVICE_PORT|g" \
    -e "s|\${DB_SERVICE_HOST}|$DB_SERVICE_HOST|g" \
    -e "s|\${DB_SERVICE_PORT}|$DB_SERVICE_PORT|g" \
    -e "s|\${DIAGNOSTICS_SERVICE_HOST}|$DIAGNOSTICS_SERVICE_HOST|g" \
    -e "s|\${DIAGNOSTICS_SERVICE_PORT}|$DIAGNOSTICS_SERVICE_PORT|g" \
    -e "s|\${PGADMIN_HOST}|$PGADMIN_HOST|g" \
    -e "s|\${PGADMIN_PORT}|$PGADMIN_PORT|g" \
    /etc/nginx/nginx.conf.template > /etc/nginx/nginx.conf

# Check if the file was created
if [ ! -f /etc/nginx/nginx.conf ]; then
    echo "Error: /etc/nginx/nginx.conf was not created"
    exit 1
fi

# Check if the file is not empty
if [ ! -s /etc/nginx/nginx.conf ]; then
    echo "Error: /etc/nginx/nginx.conf is empty"
    exit 1
fi

# Check if the file has the variables replaced
if grep -q '${' /etc/nginx/nginx.conf; then
    echo "Error: /etc/nginx/nginx.conf still has variables to be replaced"
    exit 1
fi

echo "Success: /etc/nginx/nginx.conf was created"
cat /etc/nginx/nginx.conf

exit 0