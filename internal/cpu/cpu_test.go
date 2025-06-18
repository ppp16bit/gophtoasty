package cpu

import (
	"fmt"
	"testing"
)

func TestGetCPUTemp(t *testing.T) {
	temp, err := GetCPUTemp()

	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if temp < 0 || temp > 100 {
		t.Errorf("unexpected CPU temperature: %.2f°C. Expected between 0 and 100°C.", temp)
	}

	fmt.Printf("Nice! - temp: %.2f°C\n", temp)
}
