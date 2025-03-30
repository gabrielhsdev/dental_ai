# Create a list of our current environment variables
REQUIRED_ENV_VARS=($(env | awk -F= '{print $1}'))

# Get all variables from the template, they are int the form ${VAR}
TEMPLATE_VARS=($(grep -oP '\$\{\K[^}]+(?=\})' /etc/nginx/nginx.conf.template))

# Check if all required environment variables are set
for var in ${TEMPLATE_VARS[@]}; do
    if [[ ! " ${REQUIRED_ENV_VARS[@]} " =~ " ${var} " ]]; then
        echo "Error: Required environment variable $var is not set"
        exit 1
    fi
done

# Loop through REQUIRED_ENV_VARS and replace variables in the template
cp /etc/nginx/nginx.conf.template /etc/nginx/nginx.conf
for var in ${REQUIRED_ENV_VARS[@]}; do
    sed -i "s|\${$var}|${!var}|g" /etc/nginx/nginx.conf
done

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