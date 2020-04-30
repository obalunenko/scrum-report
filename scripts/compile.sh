#!/usr/bin/env sh
set -e
echo "Building..."

BIN_OUT=./bin/scrum-report

go build -o ${BIN_OUT} ./cmd/scrum-report

echo "Binary compiled at ${BIN_OUT}"