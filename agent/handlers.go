package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
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

var uploadHandler = func(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(3000 << 20) // maxMemory 3000MB
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to parse multipart message: " + err.Error() + "\n"))
	}

	file, handler, err := r.FormFile("linux.tar.gz")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error retrieving file: " + err.Error() + "\n"))
		return
	}
	defer file.Close()

	// Create a temporary file
	tempFile, err := ioutil.TempFile("./", "linux-*.tar.gz")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error creating a temporary file: " + err.Error() + "\n"))
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into the temp file
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error reading file into memory: " + err.Error() + "\n"))
	}

	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	var resp string
	resp = "\tuploaded file: " + handler.Filename + "\n" +
		"\tfile size: " + strconv.FormatInt(handler.Size, 10) + "KiB\n"
		//"MIME Header: " + handler.Header + "\n"

	// return that we have successfully uploaded our file!
	_, err = w.Write([]byte(resp))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

var compileHandler = func(w http.ResponseWriter, r *http.Request) {
	// TODO replace ls with cd into /tmp and make
	out, err := exec.Command("find").CombinedOutput()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("\nfailed to compile linux kernel: " + err.Error() + "\n"))
		return
	}

	_, err = w.Write([]byte(out))
	if err != nil {
		fmt.Println("failed to write data: ", err)
	}
}
