include app.env
export

build:
	go build -o ./bin/wallet

run: build
	./bin/wallet

start-middleware:
	docker compose up -d
.PHONY: start-middleware

stop-middleware: ### Down docker-compose
	docker compose down --remove-orphans
.PHONY: stop-middleware

