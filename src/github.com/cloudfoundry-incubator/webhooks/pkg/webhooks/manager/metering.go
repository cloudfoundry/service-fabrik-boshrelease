package main

type MeteringOptions struct {
	ServiceID  string `json:"service_id"`
	PlanID     string
	InstanceID string
	OrgID      string
	SpaceID    string
	Type       string
}

type MeteringSpec struct {
	Options MeteringOptions `json:"options"`
}

// Metering structure holds all the details related to
// Metering event
type Metering struct {
	Spec MeteringSpec `json:"spec"`
}
