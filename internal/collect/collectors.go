package collect

import (
	"os/exec"
	"runtime"
	"strings"
)

func GetOS() string {
	switch runtime.GOOS {

	case "linux":
		cmd := exec.Command("uname", "-a")

		output, err := cmd.Output()
		if err != nil {
			return ""
		}
		return strings.TrimSpace(string(output))

	case "windows":
		cmd := exec.Command("powershell", "-NoProfile", "-Command", `Get-CimInstance Win32_Operatingsystem | Select-Object -expand Caption`)

		output, err := cmd.Output()
		if err != nil {
			return ""
		}
		return strings.TrimSpace(string(output))

	default:
		return runtime.GOOS
	}
}
