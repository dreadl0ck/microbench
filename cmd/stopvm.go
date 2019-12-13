package main

import (
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"os/exec"
)

func stopVM(l *logrus.Logger, ip net.IP, cmd *exec.Cmd) {
	// trigger VM shutdown
	http.Get("http://" + ip.String() + "/shutdown")

	l.Info("waiting for VM to exit")
	err := cmd.Wait()
	if err != nil {
		l.Fatal(err)
	}
}
