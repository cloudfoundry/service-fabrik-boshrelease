package interfaces

import (
	// shoudld be changed to common/models after creating common bean for these classes.
	gcpmodels "ha-helper/ha/gcp/models"
)

type ILoadBalancerService interface {
	Initialize(...interface{})
	GetLoadBalancer(string, string) (*gcpmodels.LoadBalancer, bool)
	CreateLoadBalancer(gcpmodels.CreateLBInput, string) bool
	UpdateLoadBalancer(gcpmodels.CreateLBInput, string) bool
	IsProvisioningSuccessful(gcpmodels.Operation) bool
}
