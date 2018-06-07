package beans

type VMGroup struct {
	Kind              string `json:"kind"`
	ID                string `json:"id"`
	CreationTimestamp string `json:"creationTimestamp"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	Network           string `json:"network"`
	Fingerprint       string `json:"fingerprint"`
	Zone              string `json:"zone"`
	SelfLink          string `json:"selfLink"`
	Size              int    `json:"size"`
	Subnetwork        string `json:"subnetwork"`
}
