TOOL_MOCKERY ?= go run github.com/vektra/mockery/v3@v3.5.5

.PHONY: unit-tests unit-tests-verbose integration-tests integration-tests-verbose separate-tests fuzzy-tests \
        generate-mocks clean-mocks

unit-tests:
	cd tests/unit && go test ./... -count=1

unit-tests-verbose:
	cd tests/unit && go test ./... -count=1 -v

integration-tests:
	cd tests/integration && go test ./... -count=1 -tags=integration

integration-tests-verbose:
	cd tests/integration && go test ./... -count=1 -tags=integration -v

separate-tests:
	cd tests/separate && go test ./... -count=1 -tags=separate

fuzzy-tests:
	cd tests/fuzzy && go test -fuzz ./... -count=1 -v

.PHONY: go-mod-tidy

go-mod-tidy:
	go mod tidy
	cd tests/unit && go mod tidy
	cd tests/separate && go mod tidy
	cd tests/integration && go mod tidy

clean-mocks:
	rm -rf tests/mocking/mocks

generate-mocks: clean-mocks
	${TOOL_MOCKERY}

sample-of-logs:
	go test github.com/Radek-Pysny/go-tests/tests/unit -run ^TestVerbose$$/^flat.*$$ -v
	go test github.com/Radek-Pysny/go-tests/tests/unit -run ^TestVerbose$$/^.*loop.*$$ -v
