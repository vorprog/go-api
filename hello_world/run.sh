#!/usr/bin/env bash
set -xe

export ENVIRONMENT=${1:-dev}
export PORT=${2:-8080}

docker run \
--env APP_ENVIRONMENT_CONFIGURATION=$ENVIRONMENT \
--publish $PORT:8080 \
--detach \
--tty \
--interactive \
hello_world_go --port 8080
