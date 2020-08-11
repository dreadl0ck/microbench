package stats

import "time"

type Summary struct {
	KernelBootup   time.Duration `json:"kernelBootup"`
	KernelLogLines int           `json:"kernelLogLines"`
	KernelLogs     string        `json:"kernelLogs"`
}
