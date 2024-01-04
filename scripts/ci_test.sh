#!/bin/sh
set -e

CHECKSUM_FILE=./checksum

CHECKSUM_NEW=$(find . -type f -name '*.out' -exec sha256sum {} \; | sort -k 2 | sha256sum | cut -d  ' ' -f1)
CHECKSUM_OLD=$(cat $CHECKSUM_FILE || true)

if [ "$CHECKSUM_NEW" != "$CHECKSUM_OLD" ];
then
    go test -coverprofile cover.out
    go tool cover -html cover.out -o cover.html
    find . -type f -name '*.out' -exec sha256sum {} \; | sort -k 2 | sha256sum | cut -d  ' ' -f1 > ./checksum
fi