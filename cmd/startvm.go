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
	l.WithFields(logrus.Fields{
		"rootfs": rootfs,
		"ether": tapEther,
	}).Info("spawning microVM")

	switch EngineType {
	case "fc":
		cmd = exec.Command(
			os.ExpandEnv("$HOME/go/bin/firectl"),
			"--kernel="+os.ExpandEnv(*flagKernel),
			"--root-drive="+rootfs,
			"-t",
			"--cpu-template=T2",
			"--log-level=Debug",
			"--firecracker-log=firecracker-vmm.log",
			"--kernel-opts='console=ttyS0 noapic reboot=k panic=1 pci=off nomodules rw'",
			//"--metadata='{"foo":"bar"}' ""
			"--tap-device=tap"+strconv.Itoa(num)+"/"+tapEther,
		)
	case "qemu":
		cmd = exec.Command(
			os.ExpandEnv("$HOME/go/src/github.com/dreadl0ck/firebench/cli/run-microvm.sh"),
		)
	default:
		l.Fatal("invalid engine type: ", EngineType)
	}

	if *flagInteractive {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd, cmd.Start()
}
