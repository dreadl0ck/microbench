# Memory benchmark

To compare the memory performance we decided to run a computationally intensive task such as hashing. A 1Gb file was created and then hashed using `sha256sum`. The same `random.txt` file is used in both tests. It is copied over during the rootfs creation. The file contains pseudo random numbers and was created by running the following command:

```bash
# the block size is 1024 * 1024
dd if=/dev/urandom of=random.txt count=1024 bs=1048576
```

## QEMU

### Launch command
```bash
root@moscow:~/go/src/github.com/dreadl0ck/firebench# sudo qemu-system-x86_64 -M microvm,rtc=off \
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
```

### Results
```bash
(none):~# time sha256sum /random.txt
84549ba4553e9897c38062ede225c315d97bfbde06016470c174ba09f514fe64  /random.txt
real    0m 19.14s
user    0m 17.97s
sys     0m 1.16s
(none):~# time sha256sum /random.txt
[  699.426481] random: crng init done
84549ba4553e9897c38062ede225c315d97bfbde06016470c174ba09f514fe64  /random.txt
real    0m 18.89s
user    0m 17.78s
sys     0m 1.11s
(none):~# time sha256sum /random.txt
84549ba4553e9897c38062ede225c315d97bfbde06016470c174ba09f514fe64  /random.txt
real    0m 18.92s
user    0m 17.88s
sys     0m 1.04s
```

## Firecracker

### Launch command
```bash
root@moscow:~/go/src/github.com/dreadl0ck/firebench# /root/go/bin/firectl \
  --kernel=/root/vmlinux \
  --root-drive=/tmp/rootfs.ext4 \
  -t \
  --cpu-template=T2 \
  --log-level=Debug \
  --firecracker-log=firecracker-vmm.log \
  --kernel-opts='console=ttyS0 noapic reboot=k panic=1 pci=off nomodules rw' \
  --tap-device=tap0/2e:60:bf:c7:64:88
```

### Results
```bash
(none):/# time sha256sum random.txt
84549ba4553e9897c38062ede225c315d97bfbde06016470c174ba09f514fe64  random.txt
real    0m 19.82s
user    0m 17.61s
sys     0m 2.08s
(none):/# time sha256sum random.txt
[  219.262005] random: crng init done
84549ba4553e9897c38062ede225c315d97bfbde06016470c174ba09f514fe64  random.txt
real    0m 19.37s
user    0m 17.72s
sys     0m 1.42s
(none):/# time sha256sum random.txt
84549ba4553e9897c38062ede225c315d97bfbde06016470c174ba09f514fe64  random.txt
real    0m 19.40s
user    0m 17.76s
sys     0m 1.40s
(none):/#
```
