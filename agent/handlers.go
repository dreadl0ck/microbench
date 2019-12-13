package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"
	"github.com/ncw/directio"
)

var statsHandler = func(w http.ResponseWriter, r *http.Request) {

	c, err := exec.Command("dmesg").CombinedOutput()
	if err != nil {
		fmt.Println("Error on read:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	stats, err := parseKernelLog(c)
	if err != nil {
		fmt.Println("failde to parse kernel logs:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(stats)
	if err != nil {
		fmt.Println("failde to marshal stats:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		fmt.Println("failed to write data: ", err)
	}

}

var serveHexdump = func(w http.ResponseWriter, r *http.Request) {

	in, err := directio.OpenFile("/etc/hexdump", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Error on open: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	block := directio.AlignedBlock(20 * directio.BlockSize)

	start := time.Now()
	n, err := io.ReadFull(in, block)
	if err != nil {
		fmt.Println("Error on read: ", err)
	}

	fmt.Println("read", n, "bytes from file, in", time.Since(start))

	_, err = w.Write(block)
	if err != nil {
		fmt.Println("failed to write data: ", err)
	}
}

var shutdown = func(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("bye"))
	fmt.Println("bye")

	// firecracker does not implement power handling
	// so issuing a reboot will result in the microVM being shut down
	fmt.Println(exec.Command("reboot").Run())
}

var hashHandler = func(w http.ResponseWriter, r *http.Request) {
	out, err := exec.Command("time", "sha256sum", "/random.txt").CombinedOutput()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("\nfailed to hash file: " + err.Error() + "\n"))
	}

	_, err = w.Write([]byte(out))
	if err != nil {
		fmt.Println("failed to write data: ", err)
	}
}
