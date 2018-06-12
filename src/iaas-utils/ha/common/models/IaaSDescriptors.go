package models

import ()

type IaaSDescriptors struct {
	ManagementURL             string
	SubscriptionId            string
	SubnetName                string
	VirtualPrivateNetworkName string
	ResourceGroupName         string
	APIVersion                string
	Location                  string
	//	Zone 					  string
	ProjectId string
}
