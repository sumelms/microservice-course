# Testing
GO_TEST_FLAGS = $(VERBOSE)
COVER_PROFILE = cover.out
GOLANGCI_LINT := golangci-lint

# Configuration
VERSION := $(shell git describe --tags --exact-match 2>/dev/null || echo latest)
MICROSERVICE_NAME := course
BINARY_NAME := sumelms-${MICROSERVICE_NAME}
RUN_FLAGS ?= 
CONFIG_PATH ?= ./config/config.yml

# Database / Migration
SUMELMS_DATABASE_DRIVER ?= postgres
SUMELMS_DATABASE_HOST ?= localhost
SUMELMS_DATABASE_PORT ?= 5432
SUMELMS_DATABASE_USER ?= postgres
SUMELMS_DATABASE_PASSWORD ?= password
SUMELMS_DATABASE_SSL ?= disable
SUMELMS_DATABASE_DATABASE = sumelms_course
DATABASE_DSN := "${SUMELMS_DATABASE_DRIVER}://${SUMELMS_DATABASE_USER}:${SUMELMS_DATABASE_PASSWORD}@${SUMELMS_DATABASE_HOST}:${SUMELMS_DATABASE_PORT}/${SUMELMS_DATABASE_DATABASE}?sslmode=${SUMELMS_DATABASE_SSL}"

# Container
DOCKERHUB_NAMESPACE ?= sumelms
CONTAINER_NAME ?= $(BINARY_NAME)
IMAGE := ${DOCKERHUB_NAMESPACE}/${CONTAINER_NAME}:${VERSION}

# Tools
CONTAINER_RUNTIME := $(shell command -v podman 2> /dev/null || echo docker)
CONTAINER_COMPOSE_RUNTIME := $(shell command -v podman-compose 2> /dev/null || echo docker-compose)

export GO111MODULE=on
export GOFLAGS=-mod=vendor
export SUMELMS_CONFIG_PATH=$(CONFIG_PATH)

##############################################################

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: install-githooks
install-githooks: ## Install the repository githooks
	git config core.hooksPath ./.githooks

## --------------------------------------
## Build/Run
## --------------------------------------

.PHONY: run
run: ## Runs the microservice
	go run cmd/server/main.go

.PHONY: build
build: ## Generate the microservice binary
	go build -o bin/${BINARY_NAME} cmd/server/main.go

.PHONY: build-proto
build-proto: ## Compiles the protobuf
	protoc --go-grpc_out=proto --go_out=proto proto/**/*.proto

.PHONY: migration-up
migration-up: ## Runs the migrations up
	migrate -path ./db/migrations -database ${DATABASE_DSN} up

.PHONY: migration-down
migration-down: ## Runs the migrations down
	migrate -path ./db/migrations -database ${DATABASE_DSN} down

.PHONY: swagger
swagger: ## Generate Swagger Documentation
	swag init -g swagger.go -d ./internal -o ./swagger

## --------------------------------------
## Linting
## --------------------------------------

.PHONY: precommit
precommit: ## Executes the pre-commit hook (check the stashed changes)
	./.githooks/pre-commit

.PHONY: lint
lint: ## Executes the linting tool (vet, sec, and others)
	$(GOLANGCI_LINT) run $(RUN_FLAGS)

.PHONY: lint-fix
lint-fix: ## Executes the linting with fix
	$(GOLANGCI_LINT) run --fix $(RUN_FLAGS)

## --------------------------------------
## Linting
## --------------------------------------

# Run the tests
.PHONY: test
test: unit ## Run all the tests

# Run the unit tests
.PHONY: unit
unit: ## Run the unit tests
	go test ./... $(VERBOSE) -coverprofile $(COVER_PROFILE)

.PHONY: unit-cover
unit-cover:
	go test -coverprofile=$(COVER_PROFILE) $(GO_TEST_FLAGS) ./...
	go tool cover -func=$(COVER_PROFILE)

.PHONE: unit-verbose
unit-verbose:
	VERBOSE=-v make unit

## --------------------------------------
## Go Module
## --------------------------------------

.PHONY: vendor
vendor: ## Runs tiny, vendor and verify the module
	go mod tidy
	go mod vendor
	go mod verify

## --------------------------------------
## Container
## --------------------------------------

.PHONY: container-buid
container-build: ## Build the container image
	$(CONTAINER_RUNTIME) build -t $(IMAGE) .

.PHONY: container-run
container-run: container-build ## Run the container image
	$(CONTAINER_RUNTIME) run \
		--rm \
		--name $(CONTAINER_NAME) \
		-e SUMELMS_CONFIG_PATH=/config/config.yml \
		-v ./config:/config \
		-p 8080:8080 \
		$(IMAGE)

.PHONY: container-image-push
container-image-push: container-build ## Pushes the container image to the registry
	$(CONTAINER_RUNTIME) push $(IMAGE)

.PHONY: compose-up
compose-up: ## Run the container using the docker-compose file
	$(CONTAINER_COMPOSE_RUNTIME) up $(VERBOSE)

## --------------------------------------
## Checks (mostly "private" targets)
## --------------------------------------

.PHONY: check-github-token
check-github-token:
ifndef GITHUB_TOKEN
	$(error GITHUB_TOKEN is undefined)
endif