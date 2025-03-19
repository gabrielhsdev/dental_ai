FROM golang:1.24-alpine

WORKDIR /app

RUN apk update && apk add --no-cache git curl && apk add postgresql-client

ARG AUTH_SERVICE_HOST

ENV AUTH_SERVICE_HOST=${AUTH_SERVICE_HOST}

COPY backend/${AUTH_SERVICE_HOST}/. .
COPY .env .

RUN go mod tidy

RUN go build -o ${AUTH_SERVICE_HOST}

EXPOSE ${AUTH_SERVICE_PORT}

CMD ["/bin/sh", "-c", "./${AUTH_SERVICE_HOST}"]
