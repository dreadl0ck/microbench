package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func ping(start time.Time, ip net.IP) {

	fmt.Println("measuring time until network stack at", ip, "becomes available...")

	out, err := exec.Command("ping", "-c 1", ip.String()).CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		fmt.Println("ping failed: ", err)
	} else {
		fmt.Println("Time until ping response:", time.Since(start))
	}
}

func measureBootTime(start time.Time, ip net.IP, cmd *exec.Cmd) {

	fmt.Println("measuring time until service at", ip, "becomes available...")

	var serviceDown bool

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
			fmt.Println("Time until HTTP reply from webservice:", time.Since(start))

			fmt.Println("waiting for VM to exit")
			err = cmd.Wait()
			if err != nil {
				log.Fatal(err)
			}

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