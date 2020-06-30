PHONY: build install test

build:
	go build

install:
	go install

test:
	go test ./...

lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.27.0
	./bin/golangci-lint run
