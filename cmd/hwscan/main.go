package main

import (
	"encoding/json"
	"fmt"
	"hw-to-terraform/internal/collect"
	"hw-to-terraform/internal/normalize"
	"hw-to-terraform/internal/render"
	"hw-to-terraform/pkg"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
	"time"
)

func main() {

	outputPath := "output/datacollection.json"
	absPath, err := filepath.Abs(outputPath)
	if err != nil {
		log.Fatalf("Failed to resolve absolute path: %v", err)
	}

	mode := os.Args[1] // "terraform" or "bicep"
	var tmplFile, outputFile string

	switch mode {
	case "terraform":
		tmplFile = "template/terraform_vm_temp.tmpl"
		outputFile = "output/main.tf"
	case "bicep":
		tmplFile = "template/bicep_temp.tmpl"
		outputFile = "output/main.bicep"
	default:
		log.Fatalf("Unknown mode: %s", mode)
	}

	render.CreateJsonOutput()

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
		Hostname:      hostname,
		OS:            osName,
		Arch:          arch,
		LogicalCores:  logicalCores,
		PhysicalCores: physicalCores,
		Memory:        ramStr,
		Disks:         []pkg.DiskStats{diskStats},
		IPAddress:     hostIP,
	}

	render.AddLinetoJson(info, absPath)

	outputJson, err := os.Open(absPath)
	if err != nil {
		log.Fatal("Could not open file:", err)
	}
	defer outputJson.Close()

	var inv pkg.InfoCollect

	decoder := json.NewDecoder(outputJson)
	if err := decoder.Decode(&inv); err != nil {
		log.Fatal("Could not decode JSON:", err)
	}

	vmSize, err := normalize.DetermineVMSize(inv.LogicalCores, inv.Memory)
	if err != nil {
		fmt.Println("Error determining size:", err)
	}

	osPub, osOffer, osSku := normalize.ParseOS("absPath")
	if err != nil {
		fmt.Println("Error getting OS:", err)
	}

	hn := inv.Hostname

	templateData := pkg.VMTemplateData{
		Hostname:    hn,
		VMSize:      vmSize,
		DiskGB:      diskStats.TotalGB,
		OSPublisher: osPub,
		OSOffer:     osOffer,
		OSSku:       osSku,
		IPAddress:   hostIP,
	}

	tmpl, err := template.ParseFiles(tmplFile) //"template/terraform_vm_temp.tmpl")
	if err != nil {
		log.Fatal("Error parsing template:", err)
	}

	out, err := os.Create(outputFile) //"output/main.tf")
	if err != nil {
		log.Fatal("Error creating output file:", err)
	}
	fmt.Println("Wrote main.tf with data:", templateData)
	defer out.Close()

	if err := tmpl.Execute(out, templateData); err != nil {
		log.Fatal("Error executing template:", err)
	}

	log.Println("Terraform file generated at output/main.tf")
}
