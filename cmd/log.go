package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const logDir = "experiment_logs"

func makeLogger(name string) (*logrus.Logger, func()) {

	// ignore error from creating logDir
	// we dont care if it exists already
	os.MkdirAll(logDir, 755)

	var fileName string
	if *flagTag != "" {
		fileName = filepath.Join(logDir, *flagEngineType + "-" + name + "-" + *flagTag + ".log")
	} else {
		fileName = filepath.Join(logDir, *flagEngineType + "-" + name + ".log")
	}

	f, err := os.Create(fileName)
	if err != nil {
		l.Fatal(err)
	}

	l := logrus.New()

	//l.SetOutput(io.MultiWriter(os.Stdout, f))
	l.SetOutput(f)

	l.Formatter = &logrus.TextFormatter{
		ForceColors:               true,
		FullTimestamp:             true,
		TimestampFormat:           "2 Jan 2006 15:04:05",
	}

	var version string

	switch *flagEngineType {
	case "firecracker":
		out, err := exec.Command("firecracker", "-V").CombinedOutput()
		if err != nil {
			l.WithError(err).Fatal("failed to get firecracker version")
		}
		version = string(out)
	case "qemu":
		out, err := exec.Command("qemu-system-x86_64", "--version").CombinedOutput()
		if err != nil {
			l.WithError(err).Fatal("failed to get firecracker version")
		}
		version = strings.Split(string(out), "\n")[0]
	}

	l.WithField("version", version).Info("using engine: ", *flagEngineType)

	hwOut, err := exec.Command("lshw", "-C", "system,memory,processor", "-short").CombinedOutput()
	if err != nil {
		l.WithError(err).Fatal("failed to get hardware infos")
	}

	l.Info("hardware: ", string(hwOut))

	unameOut, err := exec.Command("uname", "-a").CombinedOutput()
	if err != nil {
		l.WithError(err).Fatal("failed to get uname infos")
	}

	l.Info("system: ", string(unameOut))

	return l, func() {
		fmt.Println("closing file handle for logfile:", f.Name())
		f.Close()
	}
}