package collect

import (
	"fmt"
	"hw-to-terraform/pkg"
	"log"
	"math"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
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
	return localAddress.IP

}

func GetRam() (uint64, error) {

	ramAmount, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}
	return (ramAmount.Total / (1 << 30)), nil // Convert bytes to GiB
}

func GetCores() (int, int, error) {
	switch runtime.GOOS {
	case "linux":
		// Logical cores
		out, err := exec.Command("nproc").Output()
		if err != nil {
			return 0, 0, fmt.Errorf("error running nproc: %w", err)
		}
		logical, err := strconv.Atoi(strings.TrimSpace(string(out)))
		if err != nil {
			return 0, 0, fmt.Errorf("error parsing nproc output: %w", err)
		}

		// Physical cores (awk is simplest for one-liner)
		cmd := `awk '/^core id/ {print $4}' /proc/cpuinfo | sort -u | wc -l`
		out, err = exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			return logical, 0, fmt.Errorf("error running core id awk: %w", err)
		}
		physical, err := strconv.Atoi(strings.TrimSpace(string(out)))
		if err != nil {
			return logical, 0, fmt.Errorf("error parsing physical core count: %w", err)
		}

		return logical, physical, nil

	case "windows":
		logical, err := cpu.Counts(true)
		if err != nil {
			return 0, 0, err
		}
		physical, err := cpu.Counts(false)
		if err != nil {
			return 0, 0, err
		}
		return logical, physical, nil

	default:
		return 0, 0, fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}
}

func GetHostname() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", nil
	}
	return hostname, nil
}

func GetTotalDiskStats() (pkg.DiskStats, error) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return pkg.DiskStats{}, err
	}

	var total, used uint64
	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			continue // skip partitions we can't access
		}
		total += usage.Total
		used += usage.Used
	}

	var util float64
	if total > 0 {
		util = (float64(used) / float64(total)) * 100
		util = math.Round(util*100) / 100
	}

	totalGB := float64(total) / (1024 * 1024 * 1024)
	// Optional: Round to 2 decimals
	totalGB = math.Round(totalGB*100) / 100
	util = math.Round(util*100) / 100

	return pkg.DiskStats{
		TotalGB: totalGB,
		Util:    util,
	}, nil
}
