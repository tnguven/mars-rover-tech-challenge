service ?= mars-rover

build:
	@go build -o "bin/${service}" "./cmd/${service}"

run: build
	@./bin/${service}

test:
	@go run gotest.tools/gotestsum --format testname --debug --packages="./..." -- -count=2

test-watch:
	@go run gotest.tools/gotestsum --format testname --watch --packages="./..." -- -count=1

test-coverage:
	@go test -coverprofile=coverage.out ./...
	@go run gotest.tools/gotestsum --format testname --debug --packages="./..." -- -count=2
	@go tool cover -func=coverage.out
	@echo "Coverage report shown above."
