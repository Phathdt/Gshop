include .env
export

.PHONY: migrate server air

help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

server: ## run server
	go run main.go

migrate: ## migrate database
	go run migrate/main.go $(args)

air: ## hot reload app
	~/go/bin/air .
