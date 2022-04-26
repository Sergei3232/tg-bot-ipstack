.PHONY: run
run:
	go run cmd/tg-bot-ipstack/main.go


.PHONY: build
build:
	go build -o bot cmd/tg-bot-ipstack/main.go