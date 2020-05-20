run:
	go run main.go

start:
	go run main.go -- start hello

stop:
	go run main.go -- stop

clear:
	go run main.go -- clear

adjust:
	go run main.go -- adjust

qlist:
	go run main.go -- qlist

version:
	go run main.go -- version

build:
	mkdir -p bin/
	go build -o bin/timelog

lint:
	go vet ./...

install: build
	mkdir -p $(HOME)/bin/
	cp bin/timelog $(HOME)/bin/

test:
	gotestsum

coverage: test
	go test ./... -coverprofile=coverage.out -timeout=2m -parallel=4
	go tool cover -html=coverage.out
