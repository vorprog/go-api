#!/bin/bash

docker run -d \
--env AWS_REGION=$AWS_REGION \
--env S3_BUCKET_NAME=$S3_BUCKET_NAME \
--env SERVICE_NAME=$SERVICE_NAME \
--env UPLOAD_INTERVAL=$UPLOAD_INTERVAL
--publish 24224:24224 \
--publish 24224:24224/udp \
--volume $(pwd)/fluent.conf:/fluentd/etc/fluent.conf \
fluentd-s3:latest fluentd --config /fluentd/etc/fluent.conf --verbose

# --volume /data:/fluentd/log \
