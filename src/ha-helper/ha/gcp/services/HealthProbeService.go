package services

import (
	"encoding/json"
	"ha-helper/ha/common/models"
	"ha-helper/ha/common/constants"
	commoninterfaces "ha-helper/ha/common/interfaces"
	gcpmodels "ha-helper/ha/gcp/models"
	gcputils "ha-helper/ha/gcp/utils"
	"log"
	"strings"
)

type HealthProbeService struct {
	svc commoninterfaces.IServiceClient
}

func (hpService *HealthProbeService) Initialize(params ...interface{}) {
	hpService.svc = params[0].(commoninterfaces.IServiceClient)
}

func (hpService *HealthProbeService) GetHealthProbe(probeName string) (*gcpmodels.Probe, bool) {

	var healthCheckAPIUrl, responseStr, responseCode string
	var returnValue bool
	var healthProbe *gcpmodels.Probe = &gcpmodels.Probe{}
	var params models.IaaSDescriptors = hpService.svc.GetIaaSDescriptors()

	healthCheckAPIUrl = params.ManagementURL + "/compute/v1/projects/" + params.ProjectId + "/global/healthChecks/" + probeName
	/*
		The REST API URL format is as follows.
		GET https://www.googleapis.com/compute/v1/projects/{project}/global/healthChecks/{resourceId}
	*/
	responseStr, responseCode, returnValue = hpService.svc.InvokeAPI("GET", healthCheckAPIUrl, hpService.svc.GetCommonRequestHeaders(), nil)
	if returnValue == true {
		err := json.Unmarshal([]byte(responseStr), healthProbe)
		if err != nil {
			log.Println("Error occurred while unmarshalling health probe details ", err.Error())
			return nil, false
		}
		log.Println("Health check information with name :", probeName, "and link :", healthProbe.SelfLink, "retrieved successfully.")
		return healthProbe, true
	} else {
		if strings.Compare(responseCode, constants.HTTP_STATUS_NOT_FOUND) != 0 {
			// if GetHealthProbe() is called before the given health check is created, 404 status will be returned.
			// If the error is something other than 404, then something wrong might've happened. Lets log it explicitly for later analysis.
			log.Println("Error occurred during health check get call : HTTP Status: ", responseCode, " error: ", responseStr)
			return nil, false
		} else {
			// 404. It could be the case that an health check is yet to be created. Lets return true.
			return nil, true
		}
	}

}

func (hpService *HealthProbeService) CreateHealthProbe(probe gcpmodels.ProbeInput) bool {

	var createHealthcheckAPIUrl, responseStr, responseCode string
	var returnValue bool
	var currentOperation *gcpmodels.Operation = &gcpmodels.Operation{}
	var params models.IaaSDescriptors = hpService.svc.GetIaaSDescriptors()

	createHealthcheckAPIUrl = params.ManagementURL + "/compute/v1/projects/" + params.ProjectId + "/global/healthChecks"
	log.Println(probe)
	/*
		The REST API URL format is as follows.
		POST https://www.googleapis.com/compute/v1/projects/{project}/global/healthChecks
	*/
	responseStr, responseCode, returnValue = hpService.svc.InvokeAPI("POST", createHealthcheckAPIUrl, hpService.svc.GetCommonRequestHeaders(), probe)
	if returnValue == true {
		err := json.Unmarshal([]byte(responseStr), currentOperation)
		if err != nil {
			log.Println("Error occurred while unmarshalling operation details while creating health check", err.Error())
			return false
		}
		log.Println("health check creation :", probe.Name, " initiated successfully.")
		// let's check if the operation status is successful.
		return hpService.IsProvisioningSuccessful(*currentOperation)
	} else {
		log.Println("Error occurred during health check creation call : HTTP Status: ", responseCode, " error: ", responseStr)
		return false
	}

}

func (hpService *HealthProbeService) IsProvisioningSuccessful(operation gcpmodels.Operation) bool {

	return gcputils.IsResourceProvisioningSuccessful(operation, hpService.svc)

}
