#!/bin/bash

# exit script on error
set -e

vm_ip="$1"
gw_ip="$2"

if [ -z "$vm_ip" ]; then
    echo "you must pass an ip for the vm as parameter #1"
    echo "usage: ./create_rootfs.sh <ip> <gw>"
    exit 1
fi

if [ -z "$gw_ip" ]; then
    echo "you must pass a gateway ip as parameter #2"
    echo "usage: ./create_rootfs.sh <ip> <gw>"
    exit 1
fi

# hotfix for dangling network interfaces
systemctl restart docker

# bootstrap filesystem
if mount | grep -q "/tmp/my-rootfs"; then
	umount -f /tmp/my-rootfs
fi
dd if=/dev/zero of=/tmp/rootfs.ext4 bs=1M count=4500
mkfs.ext4 /tmp/rootfs.ext4

# mount it
mkdir -p /tmp/my-rootfs
mount /tmp/rootfs.ext4 /tmp/my-rootfs

# add firebench-agent
mkdir -p /tmp/my-rootfs/usr/bin
# link static against musl for running in alpine
go build --ldflags '-linkmode external -extldflags "-static"' -o /tmp/my-rootfs/usr/bin/firebench-agent -i github.com/dreadl0ck/firebench/agent

# copy init script(s)
cp $HOME/go/src/github.com/dreadl0ck/firebench/cli/init_alpine.sh /tmp/my-rootfs/init_alpine.sh
cp $HOME/go/src/github.com/dreadl0ck/firebench/bin/networking /tmp/my-rootfs/networking
cp $HOME/go/src/github.com/dreadl0ck/firebench/bin/kcbench /tmp/my-rootfs/kcbench
cp -r $HOME/linux.git /tmp/my-rootfs/linux.git

# run docker container with latest alpine image to populate filesystem
if [ "$1" == "-i" ]; then
    echo "starting interactive mode"
    # interactive mode
    docker run -it --rm -v /tmp/my-rootfs:/my-rootfs alpine
else
    echo "running init_alpine"
    # run init script and exit
    docker run --rm -v /tmp/my-rootfs:/my-rootfs alpine ash /my-rootfs/init_alpine.sh "$vm_ip" "$gw_ip"
fi

# sync & eject
echo "syncing..."
sync

echo "unmounting..."
umount /tmp/my-rootfs

echo "created rootfs at /tmp/rootfs.ext4"
exit 0