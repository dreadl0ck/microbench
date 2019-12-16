package main

import (
	"os"
	"os/exec"
	"strconv"
)

func setupTap(address string, num int) {

	l.Info("setting up tap interface...")

	cmd := exec.Command(
		"/bin/bash",
		os.ExpandEnv("$HOME/go/src/github.com/dreadl0ck/firebench/cli/create_tap.sh"),
		address,
		strconv.Itoa(num),
	)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		l.WithError(err).Fatal("failed to setup tap interface")
	}
}
