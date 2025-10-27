# Description

Just some examples for unit testing in Go. All the test-related code can be found in `tests` directory.

```
/go-tests/
  ./tests/
    ./separate/*      separate test suit with all tests failing (separate tag)
    ./integration/*   separate test suit presenting basic usage of testcontainers
    ./unit/*          basic unit tests (no separation)
```


# Usage

For running unit tests use either `make unit-tests` or `make unit-tests-verbose`.

For running of all failing tests run `make separate-tests`.

For running fuzz testing sample run `make fuzzy-tests`.

To run `go mod tidy` across all sub-modules run `make go-mod-tidy`
