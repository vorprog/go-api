#!/bin/bash

S3_BUCKET_NAME=${1:-vorprog-logs}
SERVICE_NAME=${2:-go-server}

export META_DATA_URL="http://169.254.169.254/latest/meta-data"
export INSTANCE_ID=$(curl $META_DATA_URL/instance-id)
export LOCAL_IP=$(curl $META_DATA_URL/local-ipv4) # TODO: Where to use this?
export PUBLIC_IP=$(curl $META_DATA_URL/public-ipv4) # TODO: Where to use this?
export AVAILABILITY_ZONE=$(curl $META_DATA_URL/placement/availability-zone)
export REGION=${AVAILABILITY_ZONE::-1}

CONTAINER_ID=$(sudo docker ps -aqf "name=$SERVICE_NAME")
LOG_FILE_ITERATOR=1

LOCAL_FILE_PATH=/var/lib/docker/containers/$CONTAINER_ID/$CONTAINER_ID.json.$LOG_FILE_ITERATOR
DATE=$(date "+%m/%d/%Y")
LAST_MODIFIED_TIMESTAMP=$(date -r $LOCAL_FILE_PATH +%s) # TODO: use file "birth time" instead?

aws s3 cp $LOCAL_FILE_PATH s3://$S3_BUCKET_NAME/docker/$REGION/$SERVICE_NAME/$DATE/$INSTANCE_ID/$LAST_MODIFIED_TIMESTAMP.json.gz
