package interfaces

import (
	"iaas-utils/ha/common/models"
)

type IIAASProvider interface {
	Initialize(models.ConfigParams) int
	IsHAEnabled() bool
	ManageResources() int
	GetConfig() models.ConfigParams
}
