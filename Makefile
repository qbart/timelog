TEST?=./...

test:
	go test $(TEST) -timeout=2m -parallel=4
