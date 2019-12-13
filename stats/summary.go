package stats

import "time"

type Summary struct {
	KernelBootup time.Duration `json:"kernelBootup"`
	ServiceInit  time.Duration `json:"serviceInit"`
}

