#!/bin/env bash
set -ex
go get -d ./...

export BUILD_COMMIT=$(git rev-parse HEAD)
export CURRENT_DATE_VERSION=$(date --utc +'%Y.%m.%d.%H.%M.%S') 

CGO_ENABLED=0 \
go build \
-ldflags "-X github.com/vorprog/go-api/util.BuildCommitLinkerFlag=$BUILD_COMMIT -X github.com/vorprog/go-api/util.BuildDateVersionLinkerFlag=$CURRENT_DATE_VERSION"
