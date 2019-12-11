package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var EngineType string

func main() {

	flag.Parse()

	fmt.Println("EngineType:", EngineType)

	if *flagCreateFS {
		createRootFS()
		os.Exit(0)
	}

	fmt.Println("using rootfs from", *flagRootFS)

	if len(*flagIP) == 0 {
		log.Fatal("you need to pass an IP")
	}

	ip := net.ParseIP(*flagIP)
	if ip == nil {
		log.Fatal("invalid ip: ", *flagIP)
	}

	// setup tap interface
	if *flagTap {
		setupTap()
	}

	var ether string

	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatal("failed to read interfaces")
	}

	// get hardware addr of tap interface
	for _, i := range ifaces {
		if i.Name == "tap0" {
			ether = i.HardwareAddr.String()
		}
	}

	fmt.Println("tap ether:", ether)

	// start VM
	cmd, err := spawnMicroVM(ether)
	if err != nil {
		log.Fatal("failed to start microVM: ", err)
	}
	fmt.Println("PID:", cmd.Process.Pid)

	if *flagInteractive {
		fmt.Println("waiting for VM to exit")
		err = cmd.Wait()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		start := time.Now()
		go ping(start, ip)
		measureBootTime(start, ip, cmd)
		measureResponseTime(ip, 100)
		measureThroughput(ip, "hello.txt")
		startCompilation(ip)
		stopVM(ip, cmd)
	}
}
