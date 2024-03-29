#!/bin/bash

# Copyright 2020-2021 Hewlett Packard Enterprise Development LP
# NOTE -- The following syntax will run unit tests without worrying
# about code coverage. To see an example of code coverage (which re-runs
# these same unit tests) see runCoverage.sh. The same tests are getting
# run in both places.

set -ex

docker run --rm -v ${WORKSPACE}:/mnt/workspace -w /mnt/workspace arti.dev.cray.com/baseos-docker-master-local/golang:alpine3.12 /bin/sh -c '
set -ex -o pipefail

echo "Unit Tests for Go"
TEST_OUTPUT_DIR="$PWD/build/results/unittest"
mkdir -p $TEST_OUTPUT_DIR
export GOPATH="$HOME/go"
export PATH="$PATH:$GOPATH/bin"
apk add git
go get -u github.com/jstemmer/go-junit-report
export CGO_ENABLED=0
go test *.go -v | go-junit-report > "$TEST_OUTPUT_DIR/testing.xml"
'
