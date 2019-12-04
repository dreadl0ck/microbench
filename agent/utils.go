package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type stats struct {
	KernelBootup time.Duration `json:"kernelBootup"`
	ServiceInit  time.Duration `json:"serviceInit"`
}

var tsRegExp = regexp.MustCompile("[0-9]*\\.[0-9]*")

func parseKernelLog(contents []byte) (s *stats, err error) {

	s = &stats{}

	lines := strings.Split(string(contents), "\n")
	for i, line := range lines {
		if strings.Contains(line, "random: fast init done") {

			// ex: [   17.567752] random: fast init done
			duration := strings.TrimSpace(tsRegExp.FindString(lines[i-1])) + "s"
			fmt.Println("found duration:", duration)

			d, err := time.ParseDuration(duration)
			if err != nil {
				return nil, err
			}

			s.KernelBootup = d
			break
		}
	}
	return
}
