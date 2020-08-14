package main

import (
	"os"
	"os/exec"
	"strconv"
)

func setupTap(address string, num int) {

	logger.Info("setting up tap interface... GOPATH=", os.ExpandEnv("${GOPATH}"))

	cmd := exec.Command(
		"/bin/bash",
		os.ExpandEnv("${GOPATH}/src/github.com/dreadl0ck/microbench/cli/create_tap.sh"),
		address,
		strconv.Itoa(num),
	)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		logger.WithError(err).Fatal("failed to setup tap interface")
	}
}
