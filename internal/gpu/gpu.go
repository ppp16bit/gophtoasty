package gpu

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// if your gpu is not nvidia modify here or create another function
func NvidiaSmi() (float64, error) {
	cmd := exec.Command("nvidia-smi", "--query-gpu=temperature.gpu", "--format=csv,noheader")
	output, err := cmd.Output()

	if err != nil {
		return 0, fmt.Errorf("Error: %w", err)
	}

	tempStr := strings.TrimSpace(string(output))
	temp, err := strconv.ParseFloat(tempStr, 64)

	if err != nil {
		return 0, fmt.Errorf("error parsing GPU temperature: %w", err)
	}

	return temp, nil
}

func GetGPUTemp() (float64, error) {
	return NvidiaSmi()
}
