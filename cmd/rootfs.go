package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"strconv"
)

func createRootFS(l *logrus.Logger, ip, gw string, num int) {

	l.WithFields(logrus.Fields{
		"num":     num,
		"ip":      ip,
		"gateway": gw,
	}).Info("creating rootfs")

	cmd := exec.Command(
		"/bin/bash",
		os.ExpandEnv("$HOME/go/src/github.com/dreadl0ck/microbench/cli/create_rootfs.sh"),
		ip,
		gw,
		strconv.Itoa(num),
	)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		l.WithError(err).Fatal("failed to setup rootfs")
	}
}
