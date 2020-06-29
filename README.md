# sqlclosecheck

# Template:
[![ryanrolds](https://circleci.com/github/ryanrolds/sqlclosecheck.svg?style=svg)](https://app.circleci.com/pipelines/github/ryanrolds/sqlrowsclose)

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

## Roadmap

* ~~Get linter working~~
* ~~Added some basic test cases~~
* Test across a bunch of projects
* Introduce linter to [golangci-lint](https://github.com/golangci/golangci-lint).