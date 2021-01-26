#!/bin/bash

# Fail on errors
set -e
# Print commands
set -x

protoc -I"./proto" --go_out="$GOPATH/src" ./proto/*.proto
