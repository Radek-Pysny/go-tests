.PHONY: unit-tests unit-tests-verbose integration-tests integration-tests-verbose separate-tests fuzzy-tests

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
