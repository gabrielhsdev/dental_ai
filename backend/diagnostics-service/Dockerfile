FROM golang:1.24-alpine

WORKDIR /app

RUN apk update && apk add --no-cache git curl

COPY . .

RUN go mod tidy

RUN go build -o diagnostics-service

EXPOSE ${DIAGNOSTICS_SERVICE_PORT}

CMD ["/bin/sh", "-c", "./diagnostics-service -port $DIAGNOSTICS_SERVICE_PORT"]