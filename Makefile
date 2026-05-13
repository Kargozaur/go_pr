include makefiles/migrate.make
include .env

.PHONY: up down tidy run-user run-order create_db

up:
	docker-compose up -d

down:
	docker-compose down

tidy:
	cd pkg && go mod tidy
	cd services/user-service && go mod tidy
	cd services/order-service && go mod tidy
	cd services/gateway && go mod tidy

run-order:
	cd services/order-service && go run ./cmd/main.go

run-user:
	cd services/user-service && go run ./cmd/main.go

create_db:
	docker exec -it postgres bash ./docker/entrypoint.sh
