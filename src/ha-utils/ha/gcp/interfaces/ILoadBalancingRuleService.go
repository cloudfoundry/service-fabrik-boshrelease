package interfaces

import (
	// shoudld be changed to common/models after creating common bean for these classes.
	gcpmodels "ha-helper/ha/gcp/models"
)

type ILoadBalancingRuleService interface {
	Initialize(...interface{})
	GetLBRule(string, string) (*gcpmodels.LoadBalancingRule, bool)
	CreateLBRule(gcpmodels.CreateLBRuleInput, string) bool
	IsProvisioningSuccessful(gcpmodels.Operation) bool
}
