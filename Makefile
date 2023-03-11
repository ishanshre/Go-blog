#!make
include .env

DB_URL=postgresql://${M_DB_USERNAME}:${M_DB_PASSWORD}@localhost:5432/${M_DB_NAME}?sslmode=disable


migrateUp:
	migrate -path api/v1/db/migration -database "${DB_URL}" -verbose up

migrateDown:
	migrate -path api/v1/db/migration -database "${DB_URL}" -verbose down

dockerBuild:
	docker-compose up --build

dockerUp:
	docker-compose up

dockerDown:
	docker-compose down

dockerLog:
	docker-compose logs

run:
	go build main.go && ./main