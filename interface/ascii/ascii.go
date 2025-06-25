package ascii

import (
	"fmt"

	"github.com/ppp16bit/goph-toasty/internal/cpu"
	"github.com/ppp16bit/goph-toasty/internal/gpu"
)

const (
	Reset            = "\x1b[0m"
	BlueFadeDark     = "\x1b[38;5;18m"
	BlueFadeMidDark  = "\x1b[38;5;20m"
	BlueFadeMid      = "\x1b[38;5;26m"
	BlueFadeMidLight = "\x1b[38;5;27m"
	BlueFadeLight    = "\x1b[38;5;33m"

	InterfaceColor = ""
)

func GetCPUAndGPUTemp() (cpuTemp string, gpuTemp string) {
	cpuVal, err := cpu.GetCPUTemp()
	if err != nil {
		cpuTemp = "N/A"
	} else {
		cpuTemp = fmt.Sprintf("%.2f°C", cpuVal)
	}

	gpuVal, err := gpu.GetGPUTemp()
	if err != nil {
		gpuTemp = "N/A"
	} else {
		gpuTemp = fmt.Sprintf("%.2f°C", gpuVal)
	}

	return cpuTemp, gpuTemp
}

func PrintTextAscii() {
	cpuTemp, gpuTemp := GetCPUAndGPUTemp()

	LBar1 := `╔══════════════════════════════════════════════════════════════════════════════════════════════════════════════╗`
	L1 := `║  ________  ________  ________  ___  ___  _________  ________  ________  ________  _________    ___    ___    ║`
	L2 := `║  |\   ____\|\   __  \|\   __  \|\  \|\  \|\___   ___\\   __  \|\   __  \|\   ____\|\___   ___\ |\  \  /  /|  ║`
	L3 := `║  \ \  \___|\ \  \|\  \ \  \|\  \ \  \\\  \|___ \  \_\ \  \|\  \ \  \|\  \ \  \___|\|___ \  \_| \ \  \/  / /  ║`
	L4 := `║   \ \  \  __\ \  \\\  \ \   ____\ \   __  \   \ \  \ \ \  \\\  \ \   __  \ \_____  \   \ \  \   \ \    / /   ║`
	L5 := `║    \ \  \|\  \ \  \\\  \ \  \___|\ \  \ \  \   \ \  \ \ \  \\\  \ \  \ \  \|____|\  \   \ \  \   \/  /  /    ║`
	L6 := `║     \ \_______\ \_______\ \__\    \ \__\ \__\   \ \__\ \ \_______\ \__\ \__\____\_\  \   \ \__\__/  / /      ║`
	L7 := `║      \|_______|\|_______|\|__|     \|__|\|__|    \|__|  \|_______|\|__|\|__|\_________\   \|__|\___/ /       ║`
	L8 := `║                                                                            \|_________|       \|___|/        ║`
	L9 := `║                                                                                                              ║`
	LTempCPU := fmt.Sprintf("║                                                 CPU: %-20s                                    ║", cpuTemp)
	LTempGPU := fmt.Sprintf("║                                                 GPU: %-20s                                    ║", gpuTemp)
	LBar3 := `╚══════════════════════════════════════════════════════════════════════════════════════════════════════════════╝`

	fmt.Println(BlueFadeDark + LBar1 + Reset)
	fmt.Println(BlueFadeDark + L1 + Reset)
	fmt.Println(BlueFadeMidDark + L2 + Reset)
	fmt.Println(BlueFadeMidDark + L3 + Reset)
	fmt.Println(BlueFadeMid + L4 + Reset)
	fmt.Println(BlueFadeMid + L5 + Reset)
	fmt.Println(BlueFadeMidLight + L6 + Reset)
	fmt.Println(BlueFadeLight + L7 + Reset)
	fmt.Println(BlueFadeLight + L8 + Reset)
	fmt.Println(BlueFadeLight + L9 + Reset)
	fmt.Println(BlueFadeLight + LTempCPU + Reset)
	fmt.Println(BlueFadeLight + LTempGPU + Reset)
	fmt.Println(BlueFadeLight + LBar3 + Reset)
}
