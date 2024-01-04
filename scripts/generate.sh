#!/bin/sh
set -e

openapi-generator-cli generate -i openapi.json \
    -g go \
    -c openapi-generator-config.json \
    -o internal/openapi \
    --ignore-file-override .openapi-generator-ignore