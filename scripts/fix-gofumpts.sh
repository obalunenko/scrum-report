#!/usr/bin/env bash

if [[ -f "$(go env GOPATH)/bin/gofumpt" ]] || [[ -f "/usr/local/bin/gofumpt" ]]; then
    gofmt -w $(find . -type f -name '*.go' | grep -v 'vendor' |grep -v '.git' | grep -v "*.generated.go")
else
    printf "Cannot check gogroup, please run:
    go get -u -v mvdan.cc/gofumpt/... \n"
    exit 1
fi