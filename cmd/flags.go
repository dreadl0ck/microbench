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

	flagEngineType = flag.String("engine", "firecracker", "set engine type")
	flagNumRepetitions = flag.Int("num", 1, "set number of repetitions")
	flagVersion = flag.Bool("version", false, "print firebench version and exit")

	flagNumVMs = flag.Int("numVMs", 10, "number of vms for multi mode")
	flagTag = flag.String("tag", "", "add custom tag to experiment logs")

	flagQEMUEmulatedCPU = flag.Bool("qemu-cpu-emulated", false, "use emulated cpu instead of host one for qemu")
	flagFirecrackerCPUTemplate = flag.String("firecracker-cpu-template", "T2", "set CPU template to use for firecracker")

	flagNumCPUs = flag.Int("cpus", 2, "set num of CPUs for each VM")
	flagMemorySize = flag.Int("mem", 512, "set memory in MB for each VM")

	flagDebug = flag.Bool("debug", false, "toggle debug mode")
)

