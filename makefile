# Project settings
PROJECT_NAME=food-pilot
PKG=./...
MAIN=main.go
BUILD_DIR=build

# Git info
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

# Docker setting
TEST_TAG=dev
DOCKER_FILE_PATH=.

# Default target
run: ## Run the app
	@clear
	@go run $(MAIN)

all: build ## Build the project (default)

compile: ## Compile (without installing)
	@go build -o $(BUILD_DIR)/$(PROJECT_NAME) $(MAIN)

build: clean compile
	./$(BUILD_DIR)/$(PROJECT_NAME)


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

image: ## build an image from docker file
	docker build -t $(PROJECT_NAME):$(TEST_TAG) $(DOCKER_FILE_PATH) 

drun: ## make dokcer run the image
	docker run --name $(PROJECT_NAME)_$(TEST_TAG) --env-file .env -p 8080:8080 -d $(PROJECT_NAME):$(TEST_TAG)

dstop: ## delete and stop continer
	docker stop $(PROJECT_NAME)_$(TEST_TAG)
	docker rm $(PROJECT_NAME)_$(TEST_TAG)
