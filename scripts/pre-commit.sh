#!/usr/bin/env bash

set -ex

./scripts/code.sh

git add .

ln -sf ../../scripts/pre-commit.sh .git/hooks/pre-commit
