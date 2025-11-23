.PHONY: start test

test_db_url = postgres://pudge:swaga1337@localhost:5432/main?sslmode=disable

start:
	docker-compose up

test:
	docker-compose up -d db
	./wait-for-it.sh localhost:5432 -t 60

	migrate -path=./migrations -database="$(test_db_url)" up

	docker compose exec db psql main pudge -c 'DROP DATABASE IF EXISTS app_test'
	docker compose exec db psql main pudge -c 'CREATE DATABASE app_test'

	env DATABASE_URL=$(test_db_url) go test -v ./...
