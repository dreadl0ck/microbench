#!/bin/bash

# QEMU with host cpu
# re-create the rootfs to be sure everything works
bin/firebench -createfs
# run sequential
bin/firebench -engine qemu -num 10 -tag "sequential"

# run concurrent
bin/firebench -engine qemu -multi -numVMs 10
bin/firebench -engine qemu -multi -numVMs 20

# QEMU use emulated cpu
bin/firebench -createfs
bin/firebench -engine qemu -num 10 -qemu-cpu-emulated -tag "sequential-emulated-cpu"
bin/firebench -engine qemu -multi -qemu-cpu-emulated -numVMs 10 -tag "emulated-cpu"
bin/firebench -engine qemu -multi -qemu-cpu-emulated -numVMs 20 -tag "emulated-cpu"

# Firecracker
# re-create the rootfs to be sure everything works
bin/firebench -createfs
# run sequential
bin/firebench -num 10 -tag "sequential"
# run concurrent
bin/firebench -multi -numVMs 10
bin/firebench -multi -numVMs 20

# run with default kernel
bin/firebench -createfs
# run sequential
bin/firebench -num 10 -kernel /root/hello-vmlinux.bin -tag "sequential-default-kernel"

# run concurrent
bin/firebench -multi -numVMs 10 -kernel /root/hello-vmlinux.bin -tag "default-kernel"
bin/firebench -multi -numVMs 20 -kernel /root/hello-vmlinux.bin -tag "default-kernel"

# Firecracker with C3 CPU template
bin/firebench -createfs
# run sequential with C3 CPU template
bin/firebench -num 10 -tag "sequential-C3" -firecracker-cpu-template "C3"
# run concurrent with C3 CPU template
bin/firebench -multi -numVMs 10 -tag "C3" -firecracker-cpu-template "C3"
bin/firebench -multi -numVMs 20 -tag "C3" -firecracker-cpu-template "C3"

# run with default kernel and C3 CPU template
bin/firebench -createfs
# run sequential with default kernel and C3 CPU template
bin/firebench -num 10 -kernel /root/hello-vmlinux.bin -tag "sequential-default-kernel-C3" -firecracker-cpu-template "C3"

# run concurrent with default kernel and C3 CPU template
bin/firebench -multi -numVMs 10 -kernel /root/hello-vmlinux.bin -tag "default-kernel-C3" -firecracker-cpu-template "C3"
bin/firebench -multi -numVMs 20 -kernel /root/hello-vmlinux.bin -tag "default-kernel-C3" -firecracker-cpu-template "C3"

tree experiment_logs
echo "done."