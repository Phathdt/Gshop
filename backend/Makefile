include .env
export

.PHONY: migrate server server-dev

help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

server: ## run server
	go run cmd/server/main.go

server-dev: ## hot reload server
	~/go/bin/air -c .air.server.toml


migrate: ## migrate database
	go run cmd/migrate/main.go $(args)
