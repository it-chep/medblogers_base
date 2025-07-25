LOCAL_BIN := $(CURDIR)/bin

include ./.env

export PATH := $(PATH):$(LOCAL_BIN)

.PHONY: deps
deps:
	GOBIN=$(LOCAL_BIN) go install gitlab.ozon.ru/whc/go/libs/xo@v1.0.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN=$(LOCAL_BIN) go install gotest.tools/gotestsum@latest
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: infra
infra:
	docker-compose up -d --build

.PHONY: minfra
minfra-up: infra
	sleep 2s && \
	$(LOCAL_BIN)/goose postgres "user=${DB_USER} password=${DB_PASSWORD} host=${DB_HOST} dbname=${DB_NAME} sslmode=disable" -dir=./migrations up

.PHONY: minfra-down
minfra-down:
	$(LOCAL_BIN)/goose postgres "user=${DB_USER} password=${DB_PASSWORD} host=${DB_HOST} dbname=${DB_NAME} sslmode=disable" -dir=./migrations reset

# If the first argument is "migration"...
ifeq (migration,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  MIGRATION_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(MIGRATION_ARGS):;@:)
endif

.PHONY: migration
migration:
	$(LOCAL_BIN)/goose -dir=./migrations create $(MIGRATION_ARGS) sql

# golangci-lint
GOLANGCI_BIN := $(LOCAL_BIN)/golangci-lint

lint:
	$(info Running lint against changed files...)
	$(LOCAL_BIN)/golangci-lint run \
		--new-from-rev=origin/master \
		--config=.golangci.yaml \
		--sort-results \
		--max-issues-per-linter=1000 \
		--max-same-issues=1000 \
		./...

MIGRATION_FOLDER=./tools/migrations


# ==================================================================================== #
# TESTS ENVS
# ==================================================================================== #

GO_UNIT_TEST_DIRECTORY:=./internal/...
GO_UNIT_TEST_COVER_PKG:=${GO_UNIT_TEST_DIRECTORY}

GO_UNIT_TEST_COVER_EXCLUDE:=mocks|config|.pb.go|.pb.*.go|app|swagger.go|e2e|xo|_mock.go|.pb.go|.pb.goclay.go|.pb.scratch.go|.pb.gw.go|.pb.sensitivity.go|_minimock.go|models_gen.go|generated.go|pb.*.go|internal\/config\/config.go
GO_UNIT_TEST_COVER_PROFILE?=unit.coverage.out

GO_INTEGRATION_TEST_DIRECTORY:=./e2e/...
GO_INTEGRATION_TEST_COVER_PKG:=${GO_UNIT_TEST_DIRECTORY}
GO_INTEGRATION_TEST_COVER_PROFILE?=cover.out
GO_INTEGRATION_TEST_COVER_EXCLUDE:=${GO_UNIT_TEST_COVER_EXCLUDE}
# ==================================================================================== #
# TESTS
# ==================================================================================== #

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


XO_OUTPUT_PATH=./tools/xo
XO_TEMPLATE_PATH=./tools/xo_templates
.PHONY: xo ## генерация dto базы данных
xo:
	rm -r $(XO_OUTPUT_PATH)
	mkdir -p $(XO_OUTPUT_PATH)
	xo "pgsql://$(DB_USER):$(DB_PASSWORD)@localhost:5432/$(DB_NAME)?sslmode=disable" \
	-o $(XO_OUTPUT_PATH) --template-path $(XO_TEMPLATE_PATH) --schema public --suffix ".xo.go" --custom-type-package custom

	rm $(XO_OUTPUT_PATH)/goosedbversion.xo.go
	rm $(XO_OUTPUT_PATH)/xo_db.xo.go


.PHONY: generate
.generate: