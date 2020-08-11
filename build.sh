#!/usr/bin/env bash

echo "building microbench tool"
GOOS=linux go build -o bin/microbench -i github.com/dreadl0ck/microbench/cmd
