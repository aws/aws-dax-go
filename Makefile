.DEFAULT_GOAL := help

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: lint
lint:  ## Lint Go source files
	@golangci-lint run


.PHONY: test
test: ## Run unit tests
	go test -race \
			-coverpkg=./... \
			-coverprofile=coverageunit.out \
			-covermode=atomic \
			-count=1 \
			-timeout=5m \
			./...

.PHONY: bench
bench: ## Run benchmark tests
	go test ./... -bench=. -run=XXX -benchmem -benchtime=1s -cpu 1
