package clients

import (
	"iaas-utils/ha/common/models"
	"iaas-utils/ha/common/interfaces"
	"iaas-utils/ha/common/utils/apiutils"
	"iaas-utils/ha/common/constants"
	"iaas-utils/ha/gcp/services"
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

	svc.provisioningWaitTime = constants.PROVISIONING_WAIT_TIME
	svc.provisioningPollTime = constants.PROVISIONING_POLL_TIME

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
