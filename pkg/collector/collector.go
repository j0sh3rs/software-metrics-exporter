package collector

import (
	"bytes"
	"os/exec"
	"strings"
	"time"
)

// SoftwareInfo represents the details of an installed software package
type SoftwareInfo struct {
	Name         string
	Version      string
	InstallDate  time.Time
}

// CollectSoftwareInfo retrieves information about installed software packages
func CollectSoftwareInfo() ([]SoftwareInfo, error) {
	// Command to list installed packages with their name, version, and installation date
cmd := exec.Command("sh", "-c", "dpkg-query -W -f='${Package} ${Version}\n'")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	// Split the output into lines, each representing a package
	lines := strings.Split(out.String(), "\n")

	var softwareInfos []SoftwareInfo
	for _, line := range lines {
		if line == "" {
			continue
		}

		// Split each line into its components
		parts := strings.Fields(line)
		if len(parts) < 3 {
			continue
		}

		// For simplicity, we're currently skipping the installation date as it's not directly available from dpkg-query
		// and requires parsing logs or additional commands that might not be universally reliable.
		softwareInfos = append(softwareInfos, SoftwareInfo{
			Name:    parts[0],
			Version: parts[1],
			// Placeholder for InstallDate, consider implementing a method to determine this, if possible.
            InstallDate: installDate,
        // Attempt to get the installation date from the package's .list file modification time
        listFilePath := fmt.Sprintf("/var/lib/dpkg/info/%s.list", parts[0])
        fileInfo, err := os.Stat(listFilePath)
        var installDate time.Time
        if err == nil {
            installDate = fileInfo.ModTime()
        }
		})
	}

	return softwareInfos, nil
}
