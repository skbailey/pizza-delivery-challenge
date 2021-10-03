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

You can view test coverage data
```bash
go test -race -covermode=atomic -coverprofile=coverage.out ./...
```

## Linting

We use `golangci-lint`. Once you install the program (via `brew` or [otherwise]()), run the following command at project root.

```bash
golangci-lint run
```
The `golangci.yml` file at project root contains the configuration options for the linter.

## Docker

You may use the `Dockerfile` in the project to build an image.

```bash
docker build -t directions/pizza .
```

You can then run the code in that image.

```bash
docker run --rm -it --name pizza-bae directions/pizza make run
```

## Make

As a convenience, you can use `make` as a shortcut to some of the commands listed above.

To run the application with the default options

```bash
make run
```

To pass arguments to the `make run` command

```bash
deliverers=2 input=directions/zigzag.txt make run
```

To run tests
```bash
make test
```

To build a docker image

```bash
make docker.build
```

To build a predefined image
```bash
image=goat/pizza make docker.build
```

To run a predefined image
```bash
image=goat/pizza make docker.run
```

Take a look at the `Makefile` at the root of the project to view all possible options.

## Changelog
