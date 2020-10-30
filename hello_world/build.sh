#!/usr/bin/env bash
set -xe

docker pull registry.hub.docker.com/library/golang:alpine

GIT_CURRENT_COMMIT_HASH=$(git rev-parse HEAD)
CURRENT_DATE_VERSION=$(date +'%Y.%m.%d.%H.%M.%S')

docker build \
--build-arg BUILD_COMMIT=${GIT_CURRENT_COMMIT_HASH} \
--build-arg BUILD_DATE_VERSION=${CURRENT_DATE_VERSION} \
--tag hello_world_go:latest .