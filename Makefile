.PHONY: migrate-up, migrate-down, run, build

migrate-up:
	migrate -path migrations -database $(DATABASE_URL) up

migrate-down:
	migrate -path migrations -database $(DATABASE_URL) down 1

run:
	go run main.go

build:
	go build -o app ./main.go
