/*
 * MICROBENCH - A testbed for comparing microvm technologies
 * Copyright (c) 2019 Philipp Mieden and Philippe Partarrieu
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

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
