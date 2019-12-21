package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
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

	logger.WithFields(logrus.Fields{
		"engine": *flagEngineType,
		"tag": *flagTag,
		"num": *flagNumRepetitions,
		"numVMs": *flagNumVMs,
		"multi": *flagMulti,
		"cpu": *flagNumCPUs,
		"mem": *flagMemorySize,
	}).Info("starting firebench")

	if len(*flagIP) == 0 || len(*flagGateway) == 0 {
		logger.Fatal("you need to pass an IP and gateway")
	}

	if *flagCreateFS {
		createRootFS(logger, *flagIP, *flagGateway, 0)
		return
	}

	if !*flagMulti {

		var count int
		for {
			count++

			l, cleanup := makeLogger(*flagIP+"-"+strconv.Itoa(count))
			defer cleanup()

			initVM(l, *flagIP, *flagGateway, 0)
			if count == *flagNumRepetitions {
				break
			}
		}
		return
	}

	runMulti()
	logger.Info("done. bye")
}