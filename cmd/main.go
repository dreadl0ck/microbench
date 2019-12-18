package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
)

var l = logrus.New()
const version = "v1.0"

func main() {

	flag.Parse()

	if *flagVersion {
		fmt.Println(version)
		return
	}

	l.Info("EngineType: ", *flagEngineType)

	if len(*flagIP) == 0 || len(*flagGateway) == 0 {
		l.Fatal("you need to pass an IP and gateway")
	}

	if *flagCreateFS {
		createRootFS(l, *flagIP, *flagGateway, 0)
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
	l.Info("done. bye")
}