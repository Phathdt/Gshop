include .env
export

.PHONY: migrate-up, migrate-down, run, build, air


help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

server: ## run server
	go run main.go

run:
	go run main.go $(args)

build:
	go build -o app ./main.go

migrate-up:
	go run main.go migrate up

migrate-down:
	go run main.go migrate down

air:
	./bin/air
