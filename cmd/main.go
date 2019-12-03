package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"time"
)

var flagInteractive = flag.Bool("i", false, "interactive mode")
var flagIP = flag.String("ip", "", "guest ip")

var flagCreateFS = flag.Bool("createfs", false, "create rootfs and exit")
var flagRootFS = flag.String("rootfs", "/tmp/rootfs.ext4", "use rootfs at the specified path")

var flagTap = flag.Bool("tap", true, "create tap device")

func main() {

	flag.Parse()

	if *flagCreateFS {
		setupRootFS()
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
		err = cmd.Wait()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		measureBootTime(ip, cmd)
	}
}

func setupRootFS() {

	fmt.Println("setting up rootfs...")

	cmd := exec.Command(
		"/bin/bash",
		"/home/pmieden/go/src/github.com/dreadl0ck/firebench/cli/setup_rootfs.sh",
	)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		log.Fatal("failed to setup rootfs: ", err)
	}
}

func setupTap() {

	fmt.Println("setting up tap interface...")

	cmd := exec.Command(
		"/bin/bash",
		"/home/pmieden/go/src/github.com/dreadl0ck/firebench/cli/network_setup_host.sh",
	)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		log.Fatal("failed to setup tap interface: ", err)
	}
}

func measureBootTime(ip net.IP, cmd *exec.Cmd) {

	fmt.Println("measuring time until service at", ip, "becomes available...")

	var (
		start       time.Time
		serviceDown bool
	)

	for {

		//fmt.Print("CHECKING... ")

		http.DefaultClient = &http.Client{
			Timeout: 10 * time.Millisecond,
		}

		resp, err := http.Get("http://" + ip.String())
		if err != nil || resp.StatusCode != http.StatusOK {
			//fmt.Println(err)
			if !serviceDown {
				start = time.Now()
				serviceDown = true
				fmt.Println("SERVICE DOWN:", start)
				time.Sleep(10 * time.Millisecond)
			}
			continue
		}

		fmt.Println("SERVICE UP:", resp.Status)

		// check if the service became reachable again
		if serviceDown && resp.StatusCode == http.StatusOK {

			serviceDown = false
			fmt.Println("DELTA:", time.Since(start))

			fmt.Println("waiting for VM to exit")
			time.Sleep(5 * time.Second)

			/*fmt.Println("killing firecracker process:", cmd.Process.Pid)
			err := cmd.Process.Signal(syscall.SIGTERM)
			if err != nil {
				fmt.Println("failed to kill firecracker process:", err)
			}*/

			os.Exit(0)
		}

		time.Sleep(1000 * time.Millisecond)
	}
}

func spawnMicroVM(tapEther string) (*exec.Cmd, error) {

	fmt.Println("spawning VM")

	cmd := exec.Command(
		"/home/pmieden/go/bin/firectl",
		"--kernel=/home/pmieden/hello-vmlinux.bin",
		"--root-drive=" + *flagRootFS,
		"-t",
		"--cpu-template=T2",
		"--log-level=Debug",
		"--firecracker-log=firecracker-vmm.log",
		"--kernel-opts='console=ttyS0 noapic reboot=k panic=1 pci=off nomodules rw'",
		//"--metadata='{"foo":"bar"}' ""
		"--tap-device=tap0/" + tapEther,
	)

	// TODO: potentially fucks up the terminal when the firecracker process exits
	if *flagInteractive {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd, cmd.Start()
}