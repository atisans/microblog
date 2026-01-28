# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: run quality control checks
.PHONY: audit
audit: test
	go mod tidy -diff
	go mod verify
	test -z "$(shell gofmt -l .)" 
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## upgradeable: list direct dependencies that have upgrades available
.PHONY: upgradeable
upgradeable:
	@go run github.com/oligot/go-mod-upgrade@latest

# ==================================================================================== #
# DATABASE MIGRATIONS
# ==================================================================================== #

## tidy: tidy modfiles and format .go files
.PHONY: tidy
tidy:
	sqlc generate
	go mod tidy -v
	go fmt ./...

## build: build the cmd/api application
.PHONY: build
build:
	sqlc generate
	go build -o=/tmp/bin/api ./cmd/api

## db/migrate: run all pending migrations
.PHONY: db/migrate
db/migrate:
	goose -dir migrations sqlite3 microblog.db up

## db/status: show migration status
.PHONY: db/status
db/status:
	goose -dir migrations sqlite3 microblog.db status

## db/create: create new migration
.PHONY: db/create
db/create:
	@if [ -z "$(NAME)" ]; then echo "Usage: make db/create NAME=migration_name"; exit 1; fi
	goose -dir migrations sqlite3 microblog.db create $(NAME) sql
	
## run: run the cmd/api application
.PHONY: run
run: build
	/tmp/bin/api

## run/live: run the application with reloading on file changes
.PHONY: run/live
run/live:
	go run github.com/cosmtrek/air@v1.43.0 \
		--build.cmd "make build" --build.bin "/tmp/bin/api" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"

