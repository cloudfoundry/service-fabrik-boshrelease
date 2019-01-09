package main

// MeteringOptions represents the options field of Metering Resource
type MeteringOptions struct {
	ServiceID  string
	PlanID     string
	InstanceID string
	OrgID      string
	SpaceID    string
	Type       string
	Signal     string
}

// MeteringSpec represents the spec field of metering resource
type MeteringSpec struct {
	Options MeteringOptions `json:"options"`
}

// Metering structure holds all the details related to
// Metering event
type Metering struct {
	Spec MeteringSpec `json:"spec"`
}
