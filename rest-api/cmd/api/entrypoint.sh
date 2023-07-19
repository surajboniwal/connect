#!/bin/bash -e

APP_ENV=${APP_ENV:-development}

echo "[`date`] Running entrypoint script in the '${APP_ENV}' environment..."

echo "[`date`] Loading config..."
if [ "$APP_ENV" = "development" ]; then
    source ../../internal/config/${APP_ENV}.env
else
    source production.env
fi

echo "[`date`] Running DB migrations..."
# migrate -database "${DB_URL}" -path ./migrations up

echo "[`date`] Starting $APP_ENV server..."

if [ "$APP_ENV" = "development" ]; then
    go run .
else
    ./api
fi