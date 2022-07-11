# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Testing
GO_TEST_FLAGS = $(VERBOSE)
COVER_PROFILE = cover.out
GOLANGCI_LINT := $(GOBIN)/golangci-lint

# Configuration
VERSION := $(shell git describe --tags --exact-match 2>/dev/null || echo latest)
MICROSERVICE_NAME := course
BINARY_NAME := sumelms-${MICROSERVICE_NAME}
RUN_FLAGS ?= 
CONFIG_PATH ?= ./config/config.yml

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

.PHONY: migrations-up
migrations-up: ## Runs the migrations 
	go run cmd/migration/main.go up $(args)

.PHONY: migrations-down
migrations-down: ## Revert the migrations
	go run cmd/migration/main.go down $(args)

.PHONY: migrations-create
migrations-create: ## Create a new migration
	go run cmd/migration/main.go create $(args)


## --------------------------------------
## Linting
## --------------------------------------

.PHONY: precommit
precommit: ## Executes the pre-commit hook (check the stashed changes)
	./.githooks/pre-commit

.PHONY: lint
lint: $(GOLANGCI_LINT) ## Executes the linting tool (vet, sec, and others)
	$(GOLANGCI_LINT) run $(RUN_FLAGS)

.PHONY: lint-fix
lint-fix: $(GOLANGCI_LINT) ## Executes the linting with fix
	$(GOLANGCI_LINT) run --fix $(RUN_FLAGS)

$(GOLANGCI_LINT):
	./.travis/install-golangci-lint.sh

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