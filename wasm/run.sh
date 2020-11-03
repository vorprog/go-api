#!/usr/bin/env bash
set -xe

docker run \
--publish 8080:80 \
--detach \
--tty \
--interactive \
--rm \
wasm_go
