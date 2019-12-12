package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os/exec"
	"strconv"
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
			Timeout: 50 * time.Millisecond,
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

			// retrieve VM stats
			resp, err := http.Get("http://" + ip.String() + "/stats")
			if err != nil {
				fmt.Println(err)
			} else {
				statsData, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println("statsData:", string(statsData))
				break
			}
		}

		time.Sleep(1000 * time.Millisecond)
	}
}

func measureResponseTime(ip net.IP, requests int) {
	fmt.Println("measuring response time...")

	out, err := exec.Command("ab",
		"-n"+strconv.Itoa(requests),
		"-k",
		"http://"+ip.String()+":80"+"/ping",
	).CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		fmt.Println("ab failed: ", err)
	} else {
		fmt.Println(string(out))
	}
}

func measureThroughput(ip net.IP, filepath string) {
	fmt.Println("measuring throughput by uploading file: ", filepath)

	// TODO: get throughput metrics from postFile
	postFile(ip, filepath)
}

func startCompilation(ip net.IP) {
	http.DefaultClient = &http.Client{
		Timeout: 50 * time.Millisecond,
	}

	resp, err := http.Get("http://" + ip.String() + "/compile")
	if err != nil {
		fmt.Println("compilation err: " + err.Error())
	} else {
		resp, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(resp))
	}
}
