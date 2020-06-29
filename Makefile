PHONY: build install


build:
	go build -o sqlclosecheck cmd/sqlclosecheck/main.go


install:
	go install ./cmd/sqlclosecheck