# sqlclosecheck

[![ryanrolds](https://circleci.com/gh/ryanrolds/sqlclosecheck.svg?style=svg)](https://app.circleci.com/pipelines/github/ryanrolds/sqlrowsclose)

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
go install github.com/ryanrolds/sqlclosecheck
go vet -vettool=${GOPATH}/bin/sqlclosecheck ./...
```

## Roadmap

* ~~Get linter working~~
* ~~Added some basic test cases~~
* ~~Require that Close be deferred~~
* Add sqlx checking
* Test across a bunch of projects
* Introduce linter to [golangci-lint](https://github.com/golangci/golangci-lint).