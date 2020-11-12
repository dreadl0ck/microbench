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
	"encoding/json"
	"github.com/dreadl0ck/microbench/stats"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

func fetchStats(l *logrus.Logger, ip net.IP) {

	// retrieve VM stats
	resp, err := http.Get("http://" + ip.String() + "/stats")
	if err != nil {
		l.Info(err)
	} else {
		statsData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			l.Fatal(err)
		}

		var s = new(stats.Summary)
		err = json.Unmarshal(statsData, &s)
		if err != nil {
			l.Fatal(err)
		}

		l.WithField("delta", s.KernelBootup).Info("kernel boot time received")
		l.WithField("lines", s.KernelLogLines).Info("number of kernel log lines received")

		if len(s.KernelLogs) > 0 {
			fileName := filepath.Join(logDir, *flagEngineType+"-dmesg.log")
			f, err := os.Create(fileName)
			if err != nil {
				l.Fatal(err)
			}
			defer f.Close()

			f.Write([]byte(s.KernelLogs))

			l.Info("wrote kernel logs to file: ", fileName)
		}

	}
}
