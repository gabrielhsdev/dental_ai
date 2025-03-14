FROM node:20-alpine

WORKDIR /app

RUN apk update && apk add --no-cache git curl && apk add postgresql-client

COPY backend/db-service/package*.json ./

RUN npm install --only=production

COPY backend/db-service/. .
COPY .env .

RUN npm run build

EXPOSE ${DB_SERVICE_PORT}

CMD ["node", "dist/index.js"]