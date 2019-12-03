package main

import (
	"fmt"
	"os"
	"os/exec"
)

func spawnMicroVM(tapEther string) (*exec.Cmd, error) {

	fmt.Println("spawning VM")

	cmd := exec.Command(
		"/home/pmieden/go/bin/firectl",
		"--kernel=/home/pmieden/hello-vmlinux.bin",
		"--root-drive="+*flagRootFS,
		"-t",
		"--cpu-template=T2",
		"--log-level=Debug",
		"--firecracker-log=firecracker-vmm.log",
		"--kernel-opts='console=ttyS0 noapic reboot=k panic=1 pci=off nomodules rw'",
		//"--metadata='{"foo":"bar"}' ""
		"--tap-device=tap0/"+tapEther,
	)

	// TODO: potentially fucks up the terminal when the firecracker process exits
	if *flagInteractive {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd, cmd.Start()
}
