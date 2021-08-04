#!/usr/bin/env bash
set -xe

export IMAGE_NAME=go-api
export COMMIT_SHA_TAG=$(git rev-parse HEAD)
export APP_ENVIRONMENT_CONFIGURATION=${1:-dev}
export HOST_PORT=${2:-80}
export APP_SERVER_PORT=8080
export FLUENTD_HOST="fluentd.vorprog.com:24224"

sudo docker build \
--build-arg BUILD_COMMIT=$COMMIT_SHA_TAG \
--tag $IMAGE_NAME:$COMMIT_SHA_TAG \
--tag go-api:latest \
.

sudo docker run \
--log-driver=fluentd \
--log-opt fluentd-address=$FLUENTD_HOST \
--log-opt fluentd-sub-second-precision=true \
--log-opt tag="docker.{{.ImageName}}.{{.ImageID}}.{{.ID}}" \
--env APP_ENVIRONMENT_CONFIGURATION=$APP_ENVIRONMENT_CONFIGURATION \
--env APP_SERVER_PORT=$APP_SERVER_PORT \
--publish $HOST_PORT:$APP_SERVER_PORT \
--rm \
--detach \
--name $IMAGE_NAME \
$IMAGE_NAME:$COMMIT_SHA_TAG
