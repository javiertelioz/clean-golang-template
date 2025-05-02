export

LOCAL_BIN:=$(CURDIR)/bin
PATH:=$(LOCAL_BIN):$(PATH)

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## Display this help screen
	@@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m<target>\033[0m\n\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

run: ## Run the application
	go run ./cmd/api/main.go
.PHONY: run

test: ## Clear the test cache and then execute all project tests with coverage.
	@mkdir -p coverage
	@go clean -testcache
	go test -v -failfast -race -cover -covermode=atomic ./test/... -coverpkg=./pkg/... -coverprofile=coverage/coverage.out -shuffle=on
	@echo "🧪 Test Completed"
.PHONY: test

coverage: ## Generate and visualize a test coverage report in HTML format.
	@mkdir -p coverage
	@go clean -testcache
	@go test -v -failfast -race -cover -covermode=atomic ./test/... -coverpkg=./pkg/... -coverprofile=coverage/coverage.out -shuffle=on > /dev/null
	@go tool cover -func=coverage/coverage.out
	@go tool cover -html=coverage/coverage.out -o coverage/coverage.html
	@echo "🧪 Test coverage completed"
.PHONY: coverage

bench: ## Run the benchmarks
	go test -bench=. -benchmem -benchtime=5s -count=5 ./test/bench/... -v
.PHONY: bench

linter: ## Run the golangci-lint on the project source code to detect style issues or errors.
	golangci-lint run
.PHONY: linter

generate: ## Generate the proto files
	@buf generate proto/helloworld/v1/*.proto
	@echo "Success generate"
.PHONY: generate

clean: ## Clean the generated files
	@rm -f proto/helloworld/v1/*.{go,java,py,ts,yaml} || true
	@echo "Success clean folder"
.PHONY: clean

swagger: ## Format and initialize API documentation generation with Swaggo.
	@swag fmt
	swag init -g ./cmd/api/main.go -o ./docs/specs --parseInternal true
.PHONY: swagger
