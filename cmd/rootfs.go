package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func setupRootFS() {

	fmt.Println("setting up rootfs...")

	cmd := exec.Command(
		"/bin/bash",
		"/home/pmieden/go/src/github.com/dreadl0ck/firebench/cli/setup_rootfs.sh",
	)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		log.Fatal("failed to setup rootfs: ", err)
	}
}
