#!/usr/bin/env bash
set -xe

docker pull registry.hub.docker.com/library/golang:alpine 

GIT_CURRENT_COMMIT_HASH=$(git rev-parse HEAD)
CURRENT_DATE_VERSION=$(date +'%Y.%m.%d.%H.%M.%S')
docker build \
--build-arg BUILD_COMMIT=${GIT_CURRENT_COMMIT_HASH} \
--tag hello_world_go:latest .

export PORT=8080

docker run \
--env PORT=$PORT \
--publish $PORT:$PORT \
--detach \
--tty \
--interactive \
hello_world_go
