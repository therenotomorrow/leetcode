#!/usr/bin/env bash

golangci-lint run

go test -count 1 -race -coverprofile cover.out -cover ./...
go tool cover -func cover.out | grep total
