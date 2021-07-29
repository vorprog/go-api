#!/usr/bin/env bash
set -xe

sudo docker pull fluent/fluentd:v1.13.0-debian-arm64-1.0
sudo docker build --tag fluentd-s3 .
