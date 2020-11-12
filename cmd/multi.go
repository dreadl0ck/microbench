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
	"strconv"
	"sync"
)

func makeLocalAddrPair(num int) (string, string) {
	return "10.0." + strconv.Itoa(num) + ".1", "10.0." + strconv.Itoa(num) + ".2"
}

func runMulti(jailUser string) {

	var (
		wg       sync.WaitGroup
		wgRootFS sync.WaitGroup
	)

	for i := 1; i <= *flagNumVMs; i++ {

		ipAddr, gateway := makeLocalAddrPair(i)

		wg.Add(1)
		wgRootFS.Add(1)

		l, cleanup := makeLogger(ipAddr)
		defer cleanup()

		l.WithFields(logrus.Fields{
			"num":     i,
			"ip":      ipAddr,
			"gateway": gateway,
			"logfile": ipAddr + ".log",
		}).Info("bootstrapping machine")

		// prevent capturing loop vars
		var (
			n = i
		)

		go func() {
			createRootFS(l, ipAddr, gateway, n, jailUser)
			wgRootFS.Done()
			wgRootFS.Wait()
			initVM(l, ipAddr, gateway, n)
			wg.Done()
		}()
	}

	logger.Info("waiting...")
	wg.Wait()
}
