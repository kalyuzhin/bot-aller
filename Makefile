OUTPUT=app
MAIN_PATH=cmd/main.go

ifneq (,$(wildcard .env))
	include .env
endif

GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)

ifeq ($(POSTGRES_SETUP_TEST),)
	POSTGRES_SETUP_TEST := user=$(DB_USERNAME) password=$(DB_PASSWORD) dbname=$(DB_NAME) host=$(DB_HOST) port=$(DB_PORT) sslmode=disable
endif

MIGRATION_FOLDER=$(CURDIR)/migrations

.PHONY: help
## prints help about all targets
help:
	@echo ""
	@echo "Usage:"
	@echo "  make <target>"
	@echo ""
	@echo "Targets:"
	@awk '                                \
		BEGIN { comment=""; }             \
		/^\s*##/ {                         \
		    comment = substr($$0, index($$0,$$2)); next; \
		}                                  \
		/^[a-zA-Z0-9_-]+:/ {               \
		    target = $$1;                  \
		    sub(":", "", target);          \
		    if (comment != "") {           \
		        printf "  %-17s %s\n", target, comment; \
		        comment="";                \
		    }                              \
		}' $(MAKEFILE_LIST)
	@echo ""

.PHONY: tidy
## runs go mod tidy
tidy:
	go mod tidy

.PHONY: fmt
## runs go fmt
fmt:
	go fmt ./...

.PHONY: lint
## runs linter
lint: fmt
	golangci-lint run -c .golangci.yaml

.PHONY: vet
## runs go vet
vet: fmt
	go vet ./...

.PHONY: build
## builds project
build: fmt
	go build -o $(OUTPUT) $(MAIN_PATH)


.PHONY: migration-create
## creates migration with first param as name
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create $(name) sql

.PHONY: migration-up
## applies latest migration
migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

.PHONY: migration-down
## rolls back latest migration
migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down

.PHONY: migration-status
## checks migration status
migration-status:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" status
