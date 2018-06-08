package clients

import (
	"ha-utils/ha/common/models"
	"ha-utils/ha/common/interfaces"
	"ha-utils/ha/common/utils/apiutils"
	"ha-utils/ha/gcp/services"
	"log"
)

type GCPServiceClient struct {
	iaasDescriptors      models.IaaSDescriptors
	authorizationRequest models.AuthorizationRequest
	authorizationToken   *models.AuthorizationToken
	authorizationService interfaces.IAuthorizationService
	provisioningWaitTime float64
	provisioningPollTime float64
}

func (svc *GCPServiceClient) Initialize(params ...interface{}) int {

	svc.provisioningWaitTime = 300
	svc.provisioningPollTime = 2

	svc.iaasDescriptors = params[0].(models.IaaSDescriptors)
	svc.authorizationRequest = params[1].(models.AuthorizationRequest)
	// lets intialize auth service to be used for this client.
	svc.authorizationService = &services.AuthorizationService{}
	svc.authorizationService.Initialize()

	return 0

}

func (svc *GCPServiceClient) GetIaaSDescriptors() models.IaaSDescriptors {
	return svc.iaasDescriptors
}

func (svc *GCPServiceClient) InvokeAPI(httpMethod string, apiURL string, reqHeader map[string]string, requestBody interface{}) (string, string, bool) {
	return apiutils.InvokeRESTAPI(httpMethod, apiURL, reqHeader, requestBody)
}

func (svc *GCPServiceClient) GetCommonRequestHeaders() map[string]string {

	var reqHeader map[string]string

	reqHeader = make(map[string]string)
	reqHeader["Authorization"] = "Bearer " + svc.getAcessToken()
	reqHeader["Content-Type"] = "application/json"
	return reqHeader

}

func (svc *GCPServiceClient) getAcessToken() string {

	var returnValue bool

	if svc.authorizationToken == nil { // or auth token has expired
		svc.authorizationToken, returnValue = svc.authorizationService.Authorize(svc.authorizationRequest)
		if returnValue == false {
			log.Println("Could not authorize the request")
			svc.authorizationToken = nil
			return ""
		}
	}
	return svc.authorizationToken.AccessKey

}

func (svc *GCPServiceClient) GetProvisioningWaitTime() float64 {
	return svc.provisioningWaitTime
}

func (svc *GCPServiceClient) GetProvisioningPollTime() float64 {
	return svc.provisioningPollTime
}
