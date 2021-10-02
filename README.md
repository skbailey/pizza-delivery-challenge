## Pizza Delivery

## Usage

By default, this will run the program with 1 delivery person (termed `deliverer` in the code because sometimes it's not a person at all). The `file-input` argument is required.
```bash
go run main.go -file-input <path to file>
```

You can alter the number of `deliverers` by passing the `deliverer-count` argument.
```bash
go run main.go -file-input <path to file> -deliverer-count 2
```

## Documentation

## Tests

You can run tests from the root of the project directory.
```bash
go test -v ./...
```

## Linting

We use `golangci-lint`. Once you install the program (via `brew` or [otherwise]()), run the following command at project root.

```bash
golangci-lint run
```
The `golangci.yml` file at project root contains the configuration options for the linter.

## Docker

## Changelog
