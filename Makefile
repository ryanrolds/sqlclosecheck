PHONY: build install test

build:
	go build -o sqlclosecheck cmd/sqlclosecheck/main.go

install:
	go install ./cmd/sqlclosecheck

test:
	go test ./...

lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.27.0
	./bin/golangci-lint run
