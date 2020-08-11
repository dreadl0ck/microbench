package main

import (
	"crypto/sha256"
	"encoding/json"
	"github.com/dustin/go-humanize"
	"github.com/ncw/directio"
	"github.com/sirupsen/logrus"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"
)

var statsHandler = func(w http.ResponseWriter, r *http.Request) {

	c, err := exec.Command("dmesg").CombinedOutput()
	if err != nil {
		l.WithError(err).Error("Error on read")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	stats, err := parseKernelLog(c)
	if err != nil {
		l.WithError(err).Error("failed to parse kernel logs")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(stats)
	if err != nil {
		l.WithError(err).Error("failed to marshal stats")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		l.WithError(err).Info("failed to write data")
	}

}

var serveHexdump = func(w http.ResponseWriter, r *http.Request) {

	in, err := directio.OpenFile("/etc/hexdump", os.O_RDONLY, 0666)
	if err != nil {
		l.WithError(err).Error("Error on open")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	block := directio.AlignedBlock(20 * directio.BlockSize)

	start := time.Now()
	n, err := io.ReadFull(in, block)
	if err != nil {
		l.WithError(err).Info("Error on read")
	}

	l.WithFields(logrus.Fields{
		"numBytes": n,
		"delta":    time.Since(start),
	}).Info("read data from file")

	_, err = w.Write(block)
	if err != nil {
		l.WithError(err).Error("failed to write data")
	}
}

var shutdown = func(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("bye"))
	l.Info("bye")

	// firecracker does not implement power handling
	// so issuing a reboot will result in the microVM being shut down
	l.Info(exec.Command("reboot").Run())
}

var hashFileHandler = func(w http.ResponseWriter, r *http.Request) {

	s, err := os.Stat("/random.data")
	if err != nil {
		l.WithError(err).Fatal("failed to stat file")
	}

	out, err := exec.Command("time", "sha256sum", "/random.data").CombinedOutput()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("\nfailed to hash file: " + err.Error() + "\n"))
	}

	_, err = w.Write(append(out, []byte("\nrandom data size: "+humanize.Bytes(uint64(s.Size())))...))
	if err != nil {
		l.WithError(err).Info("failed to write data")
	}
}

var hashLoopHandler = func(w http.ResponseWriter, r *http.Request) {

	var (
		numBytes      = 1024 * 1024
		numIterations = 100
		d             = make([]byte, numBytes)
	)

	_, err := rand.Read(d)
	if err != nil {
		l.WithError(err).Fatal("failed to read random data")
	}

	l.Info("starting hash loop")

	start := time.Now()
	for i := 0; i < numIterations; i++ {
		h := sha256.New()
		h.Sum(d)
	}

	res := "hash loop benchmark: delta=" + time.Since(start).String() + " " + strconv.Itoa(numBytes) + " bytes of random data hashed with sha256 " + strconv.Itoa(numIterations) + " times."
	l.Info("result: ", res)

	_, err = w.Write([]byte(res))
	if err != nil {
		l.WithError(err).Info("failed to write data")
	}
}
