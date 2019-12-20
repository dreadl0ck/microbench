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

num=$3

# hotfix for dangling network interfaces
#systemctl restart docker

# bootstrap filesystem
if mount | grep -q "/tmp/my-rootfs$num"; then
	umount -f /tmp/my-rootfs$num
fi

dd if=/dev/zero of=/tmp/rootfs$num.ext4 bs=1M count=100
mkfs.ext4 /tmp/rootfs$num.ext4

# mount it
mkdir -p /tmp/my-rootfs$num
mount /tmp/rootfs$num.ext4 /tmp/my-rootfs$num

# add firebench-agent
mkdir -p /tmp/my-rootfs$num/usr/bin
# link static against musl for running in alpine
go build --ldflags '-linkmode external -extldflags "-static"' -o /tmp/my-rootfs$num/usr/bin/firebench-agent -i github.com/dreadl0ck/firebench/agent

# copy init script(s)
cp $HOME/go/src/github.com/dreadl0ck/firebench/cli/init_alpine.sh /tmp/my-rootfs$num/init_alpine.sh
cp $HOME/go/src/github.com/dreadl0ck/firebench/bin/networking /tmp/my-rootfs$num/networking
#cp $HOME/go/src/github.com/dreadl0ck/firebench/random.data /tmp/my-rootfs$num/random.data

# run docker container with latest alpine image to populate filesystem
if [ "$1" == "-i" ]; then
    echo "starting interactive mode"
    # interactive mode
    docker run -it --rm -v /tmp/my-rootfs$num:/my-rootfs alpine
else
    echo "running init_alpine"
    # run init script and exit
    docker run --rm -v /tmp/my-rootfs$num:/my-rootfs alpine ash /my-rootfs/init_alpine.sh "$vm_ip" "$gw_ip"
fi

# sync & eject
echo "syncing..."
sync

echo "unmounting..."
umount /tmp/my-rootfs$num

echo "created rootfs at /tmp/rootfs$num.ext4"
exit 0