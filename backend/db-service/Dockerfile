FROM node:20-alpine

WORKDIR /app

# Install dependencies
RUN apk update && apk add --no-cache git curl && apk add postgresql-client

# Copy package.json files
COPY backend/db-service/package*.json ./

# Install npm dependencies
RUN npm install --only=production

# Copy all files to the container
COPY backend/db-service/. .
COPY .env .

# Build the TypeScript files
RUN npm run build

# Expose the application port
EXPOSE ${DB_SERVICE_PORT}

# Start the application
CMD ["node", "dist/index.js"]
