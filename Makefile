# See https://golangci-lint.run/usage/install/
LINTER_VERSION = v1.55.0

# To be used for dependencies not installed with gomod
LOCAL_DEPS_INSTALL_LOCATION = /usr/local/bin

.PHONY: deps
deps:
	go env -w "GOPRIVATE=github.com/ildomm/*"
	go mod download

.PHONY: unit-test
unit-test: deps
	go test -tags=testing -count=1 -v github.com/ildomm/recoverich/...

.PHONY: lint-install
lint-install:
	[ -e ${LOCAL_DEPS_INSTALL_LOCATION}/golangci-lint ] || \
	wget -O- -q https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sudo sh -s -- -b ${LOCAL_DEPS_INSTALL_LOCATION} ${LINTER_VERSION}

.PHONY: lint
lint: deps lint-install
	golangci-lint run

.PHONY: coverage
coverage: coverage-core

.PHONY: coverage-core
coverage-core: deps
	go test -tags=testing \
		-coverprofile=build/cover.out github.com/ildomm/recoverich/...

.PHONY: coverage-report
coverage-report: deps coverage-core
	go tool cover -html=build/cover.out -o build/coverage.html
	echo "** Coverage is available in build/coverage.html **"
