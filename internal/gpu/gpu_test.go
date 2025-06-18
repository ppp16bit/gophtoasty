package gpu

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestNvidiaSmi(t *testing.T) {
	_, err := exec.LookPath("nvidia-smi")
	if err != nil {
		t.Skip("skip: nvidia-smi not found")
	}

	temp, err := NvidiaSmi()
	if err != nil {
		t.Fatalf("nvidia-smi error: %v", err)
	}

	if temp < 0 || temp > 105 {
		t.Errorf("invalid gpu temperature: %.2f°C", temp)
	}
}

func TestPrintGPUTemp_Success(t *testing.T) {
	_, err := exec.LookPath("nvidia-smi")
	if err != nil {
		t.Skip("skip: nvidia-smi not found")
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintGPUTemp()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "GPU temperature:") {
		t.Errorf("missing 'GPU temperature:' in output: %q", output)
	}
	if !strings.Contains(output, "°C") {
		t.Errorf("missing '°C' in output: %q", output)
	}
}

func TestPrintGPUTemp_Error(t *testing.T) {
	originalPath := os.Getenv("PATH")
	os.Setenv("PATH", "/tmp/invalid-path")
	defer os.Setenv("PATH", originalPath)

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintGPUTemp()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if !strings.Contains(output, "Error:") {
		t.Errorf("expected error output, got: %q", output)
	}
}
