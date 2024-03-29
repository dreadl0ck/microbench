/*
 * MICROBENCH - A testbed for comparing microvm technologies
 * Copyright (c) 2019 Philipp Mieden and Philippe Partarrieu
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package main

import "flag"

var (
	flagInteractive = flag.Bool("i", false, "interactive mode")
	flagIP          = flag.String("ip", "10.10.10.2", "guest ip")
	flagGateway     = flag.String("gw", "10.10.10.1", "gateway ip")
	flagJailUser    = flag.String("jail-user", "", "user name for jail and owner of the rootfs")

	flagCreateFS = flag.Bool("createfs", false, "create rootfs and exit")
	flagRootFS   = flag.String("rootfs", "/tmp/rootfs.ext4", "use rootfs at the specified path")

	flagKernel = flag.String("kernel", "vmlinuz", "kernel to use")
	flagMulti  = flag.Bool("multi", false, "spawn multiple vms as specified in config file")

	flagEngineType     = flag.String("engine", "firecracker", "set engine type")
	flagNumRepetitions = flag.Int("num", 1, "set number of repetitions")
	flagVersion        = flag.Bool("version", false, "print microbench version and exit")

	flagNumVMs = flag.Int("numVMs", 10, "number of vms for multi mode")
	flagTag    = flag.String("tag", "", "add custom tag to experiment logs")

	flagQEMUEmulatedCPU        = flag.Bool("qemu-cpu-emulated", false, "use emulated cpu instead of host one for qemu")
	flagFirecrackerCPUTemplate = flag.String("firecracker-cpu-template", "T2", "set CPU template to use for firecracker")

	flagNumCPUs    = flag.Int("cpus", 2, "set num of CPUs for each VM")
	flagMemorySize = flag.Int("mem", 512, "set memory in MB for each VM")

	flagDebug = flag.Bool("debug", false, "toggle debug mode")

	flagExecFile      = flag.String("exec-file", "/usr/bin/firecracker", "path to firecracker binary")
	flagUID           = flag.Int("uid", 1000, "user id for jailed user")
	flagGID           = flag.Int("gid", 1000, "group id for jailed user")
	flagChrootBaseDir = flag.String("chroot-base-dir", "/srv/jailer", "path to jail")
	flagJail          = flag.String("jailer", "/usr/bin/jailer", "path to jailer binary")
	flagNode          = flag.Int("node", 0, "jailer NUMA Cpu node")
)
