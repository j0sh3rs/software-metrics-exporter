# Software Metrics Exporter

## Overview
Software Metrics Exporter is a tool designed to collect and export metrics about installed software packages on a system, specifically tailored for Prometheus monitoring. It leverages the `dpkg-query` command to list installed packages and their versions, and formats this information into the Prometheus/Open Metrics format.

## Features
- **Software Information Collection**: Collects details about installed software packages, including name, version, and installation date.
- **Prometheus-Compatible Output**: Formats the collected software metrics into a format that can be consumed by Prometheus.
- **Installation Date Collection**: Now also collects the installation date for each software package, providing a more comprehensive overview of the system's software inventory.
- **Custom Metric Sanitization**: Ensures that metric labels are safe for Prometheus by escaping quotes and backslashes.

## Getting Started

### Prerequisites
- Go 1.21.0 or higher
- Access to a system with `dpkg-query` available (typically Debian-based distributions)
- Ensure `awk` and `grep` utilities are available for processing command output

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/j0sh3rs/software-metrics-exporter.git
   ```
2. Navigate into the project directory:
   ```bash
   cd software-metrics-exporter
   ```
3. Build the project:
   ```bash
   go build -o software-metrics-exporter .
   ```

### Running the Exporter
To run the Software Metrics Exporter, execute the binary created during the installation process:
```bash
./software-metrics-exporter
```
This will collect the software metrics and write them to `installed_packages.metrics` file in the Prometheus/Open Metrics format.

## Contributing
Contributions are welcome! Please feel free to submit pull requests or create issues for bugs and feature requests.

## License
This project is licensed under the MIT License - see the LICENSE file for details.