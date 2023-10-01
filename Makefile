include .env.example

build:
	docker compose build --no-cache

start:
	docker compose up -d

test:
	MONGO_DSN=${MONGO_DSN} MONGO_DB=${MONGO_DB} go test -v ./...