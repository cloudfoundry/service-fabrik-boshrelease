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

type VMGroupService struct {
	svc commoninterfaces.IServiceClient
}

func (vmGroupService *VMGroupService) Initialize(params ...interface{}) {

	vmGroupService.svc = params[0].(commoninterfaces.IServiceClient)

}

func (vmGroupService *VMGroupService) GetVMGroup(vmGroupName string, azName string) (*gcpmodels.VMGroup, bool) {

	var vmGroupAPIUrl, responseStr, responseCode string
	var returnValue bool
	var vmGroup *gcpmodels.VMGroup = &gcpmodels.VMGroup{}
	var params models.IaaSDescriptors = vmGroupService.svc.GetIaaSDescriptors()

	vmGroupAPIUrl = params.ManagementURL + "/compute/v1/projects/" + params.ProjectId + "/zones/" + azName + "/instanceGroups/" + vmGroupName
	/*
		    The REST API URL format is as follows.
			GET https://www.googleapis.com/compute/v1/projects/{project}/zones/{zone}/instanceGroups/{resourceId}
	*/
	responseStr, responseCode, returnValue = vmGroupService.svc.InvokeAPI("GET", vmGroupAPIUrl, vmGroupService.svc.GetCommonRequestHeaders(), nil)
	if returnValue == true {
		err := json.Unmarshal([]byte(responseStr), vmGroup)
		if err != nil {
			log.Println("Error occurred while unmarshalling vm group details ", err.Error())
			return nil, false
		}
		log.Println("VM Group information with name :", vmGroupName, "and id :", vmGroup.ID, "retrieved successfully.")
		return vmGroup, true
	} else {
		if strings.Compare(responseCode, constants.HTTP_STATUS_NOT_FOUND) != 0 {
			// if GetVMGroup() is called before the given vm group is created, 404 status will be returned. If the error is something other than 404,
			// then something wrong might've happened. Lets log it explicitly for later analysis.
			log.Println("Error occurred during instance group get call : HTTP Status: ", responseCode, " error: ", responseStr)
			return nil, false
		} else {
			// 404. It could be the case that an ILB is yet to be created. Lets return true.
			return nil, true
		}
	}

}

type CreateVMGroupInput struct {
	Name       string `json:"name"`
	Network    string `json:"network"`
	Subnetwork string `json:"subnetwork"`
}

func (vmGroupService *VMGroupService) CreateVMGroup(vmGroupName string, availabilityZone string, network string, subNetwork string) bool {

	var createVMGroupAPIUrl, responseStr, responseCode, azName string
	var returnValue bool
	var vmGroupInput CreateVMGroupInput = CreateVMGroupInput{}
	var currentOperation *gcpmodels.Operation = &gcpmodels.Operation{}
	var params models.IaaSDescriptors = vmGroupService.svc.GetIaaSDescriptors()

	temp := strings.Split(availabilityZone, "/")
	azName = temp[(len(temp) - 1)]

	createVMGroupAPIUrl = params.ManagementURL + "/compute/v1/projects/" + params.ProjectId + "/zones/" + azName + "/instanceGroups"
	vmGroupInput.Name = vmGroupName
	vmGroupInput.Network = network
	vmGroupInput.Subnetwork = subNetwork

	/*
		    The REST API URL format is as follows.
			POST https://www.googleapis.com/compute/v1/projects/{project}/zones/{zone}/instanceGroups
	*/
	responseStr, responseCode, returnValue = vmGroupService.svc.InvokeAPI("POST", createVMGroupAPIUrl, vmGroupService.svc.GetCommonRequestHeaders(), vmGroupInput)
	if returnValue == true {
		err := json.Unmarshal([]byte(responseStr), currentOperation)
		if err != nil {
			log.Println("Error occurred while unmarshalling operation details while creating instance group", err.Error())
			return false
		}
		log.Println("Instance Group creation :", vmGroupName, " initiated successfully.")
		// let's check if the operation status is successful.
		return vmGroupService.IsProvisioningSuccessful(*currentOperation)
	} else {
		log.Println("Error occurred during instance group creation call : HTTP Status: ", responseCode, " error: ", responseStr)
		return false
	}

}

type VMList struct {
	Kind  string `json:"kind"`
	Items []VM   `json:"items"`
}

type VMListInput struct {
	Instances []VM `json:"instances"`
}

type VM struct {
	Instance string `json:"instance"`
	Status   string `json:"status"`
}

