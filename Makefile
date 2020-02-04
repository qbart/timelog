TEST?=./...

test:
	go test $(TEST) -timeout=2m -parallel=4

run:
	go run -race main.go
