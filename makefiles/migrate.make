USERDSN=host=localhost user=$(POSTGRES_USER) port=5433 password=$(POSTGRES_PASSWORD) dbname=$(UDBNAME) sslmode=disable
UDIR=services/user-service/migrate

.PHONY: migrate-usersvc-up migrate-usersvc-down migrate-usersvc-create

migrate-usersvc-create:
    @read -p "Enter migration name: " name; \
	goose -dir $(UDIR) postgres create $$name sql

migrate-usersvc-up:
	goose -dir $(UDIR) postgres "$(USERDSN)" up

migrate-usersvc-down:
	goose -dir $(UDIR) postgres "$(USERDSN)" down
