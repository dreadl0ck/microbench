package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ncw/directio"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		in, err := directio.OpenFile("/etc/hexdump", os.O_RDONLY, 0666)
		if err != nil {
			fmt.Println("Error on open: ", err)
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
	})

	log.Fatal(
		"failed to serve: ",
		http.ListenAndServe(":80", nil),
	)
}
