#!/usr/bin/env bash
set -xe

docker run \
--publish 8080:80 \
--detach \
--tty \
--interactive \
--rm \
--name wasm_go \
wasm_go
