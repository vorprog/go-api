#!/usr/bin/env bash
set -xe

export ENVIRONMENT=${1:-dev}
export HOST_PORT=${2:-8080}
export CONTAINER_PORT=${3:-8080}

docker run \
--env APP_ENVIRONMENT_CONFIGURATION=$ENVIRONMENT \
--publish $HOST_PORT:$CONTAINER_PORT \
--detach \
--tty \
--interactive \
--rm \
go-server --port $CONTAINER_PORT
