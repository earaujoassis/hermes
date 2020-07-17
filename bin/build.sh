#!/usr/bin/env bash

set -e

GOARCH=arm64 GOOS=linux go build -o hermes-arm64 main.go
