# sqlclosecheck

Linter that checks if SQL rows/statements are closed. Unclosed rows and statements may
cause DB connection pool exhaustion. Included in `golangci-lint` as `sqlclosecheck`.

## Analyzers

* `defer-only` - require that Close be deferred
* `closed` - require that Close be called (WIP)

## Supported packages

* `database/sql`
*	`github.com/jmoiron/sqlx`
* `github.com/jackc/pgx/v5`
* `github.com/jackc/pgx/v5/pgxpool`

## Contributing

Contributions, bug reports, and feature requests are welcome.

### Running

```
make build
make install
```

In your project directory:
```
go vet -vettool=$(which sqlclosecheck) ./...
```

### When making changes

Start by creating a test that should pass/fail.
Test are located at `pkg/analyzer/testdata`. 
All PRs that modify the analyzer should include a test.
Negative tests are just as important as positive tests.

Make changes to the analyzer (`pkg/analyzer`) and run the tests:
```
make test
```

Sometimes it's helpful to look at the SSA for a test function:
```
env GOSSAFUNC=correctDeferBlock go build pkg/analyzer/testdata/rows/common.go pkg/analyzer/testdata/rows/correct_defer_block.go
```

### CI

GitHub Actions runs on push to `main` and PRs by the project lead.
