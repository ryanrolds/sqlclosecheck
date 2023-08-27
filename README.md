# sqlclosecheck

Linter that checks if SQL rows/statements are closed. Unclosed rows and statements may
cause DB connection pool exhaustion.

## Running

```
make build
make install
```

In your project directory:
```
go vet -vettool=$(which sqlclosecheck) ./...
```

## CI

```
go install github.com/ryanrolds/sqlclosecheck@latest
go vet -vettool=${GOPATH}/bin/sqlclosecheck ./...
```
