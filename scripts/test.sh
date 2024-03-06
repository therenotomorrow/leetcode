#!/usr/bin/env bash

set -e

go test ./...

pytest solutions/python/

npm run test
