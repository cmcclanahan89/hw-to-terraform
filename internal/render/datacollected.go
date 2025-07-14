package render

import (
	"encoding/json"

	"fmt"
	"hw-to-terraform/pkg"
	"os"
)

func CreateJsonOutput() {
	file, ferr := os.Create("output/datacollection.json")
	if ferr != nil {
		fmt.Println("Failed to create file:", ferr)
		return
	}
	defer file.Close()

}

func AddLinetoJson(info pkg.InfoCollect, path string) {
	jsonFile, jerr := os.Create("output/datacollection.json")
	if jerr != nil {
		fmt.Println("Failed to open file:", jerr)
		return
	}
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "  ") // Pretty-print
	if jerr := encoder.Encode(info); jerr != nil {
		fmt.Println("Failed to encode JSON:", jerr)
		return
	}

}
