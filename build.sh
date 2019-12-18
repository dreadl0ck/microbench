#!/usr/bin/env bash

set -e

echo "building qemubench tool with QEMU support"
GOOS=linux go build -ldflags="-X main.EngineType=qemu" -o bin/qemubench -i github.com/dreadl0ck/firebench/cmd

echo "building firebench tool with Firecracker support"
GOOS=linux go build -ldflags="-X main.EngineType=fc" -o bin/firebench -i github.com/dreadl0ck/firebench/cmd
