package main

import (
	uuid "github.com/kevinburke/go.uuid"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

func spawnMicroVM(l *logrus.Logger, tapEther string, num int) (*exec.Cmd, error) {

	var (
		cmd    *exec.Cmd
		rootfs = "/tmp/rootfs" + strconv.Itoa(num) + ".ext4"
	)

	switch *flagEngineType {
	case "firecracker":
		cmd = exec.Command(
			os.ExpandEnv("${GOPATH}/bin/firectl"),
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
			"--id="+uuid.NewV4().String(),
			"--exec-file="+*flagExecFile,
			"--uid="+strconv.Itoa(*flagUID),
			"--gid="+strconv.Itoa(*flagGID),
			"--debug",
			"--chroot-base-dir="+*flagChrootBaseDir,
			"--jailer="+*flagJail,
			"--node="+strconv.Itoa(*flagNode),
		)
	case "qemu":
		if *flagQEMUEmulatedCPU {
			cmd = exec.Command(
				os.ExpandEnv("${GOPATH}/src/github.com/dreadl0ck/microbench/cli/run-qemu-microvm-emulated-cpu.sh"),
				"-k", os.ExpandEnv(*flagKernel),
				"-r", rootfs,
				"-i", "tap"+strconv.Itoa(num),
				"-c", strconv.Itoa(*flagNumCPUs),
				"-m", strconv.Itoa(*flagMemorySize),
			)
		} else {
			cmd = exec.Command(
				os.ExpandEnv("${GOPATH}/src/github.com/dreadl0ck/microbench/cli/run-qemu-microvm.sh"),
				"-k", os.ExpandEnv(*flagKernel),
				"-r", rootfs,
				"-i", "tap"+strconv.Itoa(num),
				"-c", strconv.Itoa(*flagNumCPUs),
				"-m", strconv.Itoa(*flagMemorySize),
			)
		}

		// jail the QEMU process if a user and group id have been supplied
		if *flagUID != 0 && *flagGID != 0 {
			cmd.SysProcAttr = &syscall.SysProcAttr{
				Credential: &syscall.Credential{
					Uid: uint32(*flagUID),
					Gid: uint32(*flagGID),
				},
			}
		}

	default:
		l.Fatal("invalid engine type: ", *flagEngineType)
	}

	if *flagInteractive {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else if *flagDebug {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	// dont duplicate this log message in debug mode
	if !*flagDebug {
		// log to global logger
		logger.WithFields(logrus.Fields{
			"rootfs": rootfs,
			"ether":  tapEther,
			"path":   cmd.Path,
			"args":   cmd.Args,
		}).Info("spawning microVM")
	}

	// log to vm specific logger
	l.WithFields(logrus.Fields{
		"rootfs": rootfs,
		"ether":  tapEther,
		"path":   cmd.Path,
		"args":   cmd.Args,
	}).Info("spawning microVM")

	return cmd, cmd.Start()
}
