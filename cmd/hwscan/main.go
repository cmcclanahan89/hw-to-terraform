package main

import (
	"fmt"
	"hw-to-terraform/internal/collect"
	"hw-to-terraform/internal/render"
	"hw-to-terraform/pkg"
	"time"
)

func main() {

	osName, err := collect.GetOS()
	if err != nil {
		fmt.Println("Error Detecting OS:", err)
		return
	}

	info := pkg.InfoCollect{
		SchemaVersion: "1.0.0",
		CollectedAt:   time.Now(),
		Hostname:      "your-hostname-here", // replace with actual value
		OS:            osName,
		Arch:          "amd64",           // replace with actual value
		LogicalCores:  8,                 // replace with actual value
		PhysicalCores: 4,                 // replace with actual value
		Memory:        "16GB",            // replace with actual value
		Disks:         nil,               // or fill in actual disk info
		IPAddress:     "192.168.1.100",   // replace with actual value
		AdminUsers:    []string{"admin"}, // replace with actual value
	}

	render.AddLinetoJson(info, "output/datacollection.json")

}
