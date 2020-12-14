install:
	go mod tidy && go mod download

build:
	go build -o ./bin/app ./src/main.go

server: build
	./bin/app
