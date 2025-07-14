package collect

import (
	"log"
	"net"
	"os/exec"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/v4/mem"
)

func GetOS() (string, error) {
	switch runtime.GOOS {

	case "linux":
		cmd := exec.Command("uname", "-a")

		output, err := cmd.Output()
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(string(output)), nil

	case "windows":
		cmd := exec.Command("powershell", "-NoProfile", "-Command", `Get-CimInstance Win32_Operatingsystem | Select-Object -expand Caption`)

		output, err := cmd.Output()
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(string(output)), nil

	default:
		return runtime.GOOS, nil
	}
}

func GetArch() (string, error) {
	switch runtime.GOARCH {

	case "linux":
		cmd := exec.Command("uname", "-m")
		output, err := cmd.Output()
		if err != nil {
			return "", err
		}
		return string(output), nil

	case "windows":
		cmd := exec.Command("powershell", "-NoProfile", "-Command", `Get-CimInstance Win32_Processor | Select-Object -expand Caption`)

		output, err := cmd.Output()
		if err != nil {
			return "", err
		}
		r := string(output)
		result := strings.Fields(r)
		if len(result) > 0 {
			return result[0], err
		} else {
			return string(output), nil
		}
	default:
		return runtime.GOARCH, nil
	}
}

func GetHostIP() net.IP {

	hostIP, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer hostIP.Close()

	localAddress := hostIP.LocalAddr().(*net.UDPAddr)
	// fmt.Println("IP Address:", localAddress.IP)
	return localAddress.IP

}

func GetRam() (uint64, error) {

	ramAmount, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}
	return (ramAmount.Total / (1 << 30)), nil // Convert bytes to GiB
}
