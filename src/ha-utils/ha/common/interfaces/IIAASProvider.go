package interfaces

import (
	"ha-utils/ha/common/models"
)

type IIAASProvider interface {
	Initialize(models.ConfigParams) int
	IsHAEnabled() bool
	ManageResources() int
	GetConfig() models.ConfigParams
}
