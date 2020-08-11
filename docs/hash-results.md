# Memory benchmark

To compare the memory performance we decided to run a computationally intensive task such as hashing. A 512 MB file was created and then hashed using `sha256sum`. The same `random.txt` file is used in both tests. It is copied over during the rootfs creation. The file contains pseudo random numbers and was created by running the following command:

```bash
# the block size is 1024 * 1024
dd if=/dev/urandom of=random.txt count=512 bs=1048576
```

## QEMU

### Launch command
```bash
root@moscow:~/go/src/github.com/dreadl0ck/microbench# sudo qemu-system-x86_64 -M microvm,rtc=off \
   -enable-kvm \
   -smp 2 \
   -m 1g \
   -kernel /root/vmlinuz \
   -append "earlyprintk=ttyS0 console=ttyS0 root=/dev/vda" \
   -nodefaults \
   -no-user-config \
   -cpu host \
   -nographic \
   -serial stdio \
   -drive id=test,file=/tmp/rootfs0.ext4,format=raw,if=none \
   -device virtio-blk-device,drive=test \
   -netdev tap,id=tap0,ifname=tap0,script=no,downscript=no \
   -device virtio-net-device,netdev=tap0 \
   -no-reboot
```

### Results
```bash
(none):~# time sha256sum /random.data
5f1f8917758bd02e89605a55ce53fbcf64c2f5bf7ef6386b1de35d93be56ac41  /random.data
real	0m 9.69s
user	0m 8.87s
sys	0m 0.73s
(none):~# time sha256sum /random.data
5f1f8917758bd02e89605a55ce53fbcf64c2f5bf7ef6386b1de35d93be56ac41  /random.data
real	0m 9.15s
user	0m 8.80s
sys	0m 0.35s
(none):~# time sha256sum /random.data
5f1f8917758bd02e89605a55ce53fbcf64c2f5bf7ef6386b1de35d93be56ac41  /random.data
real	0m 9.11s
user	0m 8.88s
sys	0m 0.23s
(none):~# du -h /random.data
512.0M	/random.data
```

## Firecracker

### Launch command
```bash
root@moscow:~/go/src/github.com/dreadl0ck/microbench# firectl \
  --kernel=/root/vmlinuz \
  --root-drive=/tmp/rootfs0.ext4 \
  -t \
  --cpu-template=T2 \
  --log-level=Debug \
  --firecracker-log=firecracker-vmm.log \
  --kernel-opts='console=ttyS0 noapic reboot=k panic=1 pci=off nomodules rw' \
  --tap-device=tap0/f2:09:6a:ad:81:32
```

### Results
```bash
(none):~# time sha256sum /random.data
5f1f8917758bd02e89605a55ce53fbcf64c2f5bf7ef6386b1de35d93be56ac41  /random.data
real	0m 10.14s
user	0m 8.96s
sys	0m 1.16s
(none):~# time sha256sum /random.data
5f1f8917758bd02e89605a55ce53fbcf64c2f5bf7ef6386b1de35d93be56ac41  /random.data
real	0m 9.67s
user	0m 8.85s
sys	0m 0.71s
(none):~# time sha256sum /random.data
5f1f8917758bd02e89605a55ce53fbcf64c2f5bf7ef6386b1de35d93be56ac41  /random.data
real	0m 9.68s
user	0m 8.83s
sys	0m 0.74s
(none):~# du -h /random.data
512.0M	/random.data
```
