/*
 * MICROBENCH - A testbed for comparing microvm technologies
 * Copyright (c) 2019 Philipp Mieden and Philippe Partarrieu
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"strconv"
)

func createRootFS(l *logrus.Logger, ip, gw string, num int, jailUser string) {

	l.WithFields(logrus.Fields{
		"num":     num,
		"ip":      ip,
		"gateway": gw,
		"jailUser": jailUser,
	}).Info("creating rootfs... GOPATH=", os.ExpandEnv("${GOPATH}"))

	if jailUser == "" {
		l.Warn("no jailUser set!")
	}

	cmd := exec.Command(
		"/bin/bash",
		os.ExpandEnv("${GOPATH}/src/github.com/dreadl0ck/microbench/cli/create_rootfs.sh"),
		ip,
		gw,
		strconv.Itoa(num),
		jailUser,
	)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		l.WithError(err).Fatal("failed to setup rootfs")
	}
}
