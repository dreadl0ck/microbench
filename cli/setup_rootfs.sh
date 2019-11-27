#!/bin/bash

# hotfix for dangling network interfaces
systemctl restart docker

# bootstrap filesystem
umount -f /tmp/my-rootfs
dd if=/dev/zero of=/tmp/rootfs.ext4 bs=1M count=250
mkfs.ext4 /tmp/rootfs.ext4

# mount it
mkdir -p /tmp/my-rootfs
mount /tmp/rootfs.ext4 /tmp/my-rootfs

# add direct-fs
mkdir -p /tmp/my-rootfs/usr/bin
# link static against musl for running in alpine
go build --ldflags '-linkmode external -extldflags "-static"' -o /tmp/my-rootfs/usr/bin/direct-fs -i github.com/dreadl0ck/firebench/direct-fs

# copy init script
cp cli/init_alpine.sh /tmp/my-rootfs/init_alpine.sh

# run docker container with latest alpine image to populate filesystem
if [ "$1" == "-i" ]; then
    # interactive mode
    docker run -it --rm -v /tmp/my-rootfs:/my-rootfs alpine
else
    # run init script and exit
    docker run -it --rm -v /tmp/my-rootfs:/my-rootfs alpine ash /my-rootfs/init_alpine.sh
fi

# sync & eject
sync
umount /tmp/my-rootfs

exit 0