#!/bin/bash

rm -rf /srv/jailer/firecracker

/usr/bin/jailer \
  --id b76cad58-17d9-408f-b6db-90a9610a6d7d \
  --uid 1000 \
  --gid 1000 \
  --exec-file /usr/bin/firecracker \
  --node 0 \
  --chroot-base-dir /srv/jailer \
  -- --seccomp-level 0 --api-sock /srv/jailer/firecracker/b76cad58-17d9-408f-b6db-90a9610a6d7d/root/home/pmieden/.firecracker.sock-2139-81
