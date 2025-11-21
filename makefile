# Project settings
BINARY_NAME=food-pilot
PKG=./...
MAIN=main.go
BUILD_DIR=build

# Git info
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)


# Default target
run: ## Run the app
	@clear
	@go run $(MAIN)

all: build ## Build the project (default)

compile: ## Compile (without installing)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN)

build: clean compile
	./$(BUILD_DIR)/$(BINARY_NAME)


test: ## Run tests
	go test -v $(PKG)

dep: ## Download and vendor dependencies
	go mod tidy
	go mod vendor

swag: ## Generate Swagger docs (requires github.com/swaggo/swag)
	@which swag >/dev/null 2>&1 || (echo "Installing swag..." && go install github.com/swaggo/swag/cmd/swag@latest)
	swag init -g internal/web/setup.go -o docs
	@echo "Swagger docs generated in ./docs"

clean: ## Remove build artifacts
	rm -rf $(BUILD_DIR)

push: ## Push current branch to origin
	git push origin $(BRANCH)

pull: ## Pull current branch from origin
	git pull origin $(BRANCH)

# lint: ## Lint the project (requires golangci-lint)
# 	golangci-lint run ./...

db: ## make db up
	@echo "try to turn on database..."
	@docker compose up -d    

db_off: ## make db down
	@echo "try to turn off database..."
	@docker compose down                                  

info: ## Show Current branch
	@echo "Branch:   $(BRANCH)"

help: ## Show this help
	@echo "Usage: make [target]"
	@awk 'BEGIN {FS = ":.*##"; printf "\nAvailable targets:\n"} /^[a-zA-Z0-9_-]+:.*##/ {printf "  %-12s %s\n", $$1, $$2}' $(MAKEFILE_LIST)
