package collect

import (
	"os/exec"
	"runtime"
	"strings"
)

func GetOS() (string, bool, error) {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("powershell", "-NoProfile", "-Command", `Get-CimInstance Win32_Operatingsystem | Select-Object -expand Caption`)

		output, err := cmd.Output()
		if err != nil {
			return "", true, err
		}
		return strings.TrimSpace(string(output)), true, nil

	default:
		cmd := exec.Command("uname,", "-a")

		output, err := cmd.Output()
		if err != nil {
			return "", false, err
		}
		return strings.TrimSpace(string(output)), false, nil
	}
}
