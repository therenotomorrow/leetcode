#!/usr/bin/env bash

set -e

gofmt -w -s .
golangci-lint run ./solutions/golang/...

isort .
black .

flake8 .
mypy .

yarn lint
