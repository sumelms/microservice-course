# GoLang Commands

GOCMD := go
GORUN := ${GOCMD} run
GOBUILD := ${GOCMD} build
GOCLEAN := ${GOCMD} clean
GOTEST := ${GOCMD} test -v -race
GOGET := ${GOCMD} get
GOFMT := gofmt
LINTER := golangci-lint

# Project configuration

VERSION := $(shell git describe --tags --exact-match 2>/dev/null || echo latest)
DOCKERHUB_NAMESPACE ?= sumelms
MICROSERVICE_NAME := course
BINARY_NAME := sumelms-${MICROSERVICE_NAME}
IMAGE := ${DOCKERHUB_NAMESPACE}/microservice-${MICROSERVICE_NAME}:${VERSION}

##############################################################

all: test build

# Runner

run:
	export SUMELMS_CONFIG_PATH="./config/config.yml" && \
	${GORUN} cmd/server/main.go
.PHONY: run

# Builders

build: build-proto
	${GOBUILD} -o bin/${BINARY_NAME} cmd/server/main.go
.PHONY: build

build-proto:
	protoc proto/**/*.proto --go_out=plugins=grpc:.
.PHONY: build-proto

# Quality tools

test:
	${GOTEST} $$(go list ./... | grep -v /test/) $(TEST_OPTIONS)
.PHONY: test

lint:
	${LINTER} run

format:
	${GOFMT} -d .

# Docker stuff

docker-build:
	docker build -t ${IMAGE} .

docker-push: docker-build
	docker push ${IMAGE}

docker-run: docker-build
	docker run -p 8080:8080 ${IMAGE}