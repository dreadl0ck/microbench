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
