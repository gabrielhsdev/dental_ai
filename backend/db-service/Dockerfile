FROM golang:1.24-alpine

WORKDIR /app

RUN apk update && apk add --no-cache git curl

COPY . .

RUN go mod tidy

RUN go build -o db-service

EXPOSE ${DB_SERVICE_PORT}

CMD ["./db-service"]
