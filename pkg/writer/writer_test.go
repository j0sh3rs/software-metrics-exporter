package writer

import (
	"software-metrics-exporter/pkg/collector"
	"testing"
	"time"
)

func TestWriteMetrics(t *testing.T) {
	// Placeholder for setup and mock expectations

	// Example softwareInfos
	softwareInfos := []collector.SoftwareInfo{
		{
			Name:        "example",
			Version:     "1.0",
			InstallDate: time.Now(),
		},
	}

	// Call WriteMetrics
	err := WriteMetrics(softwareInfos)

	// Placeholder for assertions
	if err != nil {
		t.Errorf("WriteMetrics() error = %v", err)
	}
}

func TestSanitizeLabelValue(t *testing.T) {
	// Test cases for sanitizing label values
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Normal case", "value", "value"},
		{"Contains quote", `va"lue`, `va\"lue`},
		{"Contains backslash", `va\lue`, `va\\lue`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sanitizeLabelValue(tt.input); got != tt.expected {
				t.Errorf("sanitizeLabelValue() = %v, want %v", got, tt.expected)
			}
		})
	}
}