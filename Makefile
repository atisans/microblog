.PHONY: pre-commit dev backend frontend lint format

pre-commit:
	@echo "Running pre-commit checks..."
	@npm run format -w web
	@npm run format -w api

dev:
	@echo "Starting backend and frontend in parallel..."
	@npm run dev -w api & npm run dev -w web & wait

backend:
	@echo "Starting backend..."
	@npm run dev -w api

frontend:
	@echo "Starting frontend..."
	@npm run dev -w web

lint:
	@echo "Running lint checks..."
	@npm run lint -w web
	@npm run lint -w api

format:
	@echo "Running format checks..."
	@npm run format -w web
	@npm run format -w api