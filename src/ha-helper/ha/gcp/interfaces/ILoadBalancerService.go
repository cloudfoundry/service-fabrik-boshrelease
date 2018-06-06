package interfaces

import (
	// shoudld be changed to common/beans after creating common bean for these classes.
	gcpbeans "ha-helper/ha/gcp/beans"
)

type ILoadBalancerService interface {
	Initialize(...interface{})
	GetLoadBalancer(string, string) (*gcpbeans.LoadBalancer, bool)
	CreateLoadBalancer(gcpbeans.CreateLBInput, string) bool
	UpdateLoadBalancer(gcpbeans.CreateLBInput, string) bool
	IsProvisioningSuccessful(gcpbeans.Operation) bool
}
