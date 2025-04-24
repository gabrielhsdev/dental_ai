FROM golang:1.24-alpine

WORKDIR /app

RUN apk update && apk add --no-cache git curl && apk add postgresql-client

ARG DB_SERVICE_FOLDER

COPY backend/${DB_SERVICE_FOLDER}/. .
COPY .env .

RUN go mod tidy

RUN go build -o app

EXPOSE ${DB_SERVICE_PORT}

CMD ["/bin/sh", "-c", "./app"]
