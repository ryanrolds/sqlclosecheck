# sqlclosecheck

Linter that checks if SQL rows/statements are closed. Unclosed rows and statements may
cause DB connection pool exhaustion. Included in `golangci-lint` as `sqlclosecheck`.

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
