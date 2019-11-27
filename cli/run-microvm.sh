#!/bin/bash

#toolchain="$(uname -m)-unknown-linux-musl"
#rm -f /tmp/firecracker.socket
#firecracker/build/cargo_target/${toolchain}/debug/firecracker --config-file vm-config.json --api-sock /tmp/firecracker.socket

firectl \
  --kernel=/home/pmieden/hello-vmlinux.bin \
  --root-drive=/tmp/rootfs.ext4 -t \
  --cpu-template=T2 \
  --firecracker-log=firecracker-vmm.log \
  --kernel-opts="console=ttyS0 noapic reboot=k panic=1 pci=off nomodules rw" \
  --metadata='{"foo":"bar"}' \
  --tap-device="tap0/$1"