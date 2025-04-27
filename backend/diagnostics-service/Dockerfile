# MOCK for our ai stuff
FROM node:16-slim

WORKDIR /app

ENV DIAGNOSTICS_SERVICE_PORT=8080

RUN npm install --global http-server

EXPOSE ${DIAGNOSTICS_SERVICE_PORT}

CMD ["http-server", "-p", "${DIAGNOSTICS_SERVICE_PORT}"]
