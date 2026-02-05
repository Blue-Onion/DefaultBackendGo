run:
	go run cmd/api/main.go 

build:
	go build -o app

start: build
	./app