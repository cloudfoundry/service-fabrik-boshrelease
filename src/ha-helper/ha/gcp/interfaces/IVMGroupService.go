package interfaces

import (
	// shoudld be changed to common/beans after creating common bean for these classes.
	gcpbeans "ha-helper/ha/gcp/beans"
)

type IVMGroupService interface {
	Initialize(...interface{})
	GetVMGroup(string, string) (*gcpbeans.VMGroup, bool)
	CreateVMGroup(string, string, string, string) bool
	AddVMToVMGroup(string, string, string) bool
	IsProvisioningSuccessful(gcpbeans.Operation) bool
}
