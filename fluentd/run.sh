#!/bin/bash

S3_BUCKET_NAME=${1:-vorprog-logs}
SERVICE_NAME=${2:-go-server}

export META_DATA_URL="http://169.254.169.254/latest/meta-data"
export INSTANCE_ID=$(curl $META_DATA_URL/instance-id)
export LOCAL_IP=$(curl $META_DATA_URL/local-ipv4) # TODO: Where to use this?
export PUBLIC_IP=$(curl $META_DATA_URL/public-ipv4) # TODO: Where to use this?
export AVAILABILITY_ZONE=$(curl $META_DATA_URL/placement/availability-zone)
export AWS_REGION=${AVAILABILITY_ZONE::-1}

sudo docker run -d \
--env AWS_REGION=$AWS_REGION \
--env S3_BUCKET_NAME=$S3_BUCKET_NAME \
--env SERVICE_NAME=$SERVICE_NAME \
--publish 24224:24224 \
--publish 24224:24224/udp \
--volume $(pwd)/fluent.conf:/fluentd/etc/fluent.conf \
--name fluentd \
fluentd-s3:latest fluentd --config /fluentd/etc/fluent.conf --verbose
