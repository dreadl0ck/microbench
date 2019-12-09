#!/bin/bash

if [ $1 == "qemu" ]; then
  GOOS=linux go build -o bin/firebench -i github.com/dreadl0ck/firebench/cmd -ldflags="-X main.EngineType=qemu"
else
  GOOS=linux go build -o bin/firebench -i github.com/dreadl0ck/firebench/cmd -ldflags="-X main.EngineType=fc"
fi
