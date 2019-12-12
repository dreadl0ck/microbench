package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func createRootFS(ip, gw string, num int) {

	fmt.Println("creating rootfs for ip", ip, "and gateway", gw)

	cmd := exec.Command(
		"/bin/bash",
		os.ExpandEnv("$HOME/go/src/github.com/dreadl0ck/firebench/cli/create_rootfs.sh"),
		ip,
		gw,
		strconv.Itoa(num),
	)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		log.Fatal("failed to setup rootfs: ", err)
	}
}
