.PHONY: migrate-up, migrate-down, run, build

run:
	go run main.go $(args)

build:
	go build -o app ./main.go

migrate-up:
	go run main.go migrate up

migrate-down:
	go run main.go migrate down
