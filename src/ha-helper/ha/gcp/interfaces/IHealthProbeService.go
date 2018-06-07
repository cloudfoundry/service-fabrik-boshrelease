package interfaces

import (
	// shoudld be changed to common/models after creating common bean for these classes.
	gcpmodels "ha-helper/ha/gcp/models"
)

type IHealthProbeService interface {
	Initialize(...interface{})
	GetHealthProbe(string) (*gcpmodels.Probe, bool)
	CreateHealthProbe(gcpmodels.ProbeInput) bool
	IsProvisioningSuccessful(gcpmodels.Operation) bool
}
