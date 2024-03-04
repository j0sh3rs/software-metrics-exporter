package main

import (
    "fmt"
    "log"
    "software-metrics-exporter/pkg/enumeration"
    "software-metrics-exporter/pkg/formatting"
    "software-metrics-exporter/pkg/output"
    "software-metrics-exporter/pkg/parsing"
)

func main() {
    // Enumerate installed packages
    packages, err := enumeration.EnumeratePackages()
    if err != nil {
        log.Fatalf("Error enumerating packages: %v", err)
    }

    // Parse enumerated packages
    parsedPackages, err := parsing.ParsePackages(packages)
    if err != nil {
        log.Fatalf("Error parsing packages: %v", err)
    }

    // Format parsed packages for Prometheus
    formattedMetrics, err := formatting.FormatForPrometheus(parsedPackages)
    if err != nil {
        log.Fatalf("Error formatting metrics: %v", err)
    }

    // Output metrics to file
    if err := output.WriteToFile(formattedMetrics, "installed_packages.metrics"); err != nil {
        log.Fatalf("Error writing metrics to file: %v", err)
    }

    fmt.Println("Metrics successfully written to installed_packages.metrics")
}
