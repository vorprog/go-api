#!/usr/bin/env bash
set -xe

export PORT=8080

docker run \
--env PORT=$PORT \
--publish $PORT:$PORT \
--detach \
--tty \
--interactive \
hello_world_go
