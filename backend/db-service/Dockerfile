FROM node:20-alpine

WORKDIR /app

RUN apk update && apk add --no-cache git curl && apk add postgresql-client

ARG DB_SERVICE_HOST

ENV DB_SERVICE_HOST=${DB_SERVICE_HOST}

COPY backend/${DB_SERVICE_HOST}/package*.json ./

RUN npm install --only=production

COPY backend/${DB_SERVICE_HOST}/. .
COPY .env .

RUN npm run build

EXPOSE ${DB_SERVICE_PORT}

CMD ["node", "dist/index.js"]
