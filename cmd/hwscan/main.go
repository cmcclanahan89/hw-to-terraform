package main

import (
	"fmt"
	"hw-to-terraform/internal/collect"
	"hw-to-terraform/internal/render"
	"hw-to-terraform/pkg"
	"strconv"
	"time"
)

func main() {

	osName, err := collect.GetOS()
	if err != nil {
		fmt.Println("Error Detecting OS:", err)
		return
	}

	arch, err := collect.GetArch()
	if err != nil {
		fmt.Println("Error Detecting arch:", err)
		return
	}

	intIP := collect.GetHostIP()
	hostIP := intIP.String()

	ramAmount, err := collect.GetRam()
	if err != nil {
		fmt.Println("Error Gathering Ram:", err)
		return
	}
	ramStr := strconv.FormatUint(ramAmount, 10) + " GiB"

	logicalCores, physicalCores, err := collect.GetCores()
	if err != nil {
		fmt.Println("Error Gathering Core Count:", err)
		return
	}

	hostname, err := collect.GetHostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		return
	}

	diskStats, err := collect.GetTotalDiskStats()
	if err != nil {
		fmt.Println("Error collecting disk stats:", err)
		return
	}

	info := pkg.InfoCollect{
		SchemaVersion: "1.0.0",
		CollectedAt:   time.Now(),
		Hostname:      hostname, // replace with actual value
		OS:            osName,
		Arch:          arch,                       // replace with actual value
		LogicalCores:  logicalCores,               // replace with actual value
		PhysicalCores: physicalCores,              // replace with actual value
		Memory:        ramStr,                     // replace with actual value
		Disks:         []pkg.DiskStats{diskStats}, // or fill in actual disk info
		IPAddress:     hostIP,                     // replace with actual value
	}

	render.AddLinetoJson(info, "output/datacollection.json")

}
