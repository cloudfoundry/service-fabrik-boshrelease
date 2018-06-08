package interfaces

import (
	// shoudld be changed to common/models after creating common bean for these classes.
	gcpmodels "ha-utils/ha/gcp/models"
)

type IVMGroupService interface {
	Initialize(...interface{})
	GetVMGroup(string, string) (*gcpmodels.VMGroup, bool)
	CreateVMGroup(string, string, string, string) bool
	AddVMToVMGroup(string, string, string) bool
	IsProvisioningSuccessful(gcpmodels.Operation) bool
}
