PHONY: build install test

build:
	go build

install:
	go install

test: build install
	go test ./...
	# Due to an issue with importing in a anaylsistest's test data some hoop jumping is required
	# I call twice to avoid collecting package downloads in output
	-go vet -vettool=${GOPATH}/bin/sqlclosecheck ./testdata/sqlx_examples
	-go vet -vettool=${GOPATH}/bin/sqlclosecheck ./testdata/sqlx_examples 2> sqlx_examples_results.txt
	diff -a sqlx_examples_results.txt ./testdata/sqlx_examples/expected_results.txt

lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.27.0
	./bin/golangci-lint run
