package interfaces

import (
	// shoudld be changed to common/beans after creating common bean for these classes.
	gcpbeans "ha-helper/ha/gcp/beans"
)

type ILoadBalancingRuleService interface {
	Initialize(...interface{})
	GetLBRule(string, string) (*gcpbeans.LoadBalancingRule, bool)
	CreateLBRule(gcpbeans.CreateLBRuleInput, string) bool
	IsProvisioningSuccessful(gcpbeans.Operation) bool
}
