package main

import "flag"

var (
	flagInteractive = flag.Bool("i", false, "interactive mode")
	flagIP          = flag.String("ip", "", "guest ip")

	flagCreateFS = flag.Bool("createfs", false, "create rootfs and exit")
	flagRootFS   = flag.String("rootfs", "/tmp/rootfs.ext4", "use rootfs at the specified path")

	flagTap = flag.Bool("tap", true, "create tap device")
)
