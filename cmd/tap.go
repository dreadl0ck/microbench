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
		os.ExpandEnv("$HOME/go/src/github.com/dreadl0ck/firebench/cli/create_tap.sh"),
		*flagGateway,
	)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		log.Fatal("failed to setup tap interface: ", err)
	}
}
