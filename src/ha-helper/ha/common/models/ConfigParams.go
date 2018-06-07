package models

import ()

type ConfigParams struct {

	// Authorization related configurations
	AuthorizationBaseURL      string
	ClientId                  string
	ClientSecret              string
	TenantId                  string
	ClientEmailId             string
	PrivateKeyId              string
	PrivateKey                string
	SubnetName                string
	VirtualPrivateNetworkName string
	// GCP IaaS related configurations
	GCPBaseURL string
	ProjectId  string
	//Zone      	string
	Scopes string
	Region string

	// 	deployment related imports.
	DeploymentGuid         string
	FloatingIP             string
	CurrentInstanceIP      string
	SFBrokerPort           int
	SFReportPort           int
	ProbeIntervalInSeconds int
	ProbeHealthCheckCount  int
	ProbeProtocol          string
	ProbePort              int
	ProbeRequestPath       string
	Landscape              string
	Action                 string
	ExecutionMode          string
}
