package beans

import ()

type LoadBalancer struct {
	Kind                string              `json:"kind"`
	ID                  string              `json:"id"`
	CreationTimestamp   string              `json:"creationTimestamp"`
	Name                string              `json:"name"`
	Description         string              `json:"description"`
	SelfLink            string              `json:"selfLink"`
	Backends            []Backend           `json:"backends"`
	HealthChecks        []string            `json:"healthChecks"`
	TimeoutSec          int                 `json:"timeoutSec"`
	Protocol            string              `json:"protocol"`
	Fingerprint         string              `json:"fingerprint"`
	SessionAffinity     string              `json:"sessionAffinity"`
	Region              string              `json:"region"`
	LoadBalancingScheme string              `json:"loadBalancingScheme"`
	ConnectionDraining  ConnectionDrainInfo `json:"connectionDraining"`
}

type ConnectionDrainInfo struct {
	DrainingTimeoutSec int `json:"drainingTimeoutSec"`
}

type Backend struct {
	Description   string `json:"description"`
	Group         string `json:"group"`
	BalancingMode string `json:"balancingMode"`
}

type CreateLBInput struct {
	Backends            []Backend           `json:"backends"`
	HealthChecks        []string            `json:"healthChecks"`
	Name                string              `json:"name"`
	LoadBalancingScheme string              `json:"loadBalancingScheme"`
	Protocol            string              `json:"protocol"`
	Fingerprint         string              `json:"fingerprint"`
	ConnectionDraining  ConnectionDrainInfo `json:"connectionDraining"`
}
