OUTPUT=app
MAIN_PATH=cmd/main.go

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
