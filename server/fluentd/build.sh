#!/usr/bin/env bash
set -xe

docker pull fluent/fluentd:latest
docker build --tag fluentd-s3 .
