package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

// Metric holds a single telemetry data point
type Metric struct {
	Timestamp       time.Time `json:"timestamp"`
	CPUUsagePercent float64   `json:"cpu_usage_percent"`
}

func main() {
	pollInterval := 2 * time.Second

	for {
		// Sample CPU usage over 1 second (more accurate than instant reads)
		usage, err := cpu.Percent(time.Second, false)
		if err != nil || len(usage) == 0 {
			log.Printf("Failed to get CPU usage: %v", err)
			continue
		}

		// Build metric with timestamp
		metric := Metric{
			Timestamp:       time.Now(),
			CPUUsagePercent: usage[0],
		}

		// Convert to JSON and print
		out, _ := json.Marshal(metric)
		fmt.Println(string(out))

		// Wait remaining poll interval (since we already delayed for 1s in cpu.Percent)
		time.Sleep(pollInterval - time.Second)
	}
}
