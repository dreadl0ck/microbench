package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type machineConfig struct {
	IP      string `yaml:"ip"`
	Gateway string `yaml:"gw"`
}

type config struct {
	Vms []machineConfig `yaml:"vms"`
}

func parseConfig() *config {

	var c = new(config)

	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatal("failed to read config:", err)
	}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		log.Fatal("failed to parse config:", err)
	}

	return c
}
