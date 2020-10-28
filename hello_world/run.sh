#!/usr/bin/env bash
set -x

docker build --quiet --tag hello_world_go:latest .

docker run \
--env PORT=8080 \
-d \
--tty \
--interactive \
hello_world_go
