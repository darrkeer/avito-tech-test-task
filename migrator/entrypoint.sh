#!/usr/bin/env bash
set -e

echo "Waiting for db to be available..."
./wait-for-it.sh db:5432 -t 60

echo "Running migrations..."
exec migrate -path=/migrations -database="$DATABASE_URL" up
