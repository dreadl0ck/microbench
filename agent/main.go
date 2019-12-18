package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

const addr = ":80"

var l = logrus.New()

func main() {

	l.Info("serving at", addr)

	http.HandleFunc("/", serveHexdump)
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/stats", statsHandler)
	http.HandleFunc("/hashFile", hashFileHandler)
	http.HandleFunc("/hashLoop", hashLoopHandler)

	l.Fatal(
		"failed to serve: ",
		http.ListenAndServe(addr, nil),
	)
}
