package beans

import ()

type SubNetwork struct {
	Kind                  string `json:"kind"`
	ID                    string `json:"id"`
	CreationTimestamp     string `json:"creationTimestamp"`
	Name                  string `json:"name"`
	Network               string `json:"network"`
	IPCidrRange           string `json:"ipCidrRange"`
	GatewayAddress        string `json:"gatewayAddress"`
	Region                string `json:"region"`
	SelfLink              string `json:"selfLink"`
	PrivateIPGoogleAccess bool   `json:"privateIpGoogleAccess"`
	Fingerprint           string `json:"fingerprint"`
}
