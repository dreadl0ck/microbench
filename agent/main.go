package main

import (
	"fmt"
	"log"
	"net/http"
)

const addr = ":80"

func main() {

	fmt.Println("serving at", addr)

	http.HandleFunc("/", serveHexdump)
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/stats", statsHandler)

	log.Fatal(
		"failed to serve: ",
		http.ListenAndServe(addr, nil),
	)
}
