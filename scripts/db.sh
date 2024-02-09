#!/usr/bin/env bash

docker run -it \
  -p 5432:5432 \
  -e POSTGRES_PASSWORD=postgres \
  postgres:latest
