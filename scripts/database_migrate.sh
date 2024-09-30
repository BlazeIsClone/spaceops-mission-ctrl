#!/bin/bash

export $(grep -v '^#' .env | xargs)

DATABASE_URL="mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}"

if [ -z "$1" ]; then
  echo "Please provide a migration command (up or down)."
  exit 1
fi

migrate -database "$DATABASE_URL" -path database/migrations "$1"