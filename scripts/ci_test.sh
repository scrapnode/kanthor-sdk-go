#!/bin/sh
set -e

export TEST_KANTHOR_SDK_HOST=${TEST_KANTHOR_SDK_HOST:-"localhost:8180"}
export TEST_KANTHOR_SDK_CREDENTIALS_FILE=${TEST_KANTHOR_SDK_CREDENTIALS_FILE:-"/tmp/workspace.credentials.plain"}
export TEST_KANTHOR_SDK_APPLICATION_FILE=${TEST_KANTHOR_SDK_APPLICATION_FILE:-"/tmp/application.plain"}
CHECKSUM_FILE=./checksum


rm -rf *.out
CHECKSUM_NEW=$(find . -type f -name '*.out' -exec sha256sum {} \; | sort -k 2 | sha256sum | cut -d  ' ' -f1)
CHECKSUM_OLD=$(cat $CHECKSUM_FILE || true)

if [ "$CHECKSUM_NEW" != "$CHECKSUM_OLD" ];
then
    go test -coverprofile cover.out
    go tool cover -html cover.out -o cover.html
    find . -type f -name '*.out' -exec sha256sum {} \; | sort -k 2 | sha256sum | cut -d  ' ' -f1 > ./checksum
fi