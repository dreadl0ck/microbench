#!/bin/bash

# usage: ./create_rootfs.sh <ip> <gw> <num> [<jail_user>]"

# exit script on error
set -e

vm_ip="$1"
gw_ip="$2"
num=$3
jail_user=$4

if [ -z "$vm_ip" ]; then
    echo "you must pass an ip for the vm as parameter #1"
    echo "usage: ./create_rootfs.sh <ip> <gw> <num> [<jail_user>]"
    echo "                          ^^^^"
    exit 1
fi

if [ -z "$gw_ip" ]; then
    echo "you must pass a gateway ip as parameter #2"
    echo "usage: ./create_rootfs.sh <ip> <gw> <num> [<jail_user>]"
    echo "                               ^^^^"
    exit 1
fi

if [ -z "$num" ]; then
    echo "you must pass an index number for the tap interface as parameter #3"
    echo "usage: ./create_rootfs.sh <ip> <gw> <num> [<jail_user>]"
    echo "                                    ^^^^"
    exit 1
fi

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

# add microbench-agent
mkdir -p /tmp/my-rootfs$num/usr/bin
# link static against musl for running in alpine
go build --ldflags '-linkmode external -extldflags "-static"' -o /tmp/my-rootfs$num/usr/bin/microbench-agent -i github.com/dreadl0ck/microbench/agent

echo "GOPATH: ${GOPATH}"

# copy init script(s)
cp "${GOPATH}/src/github.com/dreadl0ck/microbench/cli/init_alpine.sh" /tmp/my-rootfs$num/init_alpine.sh
cp "${GOPATH}/src/github.com/dreadl0ck/microbench/bin/networking" /tmp/my-rootfs$num/networking
#cp ${GOPATH}/src/github.com/dreadl0ck/microbench/random.data /tmp/my-rootfs$num/random.data

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

echo "jail user: $jail_user"
if [ -n "$jail_user" ]; then
    echo "setting jail user to $jail_user"
    chown "$jail_user" "/tmp/rootfs$num.ext4"
    ls -la "/tmp/rootfs$num.ext4"
fi

echo "created rootfs at /tmp/rootfs$num.ext4"
exit 0