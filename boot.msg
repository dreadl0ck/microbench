[    0.000000] Initializing cgroup subsys cpuset
[    0.000000] Initializing cgroup subsys cpu
[    0.000000] Linux version 3.0.101-108.84-default (geeko@buildhost) (gcc     version 4.3.4 [gcc-4_3-branch revision 152973] (SUSE Linux) ) #1 SMP Fri Nov 30 15:57:27 UTC 2018 (7a72692)
[    0.000000] Command line: BOOT_IMAGE=dev000:\EFI\SUSE\vmlinuz-    3.0.101-108.84-default root=/dev/disk/by-id/scsi-35000cca070168a20-part2  splash=verbose showopts
[    0.000000] x86/fpu: Using 'eager' FPU context switches.
[    0.000000] BIOS-provided physical RAM map:
[    0.000000]  BIOS-e820: 0000000000000000 - 00000000000a0000 (usable)
[    0.000000]  BIOS-e820: 0000000000100000 - 00000000792de000 (usable)
[    0.000000]  BIOS-e820: 00000000792de000 - 00000000798f4000 (reserved)
[    4.967393] Brought up 128 CPUs
[    4.967400] Total of 128 processors activated (512000.20 BogoMIPS).
[    5.377931] devtmpfs: initialized
[    5.420574] PM: Registering ACPI NVS region at 79a38000 (5976064 bytes)
[    5.421208] print_constraints: dummy:
[    5.421243] Time: 20:15:03  Date: 01/22/19
[    5.421817] NET: Registered protocol family 16
[    5.422067] ACPI FADT declares the system doesn't support PCIe ASPM, so disable it
[    5.422073] ACPI: bus type pci registered
[   14.053705] igb 0000:42:00.3: added PHC on eth3
[   14.066858] igb 0000:42:00.3: Intel(R) Gigabit Ethernet Network Connection
[   14.080159] igb 0000:42:00.3: eth3: (PCIe:5.0Gb/s:Width x4) 0c:c4:7a:3a:51:33
[   14.093493] igb 0000:42:00.3: eth3: PBA No: 010A00-000
[   14.106628] igb 0000:42:00.3: Using MSI-X interrupts. 8 rx queue(s), 8 tx queue(s)
[   14.800728] device-mapper: uevent: version 1.0.3
[   14.814242] device-mapper: ioctl: 4.25.0-ioctl (2012-07-25) initialised: dm-devel@redhat.com
[   15.254103] loop: module loaded
[   15.634412] SGI XFS with ACLs, security attributes, realtime, large block/inode numbers, no debug enabled
[   15.648644] SGI XFS Quota Management subsystem
[   15.661127] XFS (sda1): Mounting Filesystem
[   15.790512] XFS (sda1): Ending clean mount
[   15.802635] XFS (sdb1): Mounting Filesystem
[   15.886148] XFS (sdb1): Ending clean mount
[   15.898303] XFS (sdd1): Mounting Filesystem
[   16.010051] XFS (sdd1): Ending clean mount
[   17.567752] fuse init (API version 7.16)
Kernel logging (ksyslog) stopped.
Kernel log daemon terminating.

Waiting for device /dev/disk/by-id/scsi-35000cca070168a20-part2 to appear:  ok
fsck from util-linux 2.19.1
[/sbin/fsck.ext3 (1) -- /] fsck.ext3 -a -C0 /dev/sdc2
myhostname: clean, 1172494/36618240 files, 42233936/146465024 blocks
fsck succeeded. Mounting root device read-write.