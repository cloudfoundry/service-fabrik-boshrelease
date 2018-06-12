package models

import ()

type LoadBalancingRule struct {
	Kind                string   `json:"kind"`
	ID                  string   `json:"id"`
	CreationTimestamp   string   `json:"creationTimestamp"`
	Name                string   `json:"name"`
	Description         string   `json:"description"`
	Region              string   `json:"region"`
	IPAddress           string   `json:"IPAddress"`
	IPProtocol          string   `json:"IPProtocol"`
	Ports               []string `json:"ports"`
	SelfLink            string   `json:"selfLink"`
	LoadBalancingScheme string   `json:"loadBalancingScheme"`
	Subnetwork          string   `json:"subnetwork"`
	Network             string   `json:"network"`
	BackendService      string   `json:"backendService"`
}

type CreateLBRuleInput struct {
	Name                string   `json:"name"`
	Description         string   `json:"description"`
	IPAddress           string   `json:"IPAddress"`
	IPProtocol          string   `json:"IPProtocol"`
	Ports               []string `json:"ports"`
	LoadBalancingScheme string   `json:"loadBalancingScheme"`
	Subnetwork          string   `json:"subnetwork"`
	Network             string   `json:"network"`
	BackendService      string   `json:"backendService"`
	//	Target 				string	 `json:"target"`
}
