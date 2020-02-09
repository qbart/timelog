TEST?=./...

run:
	go run main.go

start:
	go run main.go -- start hello

stop:
	go run main.go -- stop

export:
	go run main.go -- export

build:
	mkdir -p bin/
	go build -o bin/timelog

install: build
	mkdir -p $(HOME)/bin/
	cp bin/timelog $(HOME)/bin/

test:
	go test $(TEST) -coverprofile=coverage.out -timeout=2m -parallel=4

coverage: test
	go tool cover -html=coverage.out
