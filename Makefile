## help: show all available tasks
.PHONY: help
help: 
	@echo 'Usage: '
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## clean: deletes the ./bin directory
.PHONY: clean
clean:
	@rm -rf ./bin/

## run/api: run the json api
.PHONY: run/cli
run/api:
	@go run -race ./cmd/api/main.go

## run/cli: run the command line utility
.PHONY: run/cli
run/cli:
	@go run -race ./cmd/cli/main.go

## build: builds both the json api and command line applications
.PHONY: build
build: build/api build/cli

## build/api: builds the json api into ./bin
.PHONY: build/api
build/api:
	@go build -o=./bin/arcade-tools-api ./cmd/api/main.go

## build/cli: builds the command line utility into ./bin
.PHONY: build/cli
build/cli:
	@go build -o=./bin/arcade-tools ./cmd/cli/main.go

## install: builds and installs the tools and services (run w/ sudo). once installed, use make update to keep them up-to-date.
.PHONY: install
install: build
	@./scripts/install.sh

## update: builds and updates the tools at /opt/tools (run w/ sudo)
.PHONY: update
update: build
	@./scripts/update.sh