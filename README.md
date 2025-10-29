# Description

Just some examples for unit testing in Go. All the test-related code can be found in `tests` directory.

```
/go-tests/
  ./tests/
    ./separate/*      separate test suit with all tests failing (separate tag + sub-module)
    ./fuzzy/*         simple fuzz testing sample
    ./integration/*   separate test suit presenting basic usage of testcontainers (sub-module)
    ./unit/*          basic unit tests (no separation)
```


# Usage

For running unit tests use either `make unit-tests` or `make unit-tests-verbose`.

To run integration tests using [testcontainers](https://testcontainers.com/) run `make integration-tests`
or `make integration-tests-verbose`. The latter command is convenient especially when debugging issues with
testcontainers and Docker image retrieval atc.

For running of all failing tests run `make separate-tests`.

For running fuzz testing sample run `make fuzzy-tests`.

To run `go mod tidy` across all sub-modules run `make go-mod-tidy`
