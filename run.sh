#!/usr/bin/env bash
set -xe

export ENVIRONMENT=${1:-dev}
export HOST_PORT=${2:-8080}
export CONTAINER_PORT=${3:-8080}

sudo docker run \
--log-driver=fluentd \
--log-opt tag="{{.ImageName}}/{{.ImageID}}/{{.ID}}" \
--log-opt fluentd-sub-second-precision=true \
--env APP_ENVIRONMENT_CONFIGURATION=$ENVIRONMENT \
--publish $HOST_PORT:$CONTAINER_PORT \
--detach \
--tty \
--interactive \
--rm \
--name go-api \
go-api --port $CONTAINER_PORT
