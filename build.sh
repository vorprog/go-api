#!/usr/bin/env bash
set -xe

sudo docker pull registry.hub.docker.com/library/golang:alpine

GIT_CURRENT_COMMIT_HASH=$(git rev-parse HEAD)

sudo docker build \
--build-arg BUILD_COMMIT=${GIT_CURRENT_COMMIT_HASH} \
--tag go-server:latest .
