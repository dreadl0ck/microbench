package main

import (
	"github.com/sirupsen/logrus"
	"net"
	"strconv"
	"time"
)

func initVM(l *logrus.Logger, ipAddr, gwAddr string, num int) {

	ip := net.ParseIP(ipAddr)
	if ip == nil {
		l.Fatal("invalid ip: ", ipAddr)
	}

	// setup tap interface
	setupTap(gwAddr, num)

	var ether string

	ifaces, err := net.Interfaces()
	if err != nil {
		l.Fatal("failed to read interfaces")
	}

	// get hardware addr of tap interface
	for _, i := range ifaces {
		if i.Name == "tap"+strconv.Itoa(num) {
			ether = i.HardwareAddr.String()
		}
	}

	l.WithFields(logrus.Fields{
		"tap": num,
		"ether": ether,
	}).Info("configured tap interface")

	// start VMx
	cmd, err := spawnMicroVM(ether, num)
	if err != nil {
		l.WithError(err).Fatal("failed to start microVM")
	}
	l.WithField("pid", cmd.Process.Pid).Info("VM started")

	if *flagInteractive {
		l.Info("waiting for VM to exit")
		err = cmd.Wait()
		if err != nil {
			l.WithError(err).Fatal("failed to wait for VM")
		}
	} else {
		start := time.Now()
		measureWebserviceTime(l, start, ip, cmd)
		//measureResponseTime(l, ip, 1000)
		//measureThroughput(l, ip, 500000, 5, 30)
		//startHashingFile(l, ip)
		startHashingLoop(l, ip)
		stopVM(l, ip, cmd)
	}
}
