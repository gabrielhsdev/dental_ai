FROM golang:1.24-alpine

WORKDIR /app

RUN apk update && apk add --no-cache git curl && apk add postgresql-client

COPY backend/auth-service/. .
COPY .env .

RUN go mod tidy

RUN go build -o auth-service

EXPOSE ${AUTH_SERVICE_PORT}

CMD ["/bin/sh", "-c", "./auth-service -port $AUTH_SERVICE_PORT"]