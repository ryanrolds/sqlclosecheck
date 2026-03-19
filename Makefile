BIN := bin

PHONY: build install test

$(BIN):
	mkdir -p $@

build: $(BIN)
	go build -o $(BIN)/sqlclosecheck .

install:
	go install

test: build
	go test ./...

lint:
	curl -sSfL https://golangci-lint.run/install.sh | sh -s -- v2.11.3
	./bin/golangci-lint run
