#!/usr/bin/env bash
set -xe

sudo docker pull fluent/fluentd:latest
sudo docker build --tag fluentd-s3 .
