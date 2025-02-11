#!/usr/bin/env bash

set -e

case "$1" in
  smoke)
    go test ./...
    ;;
  race)
    go test -race ./...
    ;;
  coverage)
    go test -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out
    ;;
  *)
    echo "Usage: ./test.sh [smoke|race|coverage]" && exit 1
    ;;
esac

pytest -n auto .

yarn test
