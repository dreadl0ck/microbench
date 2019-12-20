#!/bin/bash

# QEMU with host cpu
# re-create the rootfs to be sure everything works
bin/firebench -createfs
# run sequential
bin/firebench -engine qemu -num 10 -tag "sequential"

# run concurrent
bin/firebench -engine qemu -multi -numVMs 20 -tag "x20"
bin/firebench -engine qemu -multi -numVMs 50 -tag "x50"

# QEMU use emulated cpu
bin/firebench -createfs
bin/firebench -engine qemu -num 10 qemu-cpu-emulated -tag "sequential-emulated-cpu"
bin/firebench -engine qemu -multi -numVMs 20 -tag "x20-emulated-cpu"
bin/firebench -engine qemu -multi -numVMs 50 -tag "x50-emulated-cpu"

# Firecracker
# re-create the rootfs to be sure everything works
bin/firebench -createfs
# run sequential
bin/firebench -num 10 -tag "sequential"
# run concurrent
bin/firebench -multi -numVMs 20 -tag "x20"
bin/firebench -multi -numVMs 50 -tag "x50"

# run with default kernel
bin/firebench -createfs
# run sequential
bin/firebench -num 10 -kernel /root/hello-vmlinux.bin -tag "sequential-default-kernel"

# run concurrent
bin/firebench -multi -numVMs 20 -kernel /root/hello-vmlinux.bin -tag "x20-default-kernel"
bin/firebench -multi -numVMs 50 -kernel /root/hello-vmlinux.bin -tag "x50-default-kernel"

# Fireracker with C3 CPU template
bin/firebench -createfs
# run sequential
bin/firebench -num 10 -tag "sequential-C3" -firecracker-cpu-template "C3"
# run concurrent
bin/firebench -multi -numVMs 20 -tag "x20-C3" -firecracker-cpu-template "C3"
bin/firebench -multi -numVMs 50 -tag "x50-C3" -firecracker-cpu-template "C3"

# run with default kernel
bin/firebench -createfs
# run sequential
bin/firebench -num 10 -kernel /root/hello-vmlinux.bin -tag "sequential-default-kernel-C3" -firecracker-cpu-template "C3"

# run concurrent
bin/firebench -multi -numVMs 20 -kernel /root/hello-vmlinux.bin -tag "x20-default-kernel-C3" -firecracker-cpu-template "C3"
bin/firebench -multi -numVMs 50 -kernel /root/hello-vmlinux.bin -tag "x50-default-kernel-C3" -firecracker-cpu-template "C3"

tree experiments_logs
echo "done."