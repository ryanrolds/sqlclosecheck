# sqlclosecheck

Linter that checks if SQL rows/statements are closed. Unclosed rows and statements may
cause DB connection pool exhaustion. Included in `golangci-lint` as `sqlclosecheck`.

## Analyzers

* `legacy` - require that Close be called (LEGACY)
* `closed` - require that Close be called (EXPERIMENTAL)
* `defer-only` - require that Close be deferred (FUTURE)

## Running

```
make build
make install
```

In your project directory:
```
go vet -vettool=$(which sqlclosecheck) ./...
```

## Developers

Start by creating a test that should pass/fail.
Test are located at `pkg/analyzer/testdata`. 
All PRs that modify the analyzer should include a test.
Negative tests are just as important as positive tests.

Make changes to the analyzer (`pkg/analyzer`) and run the tests:
```
make test
```

## CI

GitHub Actions that runs on push to `main` and PRs.
