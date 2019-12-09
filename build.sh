#!/bin/bash

if [[ "$1" == "qemu" ]]; then
  echo "building qemubench tool with QEMU support"
  GOOS=linux go build -ldflags="-X main.EngineType=qemu" -o bin/qemubench -i github.com/dreadl0ck/firebench/cmd
else
  echo "building firebench tool with Firecracker support"
  GOOS=linux go build -ldflags="-X main.EngineType=fc" -o bin/firebench -i github.com/dreadl0ck/firebench/cmd
fi
