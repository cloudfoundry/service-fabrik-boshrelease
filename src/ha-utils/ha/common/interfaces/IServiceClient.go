package interfaces

import (
	"ha-helper/ha/common/models"
)

type IServiceClient interface {
	Initialize(...interface{}) int
	GetIaaSDescriptors() models.IaaSDescriptors
	InvokeAPI(httpMethod string, apiURL string, reqHeader map[string]string, requestBody interface{}) (string, string, bool)
	GetCommonRequestHeaders() map[string]string
	GetProvisioningWaitTime() float64
	GetProvisioningPollTime() float64
}
