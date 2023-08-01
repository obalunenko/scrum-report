BIN_DIR=./bin

SHELL := env DOCKER_REPO=$(DOCKER_REPO) $(SHELL)
DOCKER_REPO?=olegbalunenko

SHELL := env VERSION=$(VERSION) $(SHELL)
VERSION ?= $(shell git describe --tags $(git rev-list --tags --max-count=1))

APP_NAME?=scrum-report
SHELL := env APP_NAME=$(APP_NAME) $(SHELL)

GOTOOLS_IMAGE_TAG?=v0.13.0
SHELL := env GOTOOLS_IMAGE_TAG=$(GOTOOLS_IMAGE_TAG) $(SHELL)

COMPOSE_TOOLS_FILE=deployments/docker-compose/go-tools-docker-compose.yml
COMPOSE_TOOLS_CMD_BASE=docker compose -f $(COMPOSE_TOOLS_FILE)
COMPOSE_TOOLS_CMD_UP=$(COMPOSE_TOOLS_CMD_BASE) up --remove-orphans --exit-code-from
COMPOSE_TOOLS_CMD_PULL=$(COMPOSE_TOOLS_CMD_BASE) build

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



## Build project.
build: sync-vendor generate compile-app
.PHONY: build

## Compile app.
compile-app:
	$(COMPOSE_TOOLS_CMD_UP) build build
.PHONY: compile-app

## Test coverage report.
test-cover:
	./scripts/tests/coverage.sh
.PHONY: test-cover

prepare-cover-report: test-cover
	$(COMPOSE_TOOLS_CMD_UP) prepare-cover-report prepare-cover-report
.PHONY: prepare-cover-report

## Open coverage report.
open-cover-report: prepare-cover-report
	./scripts/open-coverage-report.sh
.PHONY: open-cover-report

## Update readme coverage.
update-readme-cover: build prepare-cover-report
	$(COMPOSE_TOOLS_CMD_UP) update-readme-coverage update-readme-coverage
.PHONY: update-readme-cover

## Run tests.
test:
	$(COMPOSE_TOOLS_CMD_UP) run-tests run-tests
.PHONY: test

## Run regression tests.
test-regression: test
.PHONY: test-regression

## Sync vendor and install needed tools.
configure: sync-vendor install-tools

## Sync vendor with go.mod.
sync-vendor:
	./scripts/sync-vendor.sh
.PHONY: sync-vendor

## Fix imports sorting.
imports:
	$(COMPOSE_TOOLS_CMD_UP) fix-imports fix-imports
.PHONY: imports

## Format code with go fmt.
fmt:
	$(COMPOSE_TOOLS_CMD_UP) fix-fmt fix-fmt
.PHONY: fmt

## Format code and sort imports.
format-project: fmt imports
.PHONY: format-project

## Installs vendored tools.
install-tools:
	echo "Installing ${GOTOOLS_IMAGE_TAG}"
	$(COMPOSE_TOOLS_CMD_PULL)
.PHONY: install-tools

## vet project
vet:
	$(COMPOSE_TOOLS_CMD_UP) vet vet
.PHONY: vet

## Run full linting
lint-full:
	$(COMPOSE_TOOLS_CMD_UP) lint-full lint-full
.PHONY: lint-full

## Run linting for build pipeline
lint-pipeline:
	$(COMPOSE_TOOLS_CMD_UP) lint-pipeline lint-pipeline
.PHONY: lint-pipeline

## Run linting for sonar report
lint-sonar:
	$(COMPOSE_TOOLS_CMD_UP) lint-sonar lint-sonar
.PHONY: lint-sonar

## recreate all generated code and documentation.
codegen:
	$(COMPOSE_TOOLS_CMD_UP) go-generate go-generate
.PHONY: codegen

## recreate all generated code and swagger documentation and format code.
generate: codegen format-project vet
.PHONY: generate

## Release
release:
	$(COMPOSE_TOOLS_CMD_UP) release release
.PHONY: release

## Release local snapshot
release-local-snapshot:
	$(COMPOSE_TOOLS_CMD_UP) release-local-snapshot release-local-snapshot
.PHONY: release-local-snapshot

## Check goreleaser config.
check-releaser:
	$(COMPOSE_TOOLS_CMD_UP) release-check-config release-check-config
.PHONY: check-releaser

