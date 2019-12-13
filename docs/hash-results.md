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
  --tap-device=tap0/2e:60:bf:c7:64:88 \
  -m 1024
```

### Results
```bash
(none):~# time sha256sum /random.txt
84549ba4553e9897c38062ede225c315d97bfbde06016470c174ba09f514fe64  /random.txt
real    0m 20.36s
user    0m 17.72s
sys     0m 2.61s
(none):~# time sha256sum /random.txt
[   56.912356] random: crng init done
84549ba4553e9897c38062ede225c315d97bfbde06016470c174ba09f514fe64  /random.txt
real    0m 19.30s
user    0m 17.63s
sys     0m 1.52s
(none):~# time sha256sum /random.txt
84549ba4553e9897c38062ede225c315d97bfbde06016470c174ba09f514fe64  /random.txt
real    0m 19.28s
user    0m 17.58s
sys     0m 1.55s
```
### Firecracker mem

```bash
(none):~# cat /proc/meminfo
MemTotal:        1012792 kB
MemFree:           73836 kB
MemAvailable:     871912 kB
Buffers:             688 kB
Cached:           923436 kB
SwapCached:            0 kB
Active:             8932 kB
Inactive:         916472 kB
Active(anon):       1304 kB
Inactive(anon):       68 kB
Active(file):       7628 kB
Inactive(file):   916404 kB
Unevictable:           0 kB
Mlocked:               0 kB
SwapTotal:             0 kB
SwapFree:              0 kB
Dirty:                 0 kB
Writeback:             0 kB
AnonPages:          1312 kB
Mapped:             5568 kB
Shmem:                96 kB
KReclaimable:       3600 kB
Slab:               7972 kB
SReclaimable:       3600 kB
SUnreclaim:         4372 kB
KernelStack:         812 kB
PageTables:          232 kB
NFS_Unstable:          0 kB
Bounce:                0 kB
WritebackTmp:          0 kB
CommitLimit:      506396 kB
Committed_AS:     135344 kB
VmallocTotal:   34359738367 kB
VmallocUsed:        1072 kB
VmallocChunk:          0 kB
Percpu:              184 kB
AnonHugePages:         0 kB
ShmemHugePages:        0 kB
ShmemPmdMapped:        0 kB
HugePages_Total:       0
HugePages_Free:        0
HugePages_Rsvd:        0
HugePages_Surp:        0
Hugepagesize:       2048 kB
Hugetlb:               0 kB
DirectMap4k:       18432 kB
DirectMap2M:     1030144 kB
```


### QEMU mem

```bash
(none):~# cat /proc/meminfo
MemTotal:        1012620 kB
MemFree:          983832 kB
MemAvailable:     917588 kB
Buffers:             684 kB
Cached:             8400 kB
SwapCached:            0 kB
Active:             4352 kB
Inactive:           6004 kB
Active(anon):       1308 kB
Inactive(anon):       68 kB
Active(file):       3044 kB
Inactive(file):     5936 kB
Unevictable:           0 kB
Mlocked:               0 kB
SwapTotal:             0 kB
SwapFree:              0 kB
Dirty:                36 kB
Writeback:             0 kB
AnonPages:          1332 kB
Mapped:             5684 kB
Shmem:                96 kB
KReclaimable:       1408 kB
Slab:               6888 kB
SReclaimable:       1408 kB
SUnreclaim:         5480 kB
KernelStack:         928 kB
PageTables:          244 kB
NFS_Unstable:          0 kB
Bounce:                0 kB
WritebackTmp:          0 kB
CommitLimit:      506308 kB
Committed_AS:     136184 kB
VmallocTotal:   34359738367 kB
VmallocUsed:        1144 kB
VmallocChunk:          0 kB
Percpu:              408 kB
AnonHugePages:         0 kB
ShmemHugePages:        0 kB
ShmemPmdMapped:        0 kB
HugePages_Total:       0
HugePages_Free:        0
HugePages_Rsvd:        0
HugePages_Surp:        0
Hugepagesize:       2048 kB
Hugetlb:               0 kB
DirectMap4k:       18432 kB
DirectMap2M:     1030144 kB
```
