#!/bin/bash

# hotfix
systemctl restart docker

umount -f /tmp/my-rootfs
dd if=/dev/zero of=/tmp/rootfs.ext4 bs=1M count=250
mkfs.ext4 /tmp/rootfs.ext4

mkdir -p /tmp/my-rootfs
mount /tmp/rootfs.ext4 /tmp/my-rootfs

mkdir -p /tmp/my-rootfs/usr/bin
cp ~/go/src/github.com/dreadl0ck/os3/ls/websrv/websrv /tmp/my-rootfs/usr/bin

cp scripts/init_alpine.sh /tmp/my-rootfs/init_alpine.sh

if [ "$1" == "-i" ]; then
    docker run -it --rm -v /tmp/my-rootfs:/my-rootfs alpine
else    
    docker run -it --rm -v /tmp/my-rootfs:/my-rootfs alpine ash /my-rootfs/init_alpine.sh
fi

sync
umount /tmp/my-rootfs

exit 0