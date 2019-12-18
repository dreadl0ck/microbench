package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/dreadl0ck/firebench/stats"
	"io/ioutil"
	"net"
	"net/http"
	"os/exec"
	"strconv"
	"time"
)

func ping(l *logrus.Logger, start time.Time, ip net.IP) {

	l.WithField("ip", ip).Info("measuring time until network stack becomes available...")

	out, err := exec.Command("ping", "-c 1", ip.String()).CombinedOutput()
	if err != nil {
		l.Info(string(out))
		l.WithError(err).Info("ping failed")
	} else {
		l.WithField("delta", time.Since(start)).Info("received ping response")
	}
}

func measureWebserviceTime(l *logrus.Logger, start time.Time, ip net.IP, cmd *exec.Cmd) {

	l.WithField("ip", ip).Info("measuring time until web service becomes reachable...")

	var serviceDown bool

	for {

		//fmt.Print("CHECKING... ")

		http.DefaultClient = &http.Client{
			Timeout: 100 * time.Millisecond,
		}

		resp, err := http.Get("http://" + ip.String())
		if err != nil || resp.StatusCode != http.StatusOK {
			//l.Info(err)
			if !serviceDown {
				start = time.Now()
				serviceDown = true
				l.Info("SERVICE DOWN")
				time.Sleep(10 * time.Millisecond)
			}
			continue
		}

		l.Info("SERVICE UP:", resp.Status)

		// check if the service became reachable again
		if serviceDown && resp.StatusCode == http.StatusOK {

			serviceDown = false
			l.WithField("delta", time.Since(start)).Info("time until HTTP reply from webservice")

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
				break
			}
		}

		time.Sleep(1000 * time.Millisecond)
	}
}

func measureResponseTime(l *logrus.Logger, ip net.IP, requests int) {

	l.WithField("ip", ip).Info("executing apache bench")

	out, err := exec.Command("ab",
		"-n"+strconv.Itoa(requests),
		"-k",
		"-e", "./logs/apache/responseTime/" + ip.String() + ".log",
		"http://"+ip.String()+":80"+"/ping",
	).CombinedOutput()
	if err != nil {
		l.Info(string(out))
		l.WithError(err).Info("apache bench failed")
	} else {
		//l.Info(string(out))
	}
}

func measureThroughput(l *logrus.Logger, ip net.IP, requests int, concurrentRequests int, timeInSeconds int) {

	l.WithField("ip", ip).Info("measuring throughput...")

	out, err := exec.Command("ab",
		"-n"+strconv.Itoa(requests),
		"-k",
		"-t",strconv.Itoa(timeInSeconds),
		 "-c",strconv.Itoa(concurrentRequests),
		"-e", "./logs/apache/throughput/" + ip.String() + ".log",
		"http://"+ip.String()+":80"+"/ping",
	).CombinedOutput()
	if err != nil {
		l.Info(string(out))
		l.WithError(err).Info("apache bench failed")
	} else {
		//l.Info(string(out))
	}
}

func startHashingFile(l *logrus.Logger, ip net.IP)  {
	http.DefaultClient = &http.Client{
		Timeout: 0,
	}

	resp, err := http.Get("http://" + ip.String() + "/hashFile")
	if err != nil {
		l.WithError(err).Info("hashing failed")
	} else {
		resp, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			l.WithError(err).Fatal("failed to read response body")
		}
		l.Info(string(resp))
	}
}

func startHashingLoop(l *logrus.Logger, ip net.IP)  {

	http.DefaultClient = &http.Client{
		Timeout: 0,
	}

	resp, err := http.Get("http://" + ip.String() + "/hashLoop")
	if err != nil {
		l.WithError(err).Info("hashing failed")
	} else {
		resp, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			l.WithError(err).Fatal("failed to read response body")
		}
		l.Info(string(resp))
	}
}
