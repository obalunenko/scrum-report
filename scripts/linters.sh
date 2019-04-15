#!/usr/bin/env bash


function vet(){
    echo "vet project..."
    go vet $(go list ./...)
    EXIT_CODE=$?
    if [[ ${EXIT_CODE} -ne 0 ]]; then
        exit 1
    fi
    echo ""
}

function fmt(){
    echo "fmt lint..."
    declare -a fmts=$(gofmt -s -l  $(find . -type f -name '*.go' | grep -v 'vendor' |grep -v '.git' |grep -v '*/bindata.go'))

    if [[ ${fmts} ]]; then
        echo "fix it:"
        for f in "${fmts[@]}"
        do
            echo "$f"

        done
        exit 1

    else
        echo "code is ok"
        echo ${fmts}
    fi
    echo ""
}

function go-lint(){
    echo "golint..."
    declare -a lints=$(golint $(go list ./...|grep -v 'web'))  ## its a hack to not lint generated code
    if [[ ${lints} ]]; then
        echo "fix it:"
        for l in "${lints[@]}"
        do
            echo "$l"

        done
        exit 1

    else
        echo "code is ok"
        echo ${lints}
    fi
    echo ""
}

function go-group()
{
    echo "gogroup..."

    declare -a lints=$(gogroup -order std,other,prefix=github.com/oleg-balunenko/  $(find . -type f -name "*.go" | grep -v "vendor/" |grep -v '*/bindata.go'))
    if [[ ${lints} ]]; then
        echo "fix it:"
        for l in "${lints[@]}"
        do
            echo "$l"

        done
        exit 1

    else
        echo "code is ok"
        echo ${lints}
    fi
    echo ""

}


function golangci(){
    echo "golang-ci linter running..."
    if [[ -f "$(go env GOPATH)/bin/golangci-lint" ]] || [[ -f "/usr/local/bin/golangci-lint" ]]; then
        golangci-lint run ./...
    else
        printf "Cannot check golang-ci, please run:
        curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.12.5 \n"
        exit 1
    fi
    echo ""
}
