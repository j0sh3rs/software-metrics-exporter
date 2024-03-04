// Package main is the package entrance point
package main

import (
    "log"
    "github.com/j0sh3rs/software-metrics-exporter/pkg/collector"
	"github.com/j0sh3rs/software-metrics-exporter/pkg/writer"
)

func main() {
	// Collect installed software information
	softwareInfo, err := collector.CollectSoftwareInfo()
	if err != nil {
		log.Fatalf("Error collecting software information: %v", err)
	}

	// Write software information to a file in Prometheus/Open Metrics format
	err = writer.WriteMetrics(softwareInfo)
	if err != nil {
		log.Fatalf("Error writing metrics: %v", err)
	}
}