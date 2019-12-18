package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func makeLogger(name string) (*logrus.Logger, *os.File) {

	const logDir = "experiment_logs"

	// ignore error from creating logDir
	// we dont care if it exists already
	os.MkdirAll(logDir, 755)

	f, err := os.Create(filepath.Join(logDir, name + ".log"))
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

	return l, f
}