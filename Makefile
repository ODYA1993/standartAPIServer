.PHONY: build

build:
	go build -v ./cmd/apiserver/

create migrate:
	migrate create -ext sql -dir migrations UserCreationMigration

up:
	migrate -path migrations -database "postgres://localhost:5432/postgres?sslmode=disable&user=postgres&password=postgres" up

down:
	migrate path migrations -database "postgres://localhost:5432/postgres?sslmode=disable&user=postgres&password=postgres" down
.DEFAULT_GOAL := build
