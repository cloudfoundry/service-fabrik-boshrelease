package interfaces

import (
	// shoudld be changed to common/beans after creating common bean for these classes.
	gcpbeans "ha-helper/ha/gcp/beans"
)

type IVMService interface {
	Initialize(...interface{})
	GetVirtualMachine(string) (*gcpbeans.VirtualMachine, bool)
	GetVirtualMachineByIP(string) (*gcpbeans.VirtualMachine, bool)
}
