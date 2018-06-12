package interfaces

import (
	// shoudld be changed to common/models after creating common bean for these classes.
	gcpmodels "iaas-utils/ha/gcp/models"
)

type IVMService interface {
	Initialize(...interface{})
	GetVirtualMachine(string) (*gcpmodels.VirtualMachine, bool)
	GetVirtualMachineByIP(string) (*gcpmodels.VirtualMachine, bool)
}
