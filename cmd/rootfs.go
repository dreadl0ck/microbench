package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func createRootFS() {

	fmt.Println("creating rootfs for ip", *flagIP, "and gateway", *flagGateway)

	cmd := exec.Command(
		"/bin/bash",
		os.ExpandEnv("$HOME/go/src/github.com/dreadl0ck/firebench/cli/create_rootfs.sh"),
		*flagIP,
		*flagGateway,
	)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		log.Fatal("failed to setup rootfs: ", err)
	}
}
