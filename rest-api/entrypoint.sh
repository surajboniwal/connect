#!/bin/bash -e

APP_ENV=${ENV:-development}

echo "[`date`] Running entrypoint script in the '${APP_ENV}' environment"

echo "[`date`] Loading env file"

ENV_PATH=""

if [ -f "$APP_ENV.env" ]; then  
    ENV_PATH="$APP_ENV.env"
elif [ -f "internal/config/$APP_ENV.env" ]; then
    ENV_PATH="internal/config/$APP_ENV.env"
else
    echo "Unable to find env file"
    exit 1
fi

echo "[`date`] Sourcing $ENV_PATH"
source $ENV_PATH

echo "[`date`] Running DB migrations"
migrate -database "${DB_URL}" -path migrations up > /dev/null 2>&1

echo "[`date`] Starting $APP_ENV server"

if [ "$APP_ENV" = "development" ]; then
    clear
    ENV=$APP_ENV go run ./cmd/api
else
    rm -rf ./main
    mkdir -p logs
    go build ./cmd/api/*
    clear
    ENV=$APP_ENV ./main > "logs/`date`_logs.txt"
fi