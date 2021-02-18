NAME=scrum-report
BIN_DIR=./bin

TARGET_MAX_CHAR_NUM=20

## Show help
help:
	${call colored, help is running...}
	@echo ''
	@echo 'Usage:'
	@echo '  make <target>'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  %-$(TARGET_MAX_CHAR_NUM)s %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)



## Compile executable
compile:
	./scripts/compile.sh
.PHONY: compile

## lint project
lint:
	./scripts/run-linters.sh
.PHONY: lint

lint-ci:
	./scripts/run-linters-ci.sh
.PHONY: lint-ci


## format markdown files in project
pretty-markdown:
	find . -name '*.md' -not -wholename './vendor/*' | xargs prettier --write
.PHONY: pretty-markdown

## Test all packages
test:
	./scripts/run-tests.sh
.PHONY: test

## Test all packages
test-docker:
	./scripts/run-tests-docker.sh
.PHONY: test-docker

## Test coverage
test-cover:
	./scripts/coverage.sh
.PHONY: test-cover

## Add new version
new-version: lint test compile
	./scripts/version.sh
.PHONY: new-version

## Release
release:
	./scripts/release.sh
.PHONY: release

## Fix imports sorting
imports:
	./scripts/fix-imports.sh
.PHONY: imports

fumpt:
	./scripts/fix-gofumpts.sh
.PHONY: fumpt

## fetch all dependencies for scripts
install-tools:
	./scripts/get-dependencies.sh
.PHONY: install-tools

## Sync vendor
sync-vendor:
	${call colored, gomod is running...}
	./scripts/sync-vendor.sh
.PHONY: sync-vendor

## Update dependencies
gomod-update:
	${call colored, gomod is running...}
	go get -u -v ./...
	make sync-vendor
.PHONY: gomod-update

## Vet project
vet:
	./scripts/vet.sh
.PHONY: vet

## Docker compose up
docker-up:
	docker-compose -f ./docker-compose.yml up --build -d

.PHONY: docker-up

## Docker compose down
docker-down:
	docker-compose -f ./docker-compose.yml down --volumes

.PHONY: docker-down

## Docker compose up
docker-up-dev:
	docker-compose -f ./docker-compose.dev.yml up --build -d

.PHONY: docker-up

## Docker compose down
docker-down-dev:
	docker-compose -f ./dev.docker-compose.dev.yml down --volumes

.PHONY: docker-down

.DEFAULT_GOAL := test

