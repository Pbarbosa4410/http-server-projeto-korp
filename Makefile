.PHONY: up down build test fmt logs ps compose-check docker-build

up:
	sudo docker compose up -d --build

down:
	sudo docker compose down

build:
	go build ./cmd/server

test:
	go test -v ./...

fmt:
	gofmt -w cmd/server/main.go internal/server/server.go internal/server/server_test.go

logs:
	sudo docker compose logs -f

ps:
	sudo docker compose ps

compose-check:
	sudo docker compose config

docker-build:
	sudo docker build -t http-server-projeto-korp:test .