# sqlclosecheck

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
* Test across a bunch of projects
* Introduce linter to [golangci-lint](https://github.com/golangci/golangci-lint).