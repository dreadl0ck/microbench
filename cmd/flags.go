package main

import "flag"

var (
	flagInteractive = flag.Bool("i", false, "interactive mode")
	flagIP          = flag.String("ip", "10.10.10.2", "guest ip")
	flagGateway     = flag.String("gw", "10.10.10.1", "gateway ip")

	flagCreateFS = flag.Bool("createfs", false, "create rootfs and exit")
	flagRootFS   = flag.String("rootfs", "/tmp/rootfs.ext4", "use rootfs at the specified path")
	//flagKernel = flag.String("kernel", "$HOME/hello-vmlinux.bin", "kernel to use")
	flagKernel = flag.String("kernel", "$HOME/vmlinuz", "kernel to use")
	flagMulti  = flag.Bool("multi", false, "spawn multiple vms as specified in config file")
)
