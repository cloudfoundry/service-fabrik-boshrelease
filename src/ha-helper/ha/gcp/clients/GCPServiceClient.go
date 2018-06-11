package clients

import (
	"ha-helper/ha/gcp/services"
	"ha-helper/ha/common/beans"
	"ha-helper/ha/common/interfaces"
	"ha-helper/ha/common/utils/apiutils"
	"log"
)

type GCPServiceClient struct {
	iaasDescriptors      beans.IaaSDescriptors
	authorizationRequest beans.AuthorizationRequest
	authorizationToken   *beans.AuthorizationToken
	authorizationService interfaces.IAuthorizationService
	provisioningWaitTime float64
	provisioningPollTime float64
}

func (svc *GCPServiceClient) Initialize(params ...interface{}) int {

	svc.provisioningWaitTime = 300
	svc.provisioningPollTime = 2

	svc.iaasDescriptors = params[0].(beans.IaaSDescriptors)
	svc.authorizationRequest = params[1].(beans.AuthorizationRequest)
	// lets intialize auth service to be used for this client.
	svc.authorizationService = &services.AuthorizationService{}
	svc.authorizationService.Initialize()

	return 0

}

func (svc *GCPServiceClient) GetIaaSDescriptors() beans.IaaSDescriptors {
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

