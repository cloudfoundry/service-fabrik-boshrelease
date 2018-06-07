package interfaces

import (
	// shoudld be changed to common/beans after creating common bean for these classes.
	gcpbeans "ha-helper/ha/gcp/beans"
)

type IHealthProbeService interface {
	Initialize(...interface{})
	GetHealthProbe(string) (*gcpbeans.Probe, bool)
	CreateHealthProbe(gcpbeans.ProbeInput) bool
	IsProvisioningSuccessful(gcpbeans.Operation) bool
}
