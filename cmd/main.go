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
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var logger = logrus.New()

const version = "v1.0"

func main() {
	flag.Parse()

	if *flagVersion {
		fmt.Println(version)

		return
	}

	jailUser := *flagJailUser
	if jailUser == "" {
		jailUser = os.Getenv("JAIL_USER")
	}

	if *flagCreateFS {
		createRootFS(logger, *flagIP, *flagGateway, 0, jailUser)

		return
	}

	logger.WithFields(logrus.Fields{
		"engine": *flagEngineType,
		"tag":    *flagTag,
		"num":    *flagNumRepetitions,
		"numVMs": *flagNumVMs,
		"multi":  *flagMulti,
		"cpu":    *flagNumCPUs,
		"mem":    *flagMemorySize,
	}).Info("starting microbench ", version)

	if len(*flagIP) == 0 || len(*flagGateway) == 0 {
		logger.Fatal("you need to pass an IP and gateway")
	}

	if !*flagMulti {
		var count int

		for {
			count++

			l, cleanup := makeLogger(*flagIP + "-" + strconv.Itoa(count))
			defer cleanup()

			initVM(l, *flagIP, *flagGateway, 0)

			if count == *flagNumRepetitions {
				break
			}
		}

		return
	}

	runMulti(jailUser)
	logger.Info("done. bye")
}
