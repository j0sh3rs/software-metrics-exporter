package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCmd struct {
	mock.Mock
}

func (m *MockCmd) Output() ([]byte, error) {
	args := m.Called()
	return args.Get(0).([]byte), args.Error(1)
}

func TestRunExporter(t *testing.T) {
	// Mock exec.Command
	originalExecCommand := execCommand
	defer func() { execCommand = originalExecCommand }()
	execCommand = func(name string, arg ...string) *exec.Cmd {
		mockCmd := new(MockCmd)
		mockCmd.On("Output").Return([]byte("package1 1.0.0 2023-04-01\npackage2 2.0.0 2023-04-02"), nil)
		return exec.Command("echo")
	}

	// Mock os.WriteFile
	originalWriteFile := writeFile
	defer func() { writeFile = originalWriteFile }()
	writeFile = func(filename string, data []byte, perm os.FileMode) error {
		return nil
	}

	var b bytes.Buffer
	viper.Set("output", "./metrics.txt")
	runExporter(nil, nil)
	assert.Contains(t, b.String(), "Metrics written to: ./metrics.txt")
}

func TestRunExporterFailOnCommand(t *testing.T) {
	originalExecCommand := execCommand
	defer func() { execCommand = originalExecCommand }()
	execCommand = func(name string, arg ...string) *exec.Cmd {
		mockCmd := new(MockCmd)
		mockCmd.On("Output").Return(nil, os.ErrNotExist)
		return exec.Command("echo")
	}

	var b bytes.Buffer
	runExporter(nil, nil)
	assert.Contains(t, b.String(), "Failed to execute command")
}

func TestRunExporterFailOnWriteFile(t *testing.T) {
	originalWriteFile := writeFile
	defer func() { writeFile = originalWriteFile }()
	writeFile = func(filename string, data []byte, perm os.FileMode) error {
		return os.ErrPermission
	}

	var b bytes.Buffer
	runExporter(nil, nil)
	assert.Contains(t, b.String(), "Failed to write metrics to file")
}

func TestMainFunction(t *testing.T) {
	// This test ensures that the main function can run without panicking
	assert.NotPanics(t, func() { main() }, "main function panicked")
}