package beans

import ()

type Probe struct {
	Kind               string `json:"kind"`
	ID                 string `json:"id"`
	CreationTimestamp  string `json:"creationTimestamp"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	CheckIntervalSec   int    `json:"checkIntervalSec"`
	TimeoutSec         int    `json:"timeoutSec"`
	UnhealthyThreshold int    `json:"unhealthyThreshold"`
	HealthyThreshold   int    `json:"healthyThreshold"`
	Type               string `json:"type"`
	HTTPHealthCheck    HTTPHC `json:"httpHealthCheck"`
	SelfLink           string `json:"selfLink"`
}

type HTTPHC struct {
	Port        int    `json:"port"`
	Host        string `json:"host"`
	RequestPath string `json:"requestPath"`
//	ProxyHeader string `json:"proxyHeader"`
}


type ProbeInput struct {
	Name               string `json:"name"`
	HealthyThreshold   int    `json:"healthyThreshold"`
	UnhealthyThreshold int    `json:"unhealthyThreshold"`
	CheckIntervalSec   int    `json:"checkIntervalSec"`
	TimeoutSec         int    `json:"timeoutSec"`
	Type               string `json:"type"`
	HTTPHealthCheck    HTTPHC `json:"httpHealthCheck"`
}