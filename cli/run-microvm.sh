#!/bin/bash -e

#toolchain="$(uname -m)-unknown-linux-musl"
#rm -f /tmp/firecracker.socket
#firecracker/build/cargo_target/${toolchain}/debug/firecracker --config-file vm-config.json --api-sock /tmp/firecracker.socket


# firectl \
#   --kernel=$HOME/hello-vmlinux.bin \
#   --root-drive=/tmp/rootfs.ext4 -t \
#   --cpu-template=T2 \
#   --firecracker-log=firecracker-vmm.log \
#   --kernel-opts="console=ttyS0 noapic reboot=k panic=1 pci=off nomodules rw" \
#   --metadata='{"foo":"bar"}' \
#   --tap-device="tap0/$1"

sudo qemu-system-x86_64 -M microvm,rtc=off \
	-enable-kvm \
	-smp 2 \
	-m 1g \
	-kernel /root/vmlinux \
	-append "earlyprintk=ttyS0 console=ttyS0 root=/dev/vda" \
	-nodefaults \
	-no-user-config \
	-nographic \
	-serial stdio \
	-drive id=test,file=/tmp/rootfs.ext4,format=raw,if=none \
	-device virtio-blk-device,drive=test \
	-netdev tap,id=tap0,ifname=tap0,script=no,downscript=no \
	-device virtio-net-device,netdev=tap0 \
	-no-reboot
