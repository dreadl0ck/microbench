package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
)

func stopVM(ip net.IP, cmd *exec.Cmd) {
	// trigger VM shutdown
	http.Get("http://" + ip.String() + "/shutdown")

	fmt.Println("waiting for VM to exit")
	err := cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
