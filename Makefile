
include .env

PHONY: install
install:
	go install gotest.tools/gotestsum@latest
	go install github.com/joho/godotenv/cmd/dotenv@latest

# golangci-lint
GOLANGCI_BIN := $(LOCAL_BIN)/golangci-lint

lint:
	$(info Running lint against changed files...)
	$(GOLANGCI_BIN) run \
		--new-from-rev=origin/master \
		--config=.golangci.yaml \
		--sort-results \
		--max-issues-per-linter=1000 \
		--max-same-issues=1000 \
		./...

MIGRATION_FOLDER=./tools/migrations

.PRONY: migrations-up
migrations-up:
	goose -dir ./migrations postgres "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

.PHONY: .test
.test:
	$(LOCAL_BIN)/gotestsum \
		--format pkgname-and-test-fails \
		--format-hide-empty-pkg \
		--format-icons hivis \
		--packages $(GO_UNIT_TEST_DIRECTORY) \
		-- -cover -covermode=atomic -coverprofile=$(GO_UNIT_TEST_COVER_PROFILE).tmp --race -coverpkg=$(GO_UNIT_TEST_COVER_PKG)

.PHONY: test-cover
test-cover: .test
	grep -vE '$(GO_UNIT_TEST_COVER_EXCLUDE)' $(GO_UNIT_TEST_COVER_PROFILE).tmp > $(GO_UNIT_TEST_COVER_PROFILE)
	rm $(GO_UNIT_TEST_COVER_PROFILE).tmp
