.PHONY: pre-commit dev backend frontend lint format db/migrate db/create db/status db/up db/down db/reset

pre-commit:
	@echo "Running pre-commit checks..."
	@npm run format -w web
	@npm run format -w api

dev:
	@echo "Starting backend and frontend in parallel..."
	@npm run dev -w api & npm run dev -w web & wait

backend:
	@echo "Starting backend with live reload..."
	@cd codebase && make run/live

frontend:
	@echo "Starting frontend..."
	@npm run dev -w web

lint:
	@echo "Running lint checks..."
	@cd codebase && make audit
	@npm run lint -w web

format:
	@echo "Running format checks..."
	@cd codebase && make tidy
	@npm run format -w web

db/migrate:
	@echo "Running database migrations..."
	@cd codebase && goose -dir migrations sqlite3 microblog.db up

db/create:
	@if [ -z "$(NAME)" ]; then echo "Usage: make db/create NAME=migration_name"; exit 1; fi
	@cd codebase && goose -dir migrations sqlite3 microblog.db create $(NAME) sql

db/status:
	@echo "Checking migration status..."
	@cd codebase && goose -dir migrations sqlite3 microblog.db status

db/up:
	@echo "Migrating up by one..."
	@cd codebase && goose -dir migrations sqlite3 microblog.db up-by-one

db/down:
	@echo "Migrating down by one..."
	@cd codebase && goose -dir migrations sqlite3 microblog.db down

db/reset:
	@echo "Resetting all migrations..."
	@cd codebase && goose -dir migrations sqlite3 microblog.db down-to 0