package models

type VirtualMachine struct {
	Kind              string `json:"kind"`
	ID                string `json:"id"`
	CreationTimestamp string `json:"creationTimestamp"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	Tags              struct {
		Items       []string `json:"items"`
		Fingerprint string   `json:"fingerprint"`
	} `json:"tags"`
	MachineType       string `json:"machineType"`
	Status            string `json:"status"`
	Zone              string `json:"zone"`
	Region            string // added explicitly to store the region information.
	NetworkInterfaces []struct {
		Kind        string `json:"kind"`
		Network     string `json:"network"`
		Subnetwork  string `json:"subnetwork"`
		NetworkIP   string `json:"networkIP"`
		Name        string `json:"name"`
		Fingerprint string `json:"fingerprint"`
	} `json:"networkInterfaces"`
	Disks []struct {
		Kind       string `json:"kind"`
		Type       string `json:"type"`
		Mode       string `json:"mode"`
		Source     string `json:"source"`
		DeviceName string `json:"deviceName"`
		Index      int    `json:"index"`
		Boot       bool   `json:"boot"`
		AutoDelete bool   `json:"autoDelete"`
		Interface  string `json:"interface"`
	} `json:"disks"`
	Metadata struct {
		Kind        string `json:"kind"`
		Fingerprint string `json:"fingerprint"`
		Items       []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"items"`
	} `json:"metadata"`
	SelfLink   string `json:"selfLink"`
	Scheduling struct {
		OnHostMaintenance string `json:"onHostMaintenance"`
		AutomaticRestart  bool   `json:"automaticRestart"`
		Preemptible       bool   `json:"preemptible"`
	} `json:"scheduling"`
	CPUPlatform string `json:"cpuPlatform"`
	Labels      struct {
		Deployment       string `json:"deployment"`
		Director         string `json:"director"`
		ID               string `json:"id"`
		Index            string `json:"index"`
		InstanceGroup    string `json:"instance_group"`
		Job              string `json:"job"`
		Name             string `json:"name"`
		OrganizationGUID string `json:"organization_guid"`
		Platform         string `json:"platform"`
		SpaceGUID        string `json:"space_guid"`
	} `json:"labels"`
	LabelFingerprint   string `json:"labelFingerprint"`
	StartRestricted    bool   `json:"startRestricted"`
	DeletionProtection bool   `json:"deletionProtection"`
}

type VirtualMachineList struct {
	Kind          string           `json:"kind"`
	ID            string           `json:"id"`
	Items         []VirtualMachine `json:"items"`
	SelfLink      string           `json:"selfLink"`
	NextPageToken string           `json:"nextPageToken"`
}
