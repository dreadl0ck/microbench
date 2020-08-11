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
