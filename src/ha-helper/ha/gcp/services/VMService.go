package services

import (
	"encoding/json"
	"ha-helper/ha/common/models"
	commoninterfaces "ha-helper/ha/common/interfaces"
	gcpmodels "ha-helper/ha/gcp/models"
	"log"
	"strings"
)

type VMService struct {
	svc commoninterfaces.IServiceClient
}

func (vmService *VMService) Initialize(params ...interface{}) {

	vmService.svc = params[0].(commoninterfaces.IServiceClient)

}

func (vmService *VMService) GetVirtualMachineByIP(instanceIP string) (*gcpmodels.VirtualMachine, bool) {

	var virtualMachineResult *gcpmodels.VirtualMachine = &gcpmodels.VirtualMachine{}
	var returnValue, vmInstanceFound bool

	azs, returnValue := vmService.GetAvailabilityZones()
	// No AZs found. Lets return.
	if returnValue == false || len(azs.Items) <= 0 {
		return nil, false
	}

	// Iterate through each AZ and find the VM.
	for _, currentAZ := range azs.Items {
		virtualMachineList, returnValue := vmService.GetVirtualMachines(currentAZ.Name)
		if returnValue == false {
			continue
		}
		for _, currentVirtualMachine := range virtualMachineList.Items {
			if len(currentVirtualMachine.NetworkInterfaces) <= 0 {
				continue
			}
			for _, currentNetworkInterface := range currentVirtualMachine.NetworkInterfaces {
				if currentNetworkInterface.NetworkIP == instanceIP {
					vmInstanceFound = true
					(*virtualMachineResult) = currentVirtualMachine
					virtualMachineResult.Region = currentAZ.Region
					break
				}
			}
			if vmInstanceFound == true {
				break
			}
		}
		if vmInstanceFound == true {
			break
		}
	}

	// If a VM corresponding to instanceIP cannot be found, lets return false.
	if !vmInstanceFound {
		log.Println("Error - Could not find any virtual machine with private address: ", instanceIP)
		return nil, false
	} else {
		log.Println("Found virtual machine :", virtualMachineResult.ID, "--", virtualMachineResult.Name, "corresponding to ip: ", instanceIP)
		//		log.Println(virtualMachineResult.Zone, virtualMachineResult.NetworkInterfaces[0].Network, virtualMachineResult.NetworkInterfaces[0].Subnetwork)
		return virtualMachineResult, true
	}

}

func (vmService *VMService) GetVirtualMachine(vmName string) (*gcpmodels.VirtualMachine, bool) {

	return nil, false
}

func (vmService *VMService) GetAvailabilityZones() (*gcpmodels.AvailabilityZoneList, bool) {

	var getAZsUrl, responseStr string
	var azList *gcpmodels.AvailabilityZoneList = &gcpmodels.AvailabilityZoneList{}
	var returnValue bool
	var iaasDescriptors models.IaaSDescriptors = vmService.svc.GetIaaSDescriptors()

	getAZsUrl = iaasDescriptors.ManagementURL + "/compute/v1/projects/" + iaasDescriptors.ProjectId + "/zones"
	/*
		The REST API URL for getting VM model view from Azure.

		" GET https://www.googleapis.com/compute/v1/projects/{project}/zones "

	*/
	responseStr, _, returnValue = vmService.svc.InvokeAPI("GET", getAZsUrl, vmService.svc.GetCommonRequestHeaders(), nil)
	if returnValue == true {
		err := json.Unmarshal([]byte(responseStr), azList)
		if err != nil {
			log.Println("Error occurred while unmarshalling az list details ", err.Error())
			return nil, false
		}
		log.Println("Availability zone list retrieved successfully.")
		return azList, true
	} else {
		return nil, false
	}

}

func (vmService *VMService) GetVirtualMachines(availabilityZone string) (*gcpmodels.VirtualMachineList, bool) {

	var getInstancesAPIUrl, responseStr, responseCode, currentPageToken string
	var virtualMachineList *gcpmodels.VirtualMachineList = &gcpmodels.VirtualMachineList{}
	var returnValue bool
	var iaasDescriptors models.IaaSDescriptors = vmService.svc.GetIaaSDescriptors()

	getInstancesAPIUrl = iaasDescriptors.ManagementURL + "/compute/v1/projects/" + iaasDescriptors.ProjectId + "/zones/" + availabilityZone + "/instances"

	/*
		The REST API URL for getting instances from GCP. At this point of time, we cannot pass filter with a given ip address.
		Thus we need to loop through the instances and get the vm resource id.

			" GET https://www.googleapis.com/compute/v1/projects/{project}/zones/{zone}/instances"

	*/
	currentPageToken = getInstancesAPIUrl
	for len(strings.TrimSpace(currentPageToken)) > 0 {
		var currvirtualMachineList gcpmodels.VirtualMachineList = gcpmodels.VirtualMachineList{}
		responseStr, responseCode, returnValue = vmService.svc.InvokeAPI("GET", getInstancesAPIUrl, vmService.svc.GetCommonRequestHeaders(), nil)
		if returnValue == true {
			err := json.Unmarshal([]byte(responseStr), &currvirtualMachineList)
			if err != nil {
				log.Println("Error occurred while unmarshalling virtual machine list details ", err.Error())
				return nil, false
			}
		} else {
			log.Println("Error occurred while fetching virtual machine list details : HTTP Status: ", responseCode, " error: ", responseStr)
			return nil, false
		}
		if (&currvirtualMachineList) == nil || len(currvirtualMachineList.Items) <= 0 {
			log.Println("Could not find any virtual machines in az ", availabilityZone)
			break
		} else {
			virtualMachineList.Items = append(virtualMachineList.Items, currvirtualMachineList.Items...)
			log.Println("[Debug]: List size is : ", len(virtualMachineList.Items))
			log.Println("[Debug]: Next link value is :|", currvirtualMachineList.NextPageToken, "|")
			currentPageToken = currvirtualMachineList.NextPageToken
		}
	}
	return virtualMachineList, true

}
