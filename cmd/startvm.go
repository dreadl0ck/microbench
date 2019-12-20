package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"strconv"
)

func spawnMicroVM(tapEther string, num int) (*exec.Cmd, error) {

	var (
		cmd    *exec.Cmd
		rootfs = "/tmp/rootfs" + strconv.Itoa(num) + ".ext4"
	)

	switch *flagEngineType {
	case "firecracker":
		cmd = exec.Command(
			os.ExpandEnv("$HOME/go/bin/firectl"),
			"--kernel="+os.ExpandEnv(*flagKernel),
			"--root-drive="+rootfs,
			"-t",
			"--cpu-template="+*flagFirecrackerCPUTemplate,
			"-c="+strconv.Itoa(*flagNumCPUs),
			"-m="+strconv.Itoa(*flagMemorySize),
			"--log-level=Debug",
			"--firecracker-log=firecracker-vmm.log",
			"--kernel-opts='console=ttyS0 noapic reboot=k panic=1 pci=off nomodules rw'",
			"--tap-device=tap"+strconv.Itoa(num)+"/"+tapEther,
		)
	case "qemu":
		if *flagQEMUEmulatedCPU {
			cmd = exec.Command(
				os.ExpandEnv("$HOME/go/src/github.com/dreadl0ck/firebench/cli/run-qemu-microvm-emulated-cpu.sh"),
				"-k", os.ExpandEnv(*flagKernel),
				"-r", rootfs,
				"-i", "tap" + strconv.Itoa(num),
				"-c", strconv.Itoa(*flagNumCPUs),
				"-m", strconv.Itoa(*flagMemorySize),
			)
		} else {
			cmd = exec.Command(
				os.ExpandEnv("$HOME/go/src/github.com/dreadl0ck/firebench/cli/run-qemu-microvm.sh"),
				"-k", os.ExpandEnv(*flagKernel),
				"-r", rootfs,
				"-i", "tap" + strconv.Itoa(num),
				"-c", strconv.Itoa(*flagNumCPUs),
				"-m", strconv.Itoa(*flagMemorySize),
			)
		}

	default:
		l.Fatal("invalid engine type: ", *flagEngineType)
	}

	if *flagInteractive {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	l.WithFields(logrus.Fields{
		"rootfs": rootfs,
		"ether": tapEther,
		"path": cmd.Path,
		"args": cmd.Args,
	}).Info("spawning microVM")

	return cmd, cmd.Start()
}
