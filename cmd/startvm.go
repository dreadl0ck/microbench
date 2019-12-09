package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func spawnMicroVM(tapEther string) (*exec.Cmd, error) {

	fmt.Println("spawning VM")

	var cmd *exec.Cmd

	switch EngineType {
	case "fc":
		cmd = exec.Command(
			os.ExpandEnv("$HOME/go/bin/firectl"),
			"--kernel=" + os.ExpandEnv(*flagKernel),
			"--root-drive="+*flagRootFS,
			"-t",
			"--cpu-template=T2",
			"--log-level=Debug",
			"--firecracker-log=firecracker-vmm.log",
			"--kernel-opts='console=ttyS0 noapic reboot=k panic=1 pci=off nomodules rw'",
			//"--metadata='{"foo":"bar"}' ""
			"--tap-device=tap0/"+tapEther,
		)
	case "qemu":
		cmd = exec.Command(
			os.ExpandEnv("$HOME/go/bin/firectl"),
			"--kernel=" + os.ExpandEnv(*flagKernel),
			"--root-drive="+*flagRootFS,
			"-t",
			"--cpu-template=T2",
			"--log-level=Debug",
			"--firecracker-log=firecracker-vmm.log",
			"--kernel-opts='console=ttyS0 noapic reboot=k panic=1 pci=off nomodules rw'",
			//"--metadata='{"foo":"bar"}' ""
			"--tap-device=tap0/"+tapEther,
		)
	default:
		log.Fatal("invalid engine type: ", EngineType)
	}

	// TODO: potentially fucks up the terminal when the firecracker process exits
	if *flagInteractive {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd, cmd.Start()
}
