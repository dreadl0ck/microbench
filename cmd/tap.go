package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func setupTap() {

	fmt.Println("setting up tap interface...")

	cmd := exec.Command(
		"/bin/bash",
		"/home/pmieden/go/src/github.com/dreadl0ck/firebench/cli/network_setup_host.sh",
	)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		log.Fatal("failed to setup tap interface: ", err)
	}
}
