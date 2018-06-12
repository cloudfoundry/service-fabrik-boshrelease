package services

import (
	"encoding/json"
	"iaas-utils/ha/common/models"
	"iaas-utils/ha/common/constants"
	commoninterfaces "iaas-utils/ha/common/interfaces"
	gcpmodels "iaas-utils/ha/gcp/models"
	gcputils "iaas-utils/ha/gcp/utils"
	"log"
	"strings"
)

type LoadBalancerService struct {
	svc commoninterfaces.IServiceClient
}

func (lbService *LoadBalancerService) Initialize(params ...interface{}) {

	lbService.svc = params[0].(commoninterfaces.IServiceClient)

}

func (lbService *LoadBalancerService) GetLoadBalancer(lbName string, regionName string) (*gcpmodels.LoadBalancer, bool) {

	var loadBalancerAPIUrl, responseStr, responseCode string
	var returnValue bool
	var loadBalancer *gcpmodels.LoadBalancer = &gcpmodels.LoadBalancer{}
	var params models.IaaSDescriptors = lbService.svc.GetIaaSDescriptors()

	loadBalancerAPIUrl = params.ManagementURL + "/compute/v1/projects/" + params.ProjectId + "/regions/" + regionName + "/backendServices/" + lbName
	/*
		The REST API URL format is as follows.
		// https://www.googleapis.com/compute/v1/projects/sap-picpcore-gcp-dev8/regions/europe-west1/backendServices/postgres-ha-test-lb1
		GET https://www.googleapis.com/compute/v1/projects/{project}/global/backendServices/{resourceId}
	*/
	responseStr, responseCode, returnValue = lbService.svc.InvokeAPI("GET", loadBalancerAPIUrl, lbService.svc.GetCommonRequestHeaders(), nil)
	if returnValue == true {
		err := json.Unmarshal([]byte(responseStr), loadBalancer)
		if err != nil {
			log.Println("Error occurred while unmarshalling load balancer details ", err.Error())
			return nil, false
		}
		log.Println("Load Balancer with name :", lbName, "and link :", loadBalancer.SelfLink, "retrieved successfully.")
		return loadBalancer, true
	} else {
		if strings.Compare(responseCode, constants.HTTP_STATUS_NOT_FOUND) != 0 {
			// if GetLoadBalancer() is called before the given health check is created, 404 status will be returned.
			// If the error is something other than 404, then something wrong might've happened. Lets log it explicitly for later analysis.
			log.Println("Error occurred during load balancer get call : HTTP Status: ", responseCode, " error: ", responseStr)
			return nil, false
		} else {
			// 404. It could be the case that a load balancer is yet to be created. Lets return true.
			return nil, true
		}
	}

}

func (lbService *LoadBalancerService) CreateLoadBalancer(createLBInput gcpmodels.CreateLBInput, regionName string) bool {

	var createLBAPIUrl, responseStr, responseCode string
	var returnValue bool
	var currentOperation *gcpmodels.Operation = &gcpmodels.Operation{}
	var params models.IaaSDescriptors = lbService.svc.GetIaaSDescriptors()

	createLBAPIUrl = params.ManagementURL + "/compute/v1/projects/" + params.ProjectId + "/regions/" + regionName + "/backendServices"
	/*
		The REST API URL format is as follows.
		POST https://www.googleapis.com/compute/v1/projects/{project}/global/backendServices
	*/
	responseStr, responseCode, returnValue = lbService.svc.InvokeAPI("POST", createLBAPIUrl, lbService.svc.GetCommonRequestHeaders(), createLBInput)
	if returnValue == true {
		err := json.Unmarshal([]byte(responseStr), currentOperation)
		if err != nil {
			log.Println("Error occurred while unmarshalling operation details while creating load balancer", err.Error())
			return false
		}
		log.Println("Load balancer creation :", createLBInput.Name, " initiated successfully.")
		return lbService.IsProvisioningSuccessful(*currentOperation)
	} else {
		log.Println("Error occurred during load balancer creation call : HTTP Status: ", responseCode, " error: ", responseStr)
		return false
	}

}

func (lbService *LoadBalancerService) UpdateLoadBalancer(modifyLBInput gcpmodels.CreateLBInput, regionName string) bool {

	var modifyLBAPIUrl, responseStr, responseCode string
	var returnValue bool
	var currentOperation *gcpmodels.Operation = &gcpmodels.Operation{}
	var params models.IaaSDescriptors = lbService.svc.GetIaaSDescriptors()

	modifyLBAPIUrl = params.ManagementURL + "/compute/v1/projects/" + params.ProjectId + "/regions/" + regionName + "/backendServices/" + modifyLBInput.Name
	/*
		The REST API URL format is as follows.
		PUT https://www.googleapis.com/compute/v1/projects/{project}/global/backendServices/{resourceId}
	*/
	responseStr, responseCode, returnValue = lbService.svc.InvokeAPI("PUT", modifyLBAPIUrl, lbService.svc.GetCommonRequestHeaders(), modifyLBInput)
	if returnValue == true {
		err := json.Unmarshal([]byte(responseStr), currentOperation)
		if err != nil {
			log.Println("Error occurred while unmarshalling operation details while updating load balancer", err.Error())
			return false
		}
		log.Println("Load balancer update :", modifyLBInput.Name, " initiated successfully.")
		return lbService.IsProvisioningSuccessful(*currentOperation)
	} else {
		log.Println("Error occurred during load balancer update call : HTTP Status: ", responseCode, " error: ", responseStr)
		return false
	}

}

func (lbService *LoadBalancerService) IsProvisioningSuccessful(operation gcpmodels.Operation) bool {
	return gcputils.IsResourceProvisioningSuccessful(operation, lbService.svc)
}
