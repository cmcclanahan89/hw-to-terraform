package normalize

import "strings"

func DetermineVMSize(cores int, ram string) (string, error) {
	if cores >= 16 && strings.Contains(ram, "16") {
		return "Standard_D4s_v3", nil
	}
	return "Standard_B2s", nil
}

func ParseOS(os string) (publisher, offer, sku string) {
	osLower := strings.ToLower(os)
	switch {
	case strings.Contains(osLower, "ubuntu"):
		return "Canonical", "UbuntuServer", "18.04-LTS"
	case strings.Contains(osLower, "windows"):
		return "MicrosoftWindowsServer", "WindowsServer", "2019-Datacenter"
	case strings.Contains(osLower, "debian"):
		return "Debian", "debian-11", "11"
	default:
		return "Canonical", "UbuntuServer", "latest"
	}
}
