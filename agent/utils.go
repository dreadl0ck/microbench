package main

import (
	"github.com/dreadl0ck/firebench/stats"
	"regexp"
	"strings"
	"time"
)

var tsRegExp = regexp.MustCompile("[0-9]*\\.[0-9]*")

func parseKernelLog(contents []byte) (s *stats.Summary, err error) {

	s = &stats.Summary{}

	lines := strings.Split(string(contents), "\n")
	for i, line := range lines {
		if strings.Contains(line, "random: fast init done") {

			// ex: [   17.567752] random: fast init done
			duration := strings.TrimSpace(tsRegExp.FindString(lines[i])) + "s"
			l.Info("found duration:", duration)

			d, err := time.ParseDuration(duration)
			if err != nil {
				return nil, err
			}

			s.KernelBootup = d
			s.KernelLogLines = i+1
			break
		}
	}
	return
}
