package main

import (
	"github.com/sirupsen/logrus"
	"sync"
)

func runMulti() {

	var (
		wg sync.WaitGroup
		wgRootFS sync.WaitGroup
	)

	for num, cfg := range parseConfig().Vms {

		wg.Add(1)
		wgRootFS.Add(1)

		l, f := makeLogger(cfg.IP)
		defer f.Close()

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
			wgRootFS.Done()
			wgRootFS.Wait()
			initVM(l, ip, gw, n)
			wg.Done()
		}()
	}

	l.Info("waiting...")
	wg.Wait()
}