package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

var EngineType string

func main() {

	flag.Parse()

	fmt.Println("EngineType:", EngineType)

	if len(*flagIP) == 0 || len(*flagGateway) == 0 {
		log.Fatal("you need to pass an IP and gateway")
	}

	if *flagCreateFS {
		createRootFS(*flagIP, *flagGateway, 0)
		os.Exit(0)
	}

	if !*flagMulti {
		initVM(*flagIP, *flagGateway, 0)
		os.Exit(0)
	}

	var wg sync.WaitGroup

	for num, cfg := range parseConfig().Vms {
		wg.Add(1)
		fmt.Printf("bootstrapping machine #%d (IP: %s, GW: %s)\n", num, cfg.IP, cfg.Gateway)

		// prevent capturing loop vars
		var (
			ip = cfg.IP
			gw = cfg.Gateway
			n  = num
		)

		go func() {
			createRootFS(ip, gw, n)
			initVM(ip, gw, n)
			wg.Done()
		}()
	}

	fmt.Println("waiting...")
	wg.Wait()
	fmt.Println("done. bye")
}

func initVM(ipAddr, gwAddr string, num int) {

	ip := net.ParseIP(ipAddr)
	if ip == nil {
		log.Fatal("invalid ip: ", ipAddr)
	}

	// setup tap interface
	setupTap(gwAddr, num)

	var ether string

	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatal("failed to read interfaces")
	}

	// get hardware addr of tap interface
	for _, i := range ifaces {
		if i.Name == "tap"+strconv.Itoa(num) {
			ether = i.HardwareAddr.String()
		}
	}

	fmt.Println("tap ether:", ether)

	// start VM
	cmd, err := spawnMicroVM(ether, num)
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
