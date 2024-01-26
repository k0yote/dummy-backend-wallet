build:
	go build -o ./bin/wallet

run: build
	./bin/wallet