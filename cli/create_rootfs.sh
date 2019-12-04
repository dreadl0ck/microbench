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

# add firebench-agent
mkdir -p /tmp/my-rootfs/usr/bin
# link static against musl for running in alpine
go build --ldflags '-linkmode external -extldflags "-static"' -o /tmp/my-rootfs/usr/bin/firebench-agent -i github.com/dreadl0ck/firebench/agent

# copy init script(s)
cp /home/pmieden/go/src/github.com/dreadl0ck/firebench/cli/init_alpine.sh /tmp/my-rootfs/init_alpine.sh
cp /home/pmieden/go/src/github.com/dreadl0ck/firebench/bin/networking /tmp/my-rootfs/networking

# run docker container with latest alpine image to populate filesystem
if [ "$1" == "-i" ]; then
    echo "starting interactive mode"
    # interactive mode
    docker run -it --rm -v /tmp/my-rootfs:/my-rootfs alpine
else
    echo "running init_alpine"
    # run init script and exit
    docker run --rm -v /tmp/my-rootfs:/my-rootfs alpine ash /my-rootfs/init_alpine.sh $1 $2
fi

# sync & eject
echo "syncing..."
sync

echo "unmounting..."
umount /tmp/my-rootfs

echo "created rootfs at /tmp/rootfs.ext4"
exit 0