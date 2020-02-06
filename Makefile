TEST?=./...

run:
	go run main.go -- $(ARGS)

test:
	go test $(TEST) -timeout=2m -parallel=4
