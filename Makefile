TEST?=./...

run:
	go run main.go

start:
	go run main.go -- start hello

stop:
	go run main.go -- stop

export:
	go run main.go -- export

test:
	go test $(TEST) -coverprofile=coverage.out -timeout=2m -parallel=4

coverage: test
	go tool cover -html=coverage.out
