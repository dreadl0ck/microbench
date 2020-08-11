#!/bin/bash

# for testing:

VM_ID=$1
CHROOT_PATH=/srv/jailer/firecracker/$VM_ID/root

# create chroot path
mkdir -p $CHROOT_PATH

# copy deployment file / kernel / fs
cp hello-* config.file $CHROOT_PATH

# set permissions
chmod o+x $CHROOT_PATH/hello-vmlinux.bin
chmod o+w $CHROOT_PATH/hello-rootfs.ext4

jailer --id $VM_ID --node 0 --exec-file $(which firecracker) --uid 1000 --gid 1000 -- --config-file config.file