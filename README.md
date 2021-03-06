# Pizza Delivery

## Usage

By default, this will run the program with 1 delivery person (termed `deliverer` in the code because sometimes it's not a person at all). The `--file-input` argument is required.
```bash
go run main.go --file-input <path to file>
```

You can alter the number of `deliverers` by passing the `deliverer-count` argument.
```bash
go run main.go --file-input <path to file> --deliverer-count 2
```

## Documentation

There are two main top-level packages:
- `retriever`
- `delivery`

### The `retriever` package

There is one exported function in this package: `Read`

```go
func Read(path string) (string, error)
    Read reads input from a file
```

This package reads string input from a file and returns it.

```go
input, err := retriever.Read("path/to/file")
```

### The `delivery` package

There is one exported function in this package: `Deliver`

```go
func Deliver(input string, delivererCount int) (int, error)
    Deliver coordinates the delivery of pizzas
```

This package takes the string output from above (extracted from the given file) as well as a count of the number of `deliverers` and runs the algorithm responsible for delivering pizzas. On completion, this package returns the number of unique houses visited.

```go
housesVisited, err := delivery.Deliver("^^>v<<", 2)
```

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

We use `golangci-lint`. Once you install the program (via `brew` or [otherwise](https://github.com/golangci/golangci-lint#install-golangci-lint)), run the following command at project root.

```bash
golangci-lint run
```
The `.golangci.yml` file at project root contains the configuration options for the linter. We use the [sample default linting configuration
](https://raw.githubusercontent.com/golangci/golangci-lint/master/.golangci.yml). We recommend [installing a pre commit hook](https://freshman.tech/linting-golang/#setting-up-a-pre-commit-hook) to run the linters beforehand.

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

## Release Management

We use [trunk-based workflow](https://www.atlassian.com/continuous-delivery/continuous-integration/trunk-based-development) for release management. Features are broken up into smaller releasable chunks and we use short-lived branches with few commits to merge these chunks frequently to the `main` branch. This is an alternative to `gitflow` which is also an option but the concern is that long-lived feature branches increase the likelihood of needing to handle conflicts and result in higher complexity around releases eg code freezes.

### Feature Flags
In order to continuously release the component parts of a feature, use feature flags to dark launch code into production that can then be "switched on" eg using LaunchDarkly. There is an inherent cleanup of feature flags required once the feature is turned on and functioning as expected.

### Releases
We trigger code releases/deployments from the `main` branch so this branch must always be in a deployable state. For more detail on cutting a release and creating hotfixes via an intermediate release candidate branch [follow these instructions](https://www.nebulaworks.com/insights/posts/trunk-based-development-for-beginners/).
