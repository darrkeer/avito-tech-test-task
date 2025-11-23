#!/usr/bin/env bash
set -e

echo "Waiting for db to be available..."
./wait-for-it.sh db:5432 -t 60

docker compose exec postgres psql app app -c  'DROP DATABASE IF EXISTS test'
docker compose exec postgres psql app app -c 'CREATE DATABASE test'

echo "Running migrations..."
exec migrate -path=/migrations -database="$DATABASE_URL" up

echo "Running tests..."
go test . -v