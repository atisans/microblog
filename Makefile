.PHONY: pre-commit dev backend frontend lint format db:migrate db:generate db:push db:studio db:drop

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

db:migrate:
	@echo "Running database migrations..."
	@npm run db:migrate -w api

db:generate:
	@echo "Generating database code with sqlc..."
	@sqlc generate

db:push:
	@echo "Pushing database schema..."
	@npm run db:push -w api

db:studio:
	@echo "Opening database studio..."
	@npm run db:studio -w api

db:drop:
	@echo "Dropping database..."
	@npm run db:drop -w api