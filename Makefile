.PHONY: docs
.PHONY: tests

install:
	go mod tidy && go mod download

build: docs
	go build -o ./bin/app ./src/main.go

dev-server:
	gin --appPort 8000 --port 8000

server: build
	./bin/app

docs:
	$$HOME/go/bin/swag init -g src/main.go