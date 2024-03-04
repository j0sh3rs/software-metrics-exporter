package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "software-metrics-exporter",
	Short: "Exports installed software metrics in Prometheus format",
	Run:   runExporter,
}

func init() {
	rootCmd.PersistentFlags().StringP("output", "o", "./metrics.txt", "Output file location for the Prometheus metrics")
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
	viper.SetDefault("output", "./metrics.txt")
}

func runExporter(cmd *cobra.Command, args []string) {
	outputFile := viper.GetString("output")

	cmdOut, err := exec.Command("bash", "-c", "dpkg-query -W -f='${Package} ${Version} ${Installed-Size} ${Status}\n' | grep 'install ok installed' | awk '{print $1 \" \" $2 \" \" $4 \"-\" $5 \"-\" $6}'").Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute command: %s\n", err)
		return
	}

	lines := strings.Split(string(cmdOut), "\n")
	var metrics strings.Builder
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			packageName := parts[0]
			packageVersion := parts[1]
			// Placeholder for installation date
			installDate := parts[2] // Assuming the date is now the third part of the output
			metrics.WriteString(fmt.Sprintf("package_info{name=\"%s\", version=\"%s\", install_date=\"%s\"} 1\n", packageName, packageVersion, installDate))
		}
	}

	err = os.WriteFile(outputFile, []byte(metrics.String()), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write metrics to file: %s\n", err)
		return
	}

	fmt.Println("Metrics written to:", outputFile)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
