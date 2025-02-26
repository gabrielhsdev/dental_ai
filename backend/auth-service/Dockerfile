FROM golang:1.24-alpine

WORKDIR /app

RUN apk update && apk add --no-cache git curl

COPY . .

RUN go mod tidy

RUN go build -o auth-service

EXPOSE ${AUTH_SERVICE_PORT}

CMD ["/bin/sh", "-c", "./auth-service -port $AUTH_SERVICE_PORT"]