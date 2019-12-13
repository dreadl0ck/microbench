package main

import (
	"encoding/json"
	"github.com/dustin/go-humanize"
	"github.com/ncw/directio"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"
)

var statsHandler = func(w http.ResponseWriter, r *http.Request) {

	c, err := exec.Command("dmesg").CombinedOutput()
	if err != nil {
		l.Error("Error on read: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	stats, err := parseKernelLog(c)
	if err != nil {
		l.Error("failde to parse kernel logs: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(stats)
	if err != nil {
		l.Error("failed to marshal stats: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		l.Info("failed to write data: ", err)
	}

}

var serveHexdump = func(w http.ResponseWriter, r *http.Request) {

	in, err := directio.OpenFile("/etc/hexdump", os.O_RDONLY, 0666)
	if err != nil {
		l.Error("Error on open: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	block := directio.AlignedBlock(20 * directio.BlockSize)

	start := time.Now()
	n, err := io.ReadFull(in, block)
	if err != nil {
		l.Info("Error on read: ", err)
	}

	l.Info("read", n, "bytes from file, in", time.Since(start))

	_, err = w.Write(block)
	if err != nil {
		l.Info("failed to write data: ", err)
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

var hashHandler = func(w http.ResponseWriter, r *http.Request) {

	s, err := os.Stat("/random.data")
	if err != nil {
		l.Fatal(err)
	}

	out, err := exec.Command("time", "sha256sum", "/random.data").CombinedOutput()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("\nfailed to hash file: " + err.Error() + "\n"))
	}

	_, err = w.Write(append(out, []byte("\nrandom data size:" + humanize.Bytes(uint64(s.Size())))...))
	if err != nil {
		l.Info("failed to write data: ", err)
	}
}
