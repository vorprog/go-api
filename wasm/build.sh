#!/usr/bin/env bash
set -xe

docker pull registry.hub.docker.com/library/golang:alpine

GIT_CURRENT_COMMIT_HASH=$(git rev-parse HEAD)

docker build \
--build-arg BUILD_COMMIT=${GIT_CURRENT_COMMIT_HASH} \
--tag wasm_go:latest .