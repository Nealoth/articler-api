.PHONY: clean build migrate_up migrate_down run lint default

-include ./conf/.env

db_conn_str = "postgres://${DB_HOST}:${DB_PORT}/${DB_NAME}?user=${DB_USER}&password=${DB_PASSWORD}&sslmode=${DB_SSLMODE}"

clean:
	rm -rf ./build/

build:
	go build -v -race -o ./build/articler-api ./cmd/articlerapi
	cp -R ./conf ./build

migrate_up:
	migrate -path ./migrations -database ${db_conn_str} up

migrate_down:
	migrate -path ./migrations -database ${db_conn_str} down

run:
	./build/articler-api

lint:
	golangci-lint run -c .golangci.yml

default: lint clean build run

.DEFAULT_GOAL := default