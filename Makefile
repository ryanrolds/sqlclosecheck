PHONY: build install test

build:
	go build -o sqlclosecheck cmd/sqlclosecheck/main.go

install:
	go install ./cmd/sqlclosecheck

test:
	go test ./...