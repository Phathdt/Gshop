include .env
export

.PHONY: migrate, server

help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

server: ## run server
	go run apps/web/main.go

migrate:
	go run apps/migrate/main.go $(args)
