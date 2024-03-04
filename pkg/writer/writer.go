package writer

import (
	"fmt"
	"os"
	"github.com/j0sh3rs/software-metrics-exporter/pkg/collector"
	"time"
)

// WriteMetrics formats the software information into Prometheus/Open Metrics format and writes it to a file
func WriteMetrics(softwareInfos []collector.SoftwareInfo) error {
	file, err := os.Create("software_metrics.prom")
	if err != nil {
		return err
	}
	defer file.Close()

	for _, info := range softwareInfos {
		// We're using a gauge metric type with labels for the version and install date.
		// The value for each gauge is set to 1 to indicate the presence of the software.
		// InstallDate is not included in the output as its accurate retrieval is complex and varies by system.
		metric := fmt.Sprintf(`software_installed{name="%s", version="%s"} 1`, info.Name, sanitizeLabelValue(info.Version))
		_, err := file.WriteString(metric + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

// sanitizeLabelValue ensures that the label values are safe for Prometheus metrics, escaping quotes and backslashes.
func sanitizeLabelValue(value string) string {
	replacer := strings.NewReplacer(`\`, `\\`, `"`, `\"`)
	return replacer.Replace(value)
}
