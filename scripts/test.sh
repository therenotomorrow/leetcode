#!/usr/bin/env bash

set -e

case "$1" in
  -fast)
    go test ./...
    ;;
  -race)
    go test -race ./...
    ;;
  -cover)
    go test -race -coverprofile cover.out -cover ./...
    go tool cover -func cover.out | grep total:
    ;;
  *)
    echo "Usage: ./test.sh [-fast|-race|-cover]" && exit 1
    ;;
esac