## Issue new release.
new-version: vet test-regression build
	./scripts/release/new-version.sh
.PHONY: new-release



######################################
############### DOCKER ###############
######################################

################ PROD #################

## Push all prod images to registry.
docker-push-prod-images:
	./scripts/docker/push-all-images-to-registry.sh ${DOCKER_REPO}
.PHONY: docker-push-prod-images

## Build docker base images.
docker-build-base-prod: docker-build-base-go-prod
.PHONY: docker-build-base-prod

## Build docker base image for GO
docker-build-base-go-prod:
	./scripts/docker/build/prod/go-base.sh
.PHONY: docker-build-base-go-prod

## Build all services docker prod images for deploying to gcloud.
docker-build-prod: docker-build-backend-prod
.PHONY: docker-build-prod

## Build all backend services docker prod images for deploying to gcloud.
docker-build-backend-prod: docker-build-scrum-report-prod
.PHONY: docker-build-backend-prod

## Build admin service prod docker image.
docker-build-scrum-report-prod:
	./scripts/docker/build/prod/scrum-report.sh
.PHONY: docker-build-scrum-report-prod

## Docker compose up - deploys prod containers on docker locally.
docker-compose-up:
	./scripts/docker/compose/prod/up.sh
.PHONY: docker-compose-up

## Docker compose down - remove all prod containers in docker locally.
docker-compose-down:
	./scripts/docker/compose/prod/down.sh
.PHONY: docker-compose-down

## Docker compose stop - stops all prod containers in docker locally.
docker-compose-stop:
	./scripts/docker/compose/prod/stop.sh
.PHONY: docker-compose-stop

## Build all prod images: base and services.
docker-prepare-images-prod: docker-build-base-prod docker-build-prod
.PHONY: docker-prepare-images-prod

## Prod local full deploy: build base images, build services images, deploy to docker compose
deploy-local-prod: docker-prepare-images-prod run-local-prod
.PHONY: deploy-local-prod

## Run locally: deploy to docker compose and expose tunnels.
run-local-prod: docker-compose-up
.PHONY: run-local-prod

## Stop the world and close tunnels.
stop-local-prod: docker-compose-stop
.PHONY: stop-local-prod

################## DEV ###################

## Build docker base images.
docker-build-base-dev: docker-build-base-go-dev
.PHONY: docker-build-base-dev

## Build docker base image for GO
docker-build-base-go-dev:
	./scripts/docker/build/dev/go-base.sh
.PHONY: docker-build-base-go-dev

## Build docker dev image for running locally.
docker-build-dev: docker-build-scrum-report-dev
.PHONY: docker-build-dev

## Build admin service dev docker image.
docker-build-scrum-report-dev:
	./scripts/docker/build/dev/scrum-report.sh
.PHONY: docker-build-scrum-report-dev

## Dev Docker-compose up with stubbed 3rd party dependencies.
dev-docker-compose-up:
	./scripts/docker/compose/dev/up.sh
.PHONY: dev-docker-compose-up

## Docker compose down.
dev-docker-compose-down:
	./scripts/docker/compose/dev/down.sh
.PHONY: dev-docker-compose-down

## Docker compose stop - stops all dev containers in docker locally.
dev-docker-compose-stop:
	./scripts/docker/compose/dev/stop.sh
.PHONY: dev-docker-compose-stop

## Build all dev images: base and services.
docker-prepare-images-dev: docker-build-base-dev docker-build-dev
.PHONY: docker-prepare-images-dev

## Dev local full deploy: build base images, build services images, deploy to docker compose
deploy-local-dev: docker-prepare-images-dev run-local-dev
.PHONY: deploy-local-dev

## Run locally dev: deploy to docker compose and expose tunnels.
run-local-dev: dev-docker-compose-up
.PHONY: run-local-dev

## Stop the world and close tunnels.
stop-local-dev: dev-docker-compose-stop
.PHONY: stop-local-prod

## Open containers logs service url.
open-container-logs:
	./scripts/browser-opener.sh -u 'http://localhost:9999/'
.PHONY: open-container-logs

## Opens url of web amin in browser.
open-scrum-report:
	./scripts/browser-opener.sh -u 'http://localhost.charlesproxy.com:8080/'
.PHONY: open-scrum-report

.DEFAULT_GOAL := help

