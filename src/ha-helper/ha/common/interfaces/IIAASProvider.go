package interfaces

import (
	"ha-helper/ha/common/beans"
)

type IIAASProvider interface {
	Initialize(beans.ConfigParams) int
	IsHAEnabled() bool
	ManageResources() int
	GetConfig() beans.ConfigParams
}
