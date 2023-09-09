# sqlclosecheck

Linter that checks if SQL rows/statements are closed. Unclosed rows and statements may
cause DB connection pool exhaustion.

## Analyzers

* `defer-only` - require that Close be deferred
* `closed` - require that Close be called (EXPERIMENTAL)

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

## Roadmap

* ~~Get linter working~~
* ~~Added some basic test cases~~
* ~~Require that Close be deferred~~
* ~~Add sqlx checking~~
* ~~Test across a bunch of projects~~
* ~~Introduce linter to [golangci-lint](https://github.com/golangci/golangci-lint).~~
