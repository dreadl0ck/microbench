package main

import (
	"fmt"
	"strings"
	"time"
)

type stats struct {
	KernelBootup time.Duration `json:"kernelBootup"`
	ServiceInit  time.Duration `json:"serviceInit"`
}

// TODO implement
func parseKernelLog(contents []byte) (*stats, error) {

	lines := strings.Split(string(contents), "\n")
	for i, line := range lines {
		if strings.Contains(line, "Kernel logging (ksyslog) stopped") {
			// ex: [   17.567752] fuse init (API version 7.16)
			tsLine := lines[i-1]
			fmt.Println(tsLine)
		}
	}
	return nil, nil
}
