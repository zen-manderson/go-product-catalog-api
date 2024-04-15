include env/local.env

.DEFAULT_GOAL := help

.PHONY: help
help: ## List of available commands
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: test
test: ## Run the unit tests locally
	@echo ""
	@echo "Running unit tests with coverage analysis; does not check for race conditions."
	@echo ""
	@echo "Developers should run 'make test-race' separately before pushing commits."
	@echo ""
	go test ./... -coverprofile cover.out
	go tool cover -func cover.out

.PHONY: test-race
test-race: ## Run the unit tests locally with race condition checks
	@echo ""
	@echo "Run unit tests with race condition checks."
	@echo ""
	@echo "The cache of previous test results is cleared to ensure the best race condition checks possible"
	@echo ""
	@echo "If race conditions are reported, developers should examine them carefully and disregard any that"
	@echo "arise solely because of unit tests attempting to collect and analyze log data; these are "
	@echo "erroneous reports and do not reflect how the code might behave in production use."
	@echo ""
	go clean -testcache
	go test ./... -race

## For more sophisticated examples see the Makefile in https://github.com/zenbusiness/pg-event-distributor/blob/main/Makefile
## There you will find examples of running gcloud's pub/sub emulator and postgres via docker