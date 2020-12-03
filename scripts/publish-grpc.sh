#!/bin/sh

set -e

[ -z "$DEBUG" ] || set -x

echo "\n===> Generate image...\n"

docker build --no-cache -f Dockerfile.grpc -t calculator .

echo "\n===> Docker tag...\n"

docker tag calculator fernandocagale/calculator:grpc

echo "\n===> Docker push...\n"

docker push fernandocagale/calculator:grpc