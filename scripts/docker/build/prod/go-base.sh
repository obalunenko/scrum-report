#!/bin/bash

set -Eeuo pipefail

SCRIPT_NAME="$(basename "$0")"
SCRIPT_DIR="$(dirname "$0")"
REPO_ROOT="$(cd "${SCRIPT_DIR}" && git rev-parse --show-toplevel)"
SCRIPTS_DIR="${REPO_ROOT}/scripts"

source "${SCRIPTS_DIR}/helpers-source.sh"

echo "${SCRIPT_NAME} is running... "

checkInstalled 'docker'

SHORTCOMMIT="$(git rev-parse --short HEAD)"
VERSION="$(git tag | sort -V | tail -1)"

if [ -z "${VERSION}" ] || [ "${VERSION}" = "${SHORTCOMMIT}" ]; then
  VERSION="v0.0.0"
fi

DOCKER_REPO="${DOCKER_REPO}"

docker build --rm --no-cache \
    -t "${DOCKER_REPO}scrum-report-go-base-prod:${VERSION}" \
    -t "${DOCKER_REPO}scrum-report-go-base-prod:latest" \
    -f "${REPO_ROOT}/build/docker/base/go.Dockerfile" \
    --build-arg DOCKER_REPO="${DOCKER_REPO}" .

echo "${SCRIPT_NAME} done."
