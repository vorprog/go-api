#!/usr/bin/env bash
set -x

docker pull registry.hub.docker.com/library/golang:alpine 

GIT_CURRENT_COMMIT_HASH=$(git rev-parse HEAD)
GIT_CURRENT_COMMIT_TIMESTAMP=$(git show -s --format=%ci HEAD)
docker build \
--build-arg GIT_COMMIT_HASH=${GIT_CURRENT_COMMIT_HASH} \
--build-arg GIT_COMMIT_TIMESTAMP=${GIT_CURRENT_COMMIT_TIMESTAMP} \
--quiet \
--tag hello_world_go:latest .

export PORT=8080

docker run \
--env PORT=$PORT \
--publish $PORT:$PORT \
--detach \
--tty \
--interactive \
hello_world_go
