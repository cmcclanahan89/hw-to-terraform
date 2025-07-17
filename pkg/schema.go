package pkg

import "time"

type InfoCollect struct {
	SchemaVersion string      `json:"schema_version"` // bump when the shape changes
	CollectedAt   time.Time   `json:"collected_at"`   // RFC 3339 for readability
	Hostname      string      `json:"hostname"`
	OS            string      `json:"os"`
	Arch          string      `json:"arch"`
	LogicalCores  int         `json:"logical_cores"`  // logical CPU cores
	PhysicalCores int         `json:"physical_cores"` // physical CPU cores
	Memory        string      `json:"RAM"`
	Disks         []DiskStats `json:"disks,omitempty"` // omit if empty
	IPAddress     string      `json:"IP Address,omitempty"`
}

type DiskStats struct {
	TotalGB float64 `json:"total_bytes"`
	Util    float64 `json:"used_percent"`
}
