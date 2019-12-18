#!/usr/bin/env bash

ERROR_MESSAGE='usage: ./run-qemu.sh -k "/root/vmlinux" -r "/tmp/rootfs.ext4" -i "tap0"'

while getopts :k:r:i: option; do
	case ${option} in
		k) KERNEL=${OPTARG};;
		r) ROOTFS=${OPTARG};;
		i) INTERFACE=${OPTARG};;
		\?) echo -e "$ERROR_MESSAGE";;
	esac
done

echo $KERNEL
echo $ROOTFS
echo $INTERFACE

function runQemu() {
	sudo qemu-system-x86_64 -M microvm,rtc=off \
		-enable-kvm \
		-smp 2 \
		-m 1g \
		-kernel "$KERNEL" \
		-append "earlyprintk=ttyS0 console=ttyS0 root=/dev/vda" \
		-nodefaults \
		-no-user-config \
		-nographic \
		-serial stdio \
		-drive id=test,file="$ROOTFS",format=raw,if=none \
		-device virtio-blk-device,drive=test \
		-netdev tap,id="$INTERFACE",ifname="$INTERFACE",script=no,downscript=no \
		-device virtio-net-device,netdev="$INTERFACE" \
		-no-reboot
}

runQemu
