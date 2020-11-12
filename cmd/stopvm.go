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
	"net"
	"net/http"
	"os/exec"
	"time"
)

func stopVM(l *logrus.Logger, ip net.IP, cmd *exec.Cmd) {

	start := time.Now()

	// trigger VM shutdown
	http.Get("http://" + ip.String() + "/shutdown")

	l.Info("waiting for VM to exit")
	err := cmd.Wait()
	if err != nil {
		l.WithError(err).Fatal("failed to wait for VM")
	}

	l.WithField("delta", time.Since(start)).Info("shutdown complete")
}
