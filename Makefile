LOCAL_BIN := $(CURDIR)/bin

include ./.env

export PATH := $(PATH):$(LOCAL_BIN)

.PHONY: deps
deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN=$(LOCAL_BIN) go install gotest.tools/gotestsum@latest
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@latest
	GOBIN=$(LOCAL_BIN) go install github.com/onsi/ginkgo/v2/ginkgo@latest

.PHONY: build
build:
	go build -o bin/app cmd/main.go

.PHONY: infra
infra:
	docker-compose up -d --build --force-recreate --wait

.PHONY: minfra
minfra-up: infra
	sleep 2s && \
	$(LOCAL_BIN)/goose postgres "user=${DB_USER} password=${DB_PASSWORD} host=${DB_HOST} dbname=${DB_NAME} sslmode=disable" -dir=./migrations up

.PHONY: minfra-down
minfra-down:
	$(LOCAL_BIN)/goose postgres "user=${DB_USER} password=${DB_PASSWORD} host=${DB_HOST} dbname=${DB_NAME} sslmode=disable" -dir=./migrations reset

.PHONY: migrations-up
migrations-up:
	$(LOCAL_BIN)/goose postgres "user=${DB_USER} password=${DB_PASSWORD} host=${DB_HOST} dbname=${DB_NAME} sslmode=disable" -dir=./migrations up

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


.PRONY: migrations-e2e-up ## накатывает миграции на базу данных для тестов
migrations-e2e-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_FOLDER} postgres "host=localhost port=5432 user=${DB_USER} password=${DB_PASSWORD} dbname=${E2E_DB_NAME} sslmode=disable" up

# golangci-lint
GOLANGCI_BIN := $(LOCAL_BIN)/golangci-lint

.PRONY: lint
lint:
	$(info Running lint against changed files...)
	$(LOCAL_BIN)/golangci-lint run \
	--new-from-rev=origin/master \
	--config=.golangci.yaml \
	--sort-results \
	--max-issues-per-linter=1000 \
	--max-same-issues=1000 \
	--verbose \
	./...


MIGRATION_FOLDER=./migrations


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

.PHONY: test
test:
	$(LOCAL_BIN)/gotestsum \
		--format pkgname-and-test-fails \
		--format-hide-empty-pkg \
		--format-icons hivis \
		--packages $(GO_UNIT_TEST_DIRECTORY) \
		-- -cover -covermode=atomic -coverprofile=$(GO_UNIT_TEST_COVER_PROFILE).tmp --race -coverpkg=$(GO_UNIT_TEST_COVER_PKG)

.PHONY: test-cover
test-cover: test
	grep -vE '$(GO_UNIT_TEST_COVER_EXCLUDE)' $(GO_UNIT_TEST_COVER_PROFILE).tmp > $(GO_UNIT_TEST_COVER_PROFILE)
	rm $(GO_UNIT_TEST_COVER_PROFILE).tmp

.PHONY: print-test-cover
print-test-cover:
	go tool cover -func=$(GO_UNIT_TEST_COVER_PROFILE) | grep total:


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
generate:
	protoc \
		-I ./api \
		-I ./vendor.protogen \
		--go_out=./internal/pb/medblogers_base \
		--go-grpc_out=./internal/pb/medblogers_base \
		--grpc-gateway_out=./internal/pb/medblogers_base \
		--validate_out="lang=go:./internal/pb/medblogers_base" \
		--openapiv2_out=./internal/pb/medblogers_base \
		--openapiv2_opt logtostderr=true \
		--openapiv2_opt allow_merge=true \
		--openapiv2_opt merge_file_name=medblogers_api \
		./api/doctors/v1/* ./api/freelancers/v1/* ./api/auth/v1/*


.PHONY: e2e ## запускает локальные интеграционные тесты
e2e: infra e2e-run

# Вы можете включить:
#	- сжатый режим с помощью ginkgo --succinct,
#	- подробный режим с помощью ginkgo -v
#	- очень подробный режим с помощью ginkgo -vv.
.PHONY: e2e-run  ## запускает интеграционные тесты
e2e-run:
	SERVICE_NAME=e2e $(LOCAL_BIN)/ginkgo \
		--junit-report=./junit.xml \
		-tags=e2e -cover -covermode=count \
		-coverprofile=$(GO_INTEGRATION_TEST_COVER_PROFILE).tmp \
		-coverpkg=$(GO_INTEGRATION_TEST_COVER_PKG) -succinct ./e2e/...

.PHONY: e2e-cover
e2e-cover:
	grep -vE '$(GO_INTEGRATION_TEST_COVER_EXCLUDE)' $(GO_INTEGRATION_TEST_COVER_PROFILE).tmp > $(GO_INTEGRATION_TEST_COVER_PROFILE)
	rm $(GO_INTEGRATION_TEST_COVER_PROFILE).tmp
