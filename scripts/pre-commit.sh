#!/usr/bin/env bash

set -ex

./scripts/code.sh
./scripts/test.sh

git add .

ln -sf ../../scripts/pre-commit.sh .git/hooks/pre-commit
