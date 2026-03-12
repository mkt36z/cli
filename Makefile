# mkt36z CLI — Build System
# Usage: make [target]

# Build variables
VERSION   ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT    ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE      ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS    = -s -w \
             -X github.com/mkt36z/cli/internal/version.Version=$(VERSION) \
             -X github.com/mkt36z/cli/internal/version.Commit=$(COMMIT) \
             -X github.com/mkt36z/cli/internal/version.Date=$(DATE)

BINARY     = mkt36z
GOFLAGS   ?=

.PHONY: build dev test lint clean fmt vet \
        worker-dev worker-deploy worker-typecheck \
        seed-kv coverage help

## Build

build: ## Build the CLI binary with version info
	go build $(GOFLAGS) -ldflags '$(LDFLAGS)' -o bin/$(BINARY) ./cmd/mkt36z

dev: ## Run the CLI in development mode
	go run ./cmd/mkt36z $(ARGS)

install: build ## Install the binary to $GOPATH/bin
	cp bin/$(BINARY) $(shell go env GOPATH)/bin/

## Quality

test: ## Run all tests
	go test ./... -race -count=1

coverage: ## Run tests with coverage report
	go test ./... -race -coverprofile=coverage.out -covermode=atomic
	go tool cover -func=coverage.out
	@echo ""
	@echo "HTML report: go tool cover -html=coverage.out"

lint: ## Run golangci-lint
	golangci-lint run ./...

fmt: ## Format all Go files
	gofmt -s -w .

vet: ## Run go vet
	go vet ./...

check: fmt vet lint test ## Run all quality checks

## Workers

worker-dev: ## Start Workers dev server
	cd workers && npm run dev

worker-deploy: ## Deploy Workers to Cloudflare
	cd workers && npm run deploy

worker-typecheck: ## Type-check Workers TypeScript
	cd workers && npm run typecheck

worker-db-migrate: ## Apply D1 schema
	cd workers && npm run db:migrate

## Assets

seed-kv: ## Generate KV manifests from assets and seed to Workers KV
	@if [ -f scripts/generate-kv-manifests.js ]; then node scripts/generate-kv-manifests.js; fi

## Utility

clean: ## Remove build artifacts
	rm -rf bin/ coverage.out

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
