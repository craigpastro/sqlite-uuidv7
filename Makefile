.PHONY: build
build:
	go build -buildmode=c-shared -o uuidv7.so .

.PHONY: lint
lint:
	golangci-lint run
