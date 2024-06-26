## help: show all available tasks
.PHONY: help
help: 
	@echo 'Usage: '
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## clean: deletes the ./bin and ./ui/build directories
.PHONY: clean
clean:
	@rm -rf ./bin/
	@rm -rf ./ui/build

## run/api: run the json api
.PHONY: run/cli
run/api:
	@go run ./cmd/api/main.go

## run/cli: run the command line utility
.PHONY: run/cli
run/cli:
	@go run ./cmd/cli/main.go

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

## install: [Run this on the Pi] builds and installs the tools and services (run w/ sudo). once installed, use make update to keep them up-to-date.
.PHONY: install
install: build
	@./scripts/install.sh

## update: [Run this on the Pi] builds and updates the tools at /opt/tools (run w/ sudo)
.PHONY: update
update:
	@./scripts/update.sh

## deploy/fe: builds and deploys the front end via ssh/scp
.PHONY: deploy/fe
deploy/fe:
	@if [ -z "${target}" ]; then echo "Usage: make deploy/fe target=<user@host>"; exit 1; fi
	@./scripts/deploy.fe.sh ${target}