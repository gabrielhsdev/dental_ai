FROM golang:1.24-alpine

WORKDIR /app

RUN apk update && apk add --no-cache git curl

COPY . .

RUN go mod tidy

RUN go build -o auth-service

EXPOSE 8081

CMD ["./auth-service"]