package beans

type AvailabilityZone struct {
	Kind                  string   `json:"kind"`
	ID                    string   `json:"id"`
	CreationTimestamp     string   `json:"creationTimestamp"`
	Name                  string   `json:"name"`
	Description           string   `json:"description"`
	Status                string   `json:"status"`
	Region                string   `json:"region"`
	SelfLink              string   `json:"selfLink"`
	AvailableCPUPlatforms []string `json:"availableCpuPlatforms,omitempty"`
}

type AvailabilityZoneList struct {
	Kind     string             `json:"kind"`
	ID       string             `json:"id"`
	Items    []AvailabilityZone `json:"items"`
	SelfLink string             `json:"selfLink"`
}
