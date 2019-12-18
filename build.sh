#!/bin/bash

echo "building firebench tool"
GOOS=linux go build -o bin/firebench -i github.com/dreadl0ck/firebench/cmd
