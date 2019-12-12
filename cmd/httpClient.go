package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

func postFile(ip net.IP, filename string) {
	fmt.Println("starting to post file...")

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file: " + err.Error())
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("linux.tar.gz", filepath.Base(file.Name()))
	if err != nil {
		fmt.Println("error creating form file: " + err.Error())
	}

	io.Copy(part, file)
	writer.Close()
	request, err := http.NewRequest(
		"POST",
		"http://"+ip.String()+"/upload",
		body,
	)
	if err != nil {
		fmt.Println("error creating http request: " + err.Error())
	}

	request.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		fmt.Println("error making upload request: " + err.Error())
	} else {
		resp, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(resp))
	}
	defer response.Body.Close()
}
