package cpu

import (
	"fmt"
	"strings"

	"github.com/shirou/gopsutil/v3/host"
)

func GetCPUTemp() (float64, error) {
	sensors, err := host.SensorsTemperatures()

	if err != nil {
		return 0, fmt.Errorf("Error: %w", err)
	}

	/* if your motherboards ACPI temperature sensor is different
	please replace it here */
	for _, sensor := range sensors {
		if strings.Contains(sensor.SensorKey, "acpitz") {
			return sensor.Temperature, nil
		}
	}

	return 0, fmt.Errorf("Not found")
}

func PrintCPUTemp() {
	temp, err := GetCPUTemp()

	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	fmt.Printf("CPU temperature: %.2f°C\n", temp)
}
