#!/bin/bash

# Firecracker
# re-create the rootfs to be sure everything works
bin/microbench -createfs
# run sequential
bin/microbench -num 1 -tag "sequential"