func (vmGroupService *VMGroupService) AddVMToVMGroup(vmGroupName string, azName string, VMURL string) bool {

	var vmList *VMList
	var vmListInput VMListInput
	var currentOperation *gcpmodels.Operation = &gcpmodels.Operation{}
	var returnValue, isAssociationRequired bool
	var addVMsToVMGroupAPIUrl, responseStr, responseCode string
	var params models.IaaSDescriptors = vmGroupService.svc.GetIaaSDescriptors()

	vmList, returnValue = vmGroupService.GetVMsInVMGroup(vmGroupName, azName)
	if returnValue == false {
		return false
	}
	isAssociationRequired = false
	if vmList == nil || len(vmList.Items) <= 0 {
		log.Println("No VMs found in this VM group. Current VM will be associated with the VM group")
		isAssociationRequired = true
	} else if len(vmList.Items) > 0 {
		// Lets identify if the VM is already made part of this vm group.
		isAssociationRequired = true
		for _, currentVM := range vmList.Items {
			if currentVM.Instance == VMURL {
				isAssociationRequired = false
				log.Println("VM: ", VMURL, " is already associated with the VM group : ", vmGroupName)
			}
		}
	}
	// vm is already a part of the vm group - no action needs to be performed.
	if isAssociationRequired == false {
		return true
	}

	addVMsToVMGroupAPIUrl = params.ManagementURL + "/compute/v1/projects/" + params.ProjectId + "/zones/" + azName + "/instanceGroups/" +
		vmGroupName + "/addInstances"
	vmListInput = VMListInput{}
	vmListInput.Instances = []VM{
		{
			Instance: VMURL,
		},
	}
	/*
		The REST API URL format is as follows.
		POST https://www.googleapis.com/compute/v1/projects/{project}/zones/{zone}/instanceGroups/{resourceId}/addInstances
	*/
	responseStr, responseCode, returnValue = vmGroupService.svc.InvokeAPI("POST", addVMsToVMGroupAPIUrl, vmGroupService.svc.GetCommonRequestHeaders(), vmListInput)
	if returnValue == true {
		err := json.Unmarshal([]byte(responseStr), currentOperation)
		if err != nil {
			log.Println("Error occurred while unmarshalling operation details while adding VM to VM group", err.Error())
			return false
		}
		log.Println("Adding VM : ", VMURL, " to VM Group :", vmGroupName, " initiated successfully.")
		// let's check if the operation status is successful.
		return vmGroupService.IsProvisioningSuccessful(*currentOperation)
	} else {
		log.Println("Error occurred while addVMsToVMGroup API call : HTTP Status: ", responseCode, " error: ", responseStr)
		return false
	}
	return true

}

func (vmGroupService *VMGroupService) GetVMsInVMGroup(vmGroupName string, azName string) (*VMList, bool) {

	var listVMsinVMGroupAPIUrl, responseStr, responseCode string
	var returnValue bool
	var vmList *VMList = &VMList{}
	var params models.IaaSDescriptors = vmGroupService.svc.GetIaaSDescriptors()

	listVMsinVMGroupAPIUrl = params.ManagementURL + "/compute/v1/projects/" + params.ProjectId + "/zones/" + azName + "/instanceGroups/" +
		vmGroupName + "/listInstances"
	/*
		    The REST API URL format is as follows.
			POST https://www.googleapis.com/compute/v1/projects/{project}/zones/{zone}/instanceGroups/{resourceId}/listInstances
	*/
	responseStr, responseCode, returnValue = vmGroupService.svc.InvokeAPI("POST", listVMsinVMGroupAPIUrl, vmGroupService.svc.GetCommonRequestHeaders(), nil)
	if returnValue == true {
		err := json.Unmarshal([]byte(responseStr), vmList)
		if err != nil {
			log.Println("Error occurred while unmarshalling vm list details ", err.Error())
			return nil, false
		}
		log.Println("VMs in vm group : ", vmGroupName, " retrieved successfully.")
		return vmList, true
	} else {
		if strings.Compare(responseCode, constants.HTTP_STATUS_NOT_FOUND) != 0 {
			// if GetVMsInVMGroup() is called before the given vm is added, 404 status may be returned. If the error is something other than 404,
			// then something wrong might've happened. Lets log it explicitly for later analysis.
			log.Println("Error occurred during vms in vm-group get call : HTTP Status: ", responseCode, " error: ", responseStr)
			return nil, false
		} else {
			// 404. It could be the case that an ILB is yet to be created. Lets return true.
			return nil, true
		}
	}
	return nil, false
}

func (vmGroupService *VMGroupService) IsProvisioningSuccessful(operation gcpmodels.Operation) bool {

	return gcputils.IsResourceProvisioningSuccessful(operation, vmGroupService.svc)

}
