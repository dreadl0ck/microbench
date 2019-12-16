package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"io"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

var EngineType string

var l = logrus.New()

func main() {

	flag.Parse()

	l.Info("EngineType:", EngineType)

	if len(*flagIP) == 0 || len(*flagGateway) == 0 {
		l.Fatal("you need to pass an IP and gateway")
	}

	if *flagCreateFS {
		createRootFS(l, *flagIP, *flagGateway, 0)
		os.Exit(0)
	}

	if !*flagMulti {
		initVM(l, *flagIP, *flagGateway, 0)
		os.Exit(0)
	}

	var wg sync.WaitGroup

	for num, cfg := range parseConfig().Vms {

		wg.Add(1)

		f, err := os.Create(cfg.IP + ".log")
		if err != nil {
			l.Fatal(err)
		}
		defer f.Close()

		l := logrus.New()
		l.SetOutput(io.MultiWriter(os.Stdout, f))
		l.Formatter = &logrus.TextFormatter{
			ForceColors:               true,
			FullTimestamp:             true,
			TimestampFormat:           "2 Jan 2006 15:04:05",
		}

		l.WithFields(logrus.Fields{
			"num": num,
			"ip": cfg.IP,
			"gateway": cfg.Gateway,
			"logfile": cfg.IP + ".log",
		}).Info("bootstrapping machine")

		// prevent capturing loop vars
		var (
			ip = cfg.IP
			gw = cfg.Gateway
			n  = num
		)

		go func() {
			createRootFS(l, ip, gw, n)
			initVM(l, ip, gw, n)
			wg.Done()
		}()
	}

	l.Info("waiting...")
	wg.Wait()
	l.Info("done. bye")
}

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

	// start VM
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
		go ping(l, start, ip)
		measureBootTime(l, start, ip, cmd)
		measureResponseTime(l, ip, 1000)
		startHashing(l, ip)
		stopVM(l, ip, cmd)
	}
}
