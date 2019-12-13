package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
		l.Fatal("failed to read config:", err)
	}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		l.Fatal("failed to parse config:", err)
	}

	return c
}
