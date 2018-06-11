package providers

import (
	"encoding/json"
	"fmt"
	"ha-utils/ha/common/models"
	commoninterfaces "ha-utils/ha/common/interfaces"
	gcpmodels "ha-utils/ha/gcp/models"
	"ha-utils/ha/gcp/clients"
	"ha-utils/ha/gcp/interfaces"
	"ha-utils/ha/gcp/services"
	"log"
	"strings"
)

type GCPIAAS struct {
	Config         models.ConfigParams
	serviceClient  commoninterfaces.IServiceClient
	vmService      interfaces.IVMService
	vmGroupService interfaces.IVMGroupService
	hpService      interfaces.IHealthProbeService
	lbService      interfaces.ILoadBalancerService
	lbRuleService  interfaces.ILoadBalancingRuleService
}

func (iaasProvider *GCPIAAS) Initialize(configParams models.ConfigParams) int {

	var iaasDescriptors models.IaaSDescriptors
	var authorizationRequest models.AuthorizationRequest
	var returnValue int
	//
	// Initialize relevant configurations
	iaasProvider.Config = configParams

	iaasDescriptors = models.IaaSDescriptors{
		ManagementURL: iaasProvider.Config.GCPBaseURL,
		ProjectId: iaasProvider.Config.ProjectId,
	}

	authorizationRequest = models.AuthorizationRequest{
		AuthBaseURL:   iaasProvider.Config.AuthorizationBaseURL,
		PrivateKeyId:  iaasProvider.Config.PrivateKeyId,
		PrivateKey:    iaasProvider.Config.PrivateKey,
		ClientEmailId: iaasProvider.Config.ClientEmailId,
		Scopes:        strings.Split(iaasProvider.Config.Scopes, ","),
	}

	// Initialize the service client so that it can be used by other services.
	iaasProvider.serviceClient = &clients.GCPServiceClient{}
	returnValue = iaasProvider.serviceClient.Initialize(iaasDescriptors, authorizationRequest)
	if returnValue != 0 {
		// Initialization has failed and cannot proceed any further
		return returnValue
	}

	// Initialize the services.
	iaasProvider.vmService = &services.VMService{}
	iaasProvider.vmService.Initialize(iaasProvider.serviceClient)

	iaasProvider.vmGroupService = &services.VMGroupService{}
	iaasProvider.vmGroupService.Initialize(iaasProvider.serviceClient)

	iaasProvider.hpService = &services.HealthProbeService{}
	iaasProvider.hpService.Initialize(iaasProvider.serviceClient)

	iaasProvider.lbService = &services.LoadBalancerService{}
	iaasProvider.lbService.Initialize(iaasProvider.serviceClient)

	iaasProvider.lbRuleService = &services.LoadBalancingRuleService{}
	iaasProvider.lbRuleService.Initialize(iaasProvider.serviceClient)

	return 0

}

func (iaasProvider *GCPIAAS) IsHAEnabled() bool {
	return true
}

func (iaasProvider *GCPIAAS) GetConfig() models.ConfigParams {
	return iaasProvider.Config
}

func (iaasProvider *GCPIAAS) initializeLandscapeInfo(virtualMachine *gcpmodels.VirtualMachine) (string, string, string, string, string, string) {

	var region, regionName, availabilityZone, azName, network, subNetwork string

	region = virtualMachine.Region
	temp := strings.Split(region, "/")
	regionName = temp[(len(temp) - 1)]
	availabilityZone = virtualMachine.Zone
	temp = strings.Split(availabilityZone, "/")
	azName = temp[(len(temp) - 1)]
	log.Println("Region name is ", regionName, ". and AZ name is ", azName)

	for _, currentNetworkInterface := range virtualMachine.NetworkInterfaces {
		if currentNetworkInterface.NetworkIP == iaasProvider.Config.CurrentInstanceIP {
			network = currentNetworkInterface.Network
			subNetwork = currentNetworkInterface.Subnetwork
			break
		}
	}
	return region, regionName, availabilityZone, azName, network, subNetwork

}

func (iaasProvider *GCPIAAS) getLoadBalancerName() string {
	return "lb-" + iaasProvider.Config.DeploymentGuid
}

func (iaasProvider *GCPIAAS) getVMGroupName() string {
	return iaasProvider.getLoadBalancerName() + "-" + strings.Replace(iaasProvider.Config.CurrentInstanceIP, ".", "-", -1)
}

func (iaasProvider *GCPIAAS) getHealhProbeName() string {
	return "health-check--" + iaasProvider.getLoadBalancerName()
}

func (iaasProvider *GCPIAAS) getLoadBalancingRuleName() string {
	return "lbrule-" + iaasProvider.getLoadBalancerName()
}

func (iaasProvider *GCPIAAS) ManageResources() int {

	var loadBalancerName, vmGroupName, healthProbeName, lbRuleName, regionName, availabilityZone, azName, network, subNetwork string
	var virtualMachine *gcpmodels.VirtualMachine
	var vmGroup *gcpmodels.VMGroup
	var probe *gcpmodels.Probe
	var loadBalancer *gcpmodels.LoadBalancer
	var lbRule *gcpmodels.LoadBalancingRule
	var returnValue bool

	//	MANAGE INTERNAL LOAD BALANCER
	loadBalancerName = iaasProvider.getLoadBalancerName()
	log.Println("load balancer name is ", loadBalancerName)

	// Lets identify the virtual machine and its related details.
	virtualMachine, returnValue = iaasProvider.vmService.GetVirtualMachineByIP(iaasProvider.Config.CurrentInstanceIP)
	if returnValue == false || virtualMachine == nil {
		log.Println("Failed to retrieve virtual machine details - ", virtualMachine.Name)
		return 1
	}

	_, regionName, availabilityZone, azName, network, subNetwork = iaasProvider.initializeLandscapeInfo(virtualMachine)

	// Identify whether a instance group exists for this instance.
	vmGroupName = iaasProvider.getVMGroupName()
	vmGroup, returnValue = iaasProvider.vmGroupService.GetVMGroup(vmGroupName, azName)
	if returnValue == false {
		log.Println("Failed to retrieve vm group details - ", vmGroupName)
		return 2
	} else if returnValue == true && vmGroup != nil {
		log.Println("VM group with name:", vmGroupName, "already exists in GCP.")
	}
	// If the vm group doesn't exist, then create one.
	if vmGroup == nil {
		log.Println("VM group with name ", vmGroupName, "does not exist in GCP. Initiating creation of VM group.")
		returnValue = iaasProvider.vmGroupService.CreateVMGroup(vmGroupName, availabilityZone, network, subNetwork)
		if returnValue == false {
			return 3
		}
		vmGroup, returnValue = iaasProvider.vmGroupService.GetVMGroup(vmGroupName, azName)
		if returnValue == false {
			return 5
		}
	}
	// Let's add this instance to vm group
	returnValue = iaasProvider.vmGroupService.AddVMToVMGroup(vmGroupName, azName, virtualMachine.SelfLink)
	if returnValue == false {
		log.Println("Failed to associate instance:", iaasProvider.Config.CurrentInstanceIP, "with vm group:", vmGroupName)
		return 6
	}

	// Identify whether a health check exists with the given name.
	healthProbeName = iaasProvider.getHealhProbeName()
	probe, returnValue = iaasProvider.hpService.GetHealthProbe(healthProbeName)
	if returnValue == false {
		log.Println("Failed to retrieve health check details - ", loadBalancerName)
		return 7
	} else if returnValue == true && probe != nil {
		log.Println("Health check with name:", healthProbeName, "already exists in GCP.")
	}
	// If the load balancer doesn't exist, then create one.
	if probe == nil {
		log.Println("Health check with name ", healthProbeName, "does not exist in GCP. Initiating creation of load balancer resource.")
		returnValue = iaasProvider.createHealthProbe(healthProbeName)
		if returnValue == false {
			return 8
		}
		probe, returnValue = iaasProvider.hpService.GetHealthProbe(healthProbeName)
		if returnValue == false {
			return 9
		}
	}

	// Identify whether a load balancer exists with the given name.
	loadBalancer, returnValue = iaasProvider.lbService.GetLoadBalancer(loadBalancerName, regionName)
	if returnValue == false {
		log.Println("Failed to retrieve load balancer details - ", loadBalancerName)
		return 10
	} else if returnValue == true && loadBalancer != nil {

		log.Println("Load balancer with name:", loadBalancerName, "already exists in GCP.")
		var modifyLBInput gcpmodels.CreateLBInput = gcpmodels.CreateLBInput{}
		var connDrain gcpmodels.ConnectionDrainInfo = gcpmodels.ConnectionDrainInfo{30}
		var isUpdateRequired bool = false
		var modifyBackend gcpmodels.Backend = gcpmodels.Backend{}

		modifyBackend.BalancingMode = "CONNECTION"
		modifyBackend.Group = vmGroup.SelfLink
		modifyLBInput.Name = loadBalancerName
		modifyLBInput.LoadBalancingScheme = "INTERNAL"
		modifyLBInput.Protocol = "TCP"
		modifyLBInput.HealthChecks = []string{probe.SelfLink}
		modifyLBInput.Fingerprint = loadBalancer.Fingerprint
		modifyLBInput.ConnectionDraining = connDrain

		if loadBalancer.Backends == nil || len(loadBalancer.Backends) <= 0 {
			isUpdateRequired = true
		} else {
			for _, currBackend := range loadBalancer.Backends {
				isUpdateRequired = true
				if strings.Compare(strings.TrimSpace(currBackend.Group), strings.TrimSpace(vmGroup.SelfLink)) == 0 {
					log.Println("This VM group is already present in load balancer. IG: ", currBackend.Group)
					isUpdateRequired = false
					break
				}
			}
		}
		if isUpdateRequired == true {
			modifyLBInput.Backends = append(loadBalancer.Backends, modifyBackend)
			returnValue = iaasProvider.lbService.UpdateLoadBalancer(modifyLBInput, regionName)
			if returnValue == false {
				return 11
			}
			log.Println("Load balancer with name ", loadBalancerName, "updated successfully.")
			loadBalancer, returnValue = iaasProvider.lbService.GetLoadBalancer(loadBalancerName, regionName)
			if returnValue == false {
				return 12
			}
		}
	}
	// If the load balancer doesn't exist, then create one.
	if loadBalancer == nil {
		log.Println("Load balancer with name ", loadBalancerName, "does not exist in GCP. Initiating creation of load balancer resource.")

		var createLBInput gcpmodels.CreateLBInput = gcpmodels.CreateLBInput{}
		var connDrain gcpmodels.ConnectionDrainInfo = gcpmodels.ConnectionDrainInfo{30}
		var backend gcpmodels.Backend = gcpmodels.Backend{}
		backend.BalancingMode = "CONNECTION"
		backend.Group = vmGroup.SelfLink
		
		createLBInput.Name = loadBalancerName
		createLBInput.LoadBalancingScheme = "INTERNAL"
		createLBInput.Protocol = "TCP"
		createLBInput.HealthChecks = []string{probe.SelfLink}
		createLBInput.Backends = []gcpmodels.Backend{backend}
		createLBInput.ConnectionDraining = connDrain
		returnValue = iaasProvider.lbService.CreateLoadBalancer(createLBInput, regionName)
		if returnValue == false {
			return 13
		}
		loadBalancer, returnValue = iaasProvider.lbService.GetLoadBalancer(loadBalancerName, regionName)
		if returnValue == false {
			return 14
		}
	}

	// Identify whether a load balancing rule exists with the given name.
	lbRuleName = iaasProvider.getLoadBalancingRuleName()
	lbRule, returnValue = iaasProvider.lbRuleService.GetLBRule(lbRuleName, regionName)
	if returnValue == false {
		log.Println("Failed to retrieve load balancing rule details - ", loadBalancerName)
		return 15
	} else if returnValue == true && lbRule != nil {
		log.Println("Load balancing rule with name:", lbRuleName, "already exists in GCP.")
	}
	// If the load balancing rule doesn't exist, then create one.
	if lbRule == nil {
		log.Println("Load balancing rule with name ", lbRuleName, "does not exist in GCP. Initiating creation of load balancing rule resource.")

		floatingIPNetwork, floatingIPSubNetwork, returnValue := iaasProvider.getNetworkInfo(iaasProvider.Config.SubnetName, regionName)
		var createLBRuleInput gcpmodels.CreateLBRuleInput = gcpmodels.CreateLBRuleInput{}
		createLBRuleInput.Name = lbRuleName
		createLBRuleInput.IPAddress = iaasProvider.Config.FloatingIP
		createLBRuleInput.IPProtocol = "TCP"
		createLBRuleInput.Ports = []string{fmt.Sprintf("%d", iaasProvider.Config.SFBrokerPort), fmt.Sprintf("%d", iaasProvider.Config.SFReportPort), fmt.Sprintf("%d", iaasProvider.Config.SFExternalPort), fmt.Sprintf("%d", iaasProvider.Config.SFDephooksPort)}
		createLBRuleInput.LoadBalancingScheme = "INTERNAL"
		createLBRuleInput.Network = floatingIPNetwork
		createLBRuleInput.Subnetwork = floatingIPSubNetwork
		createLBRuleInput.BackendService = loadBalancer.SelfLink

		returnValue = iaasProvider.lbRuleService.CreateLBRule(createLBRuleInput, regionName)
		if returnValue == false {
			return 16
		}
		loadBalancer, returnValue = iaasProvider.lbService.GetLoadBalancer(loadBalancerName, regionName)
		if returnValue == false {
			return 17
		}
	}

	// Everything went fine. Lets return 0.
	return 0

}

func (iaasProvider *GCPIAAS) getNetworkInfo(subNetworkName string, regionName string) (string, string, bool) {

	var subNetworkAPIUrl, responseStr string
	var subNetworkResult *gcpmodels.SubNetwork = &gcpmodels.SubNetwork{}
	var returnValue bool
	var params models.IaaSDescriptors = iaasProvider.serviceClient.GetIaaSDescriptors()

	subNetworkAPIUrl = params.ManagementURL + "/compute/v1/projects/" + params.ProjectId + "/regions/" + regionName + "/subnetworks/" + subNetworkName
	/*
		REST API format for getting the details of load balancing rule of a given load balancer is as follows.
		GET https://www.googleapis.com/compute/v1/projects/{project}/regions/{region}/subnetworks/{resourceId}
	*/
	responseStr, _, returnValue = iaasProvider.serviceClient.InvokeAPI("GET", subNetworkAPIUrl, iaasProvider.serviceClient.GetCommonRequestHeaders(), nil)
	if returnValue == true {
		err := json.Unmarshal([]byte(responseStr), subNetworkResult)
		if err != nil {
			log.Println("Error occurred while unmarshalling load balancing rule details ", err.Error())
			return "", "", false
		}
		log.Println("Subnetwork details with name : ", subNetworkResult.Name, "retrieved successfully.")
		return subNetworkResult.Network, subNetworkResult.SelfLink, true
	} else {
		return "", "", false
	}
}

func (iaasProvider *GCPIAAS) createHealthProbe(healthProbeName string) bool {

	var returnValue bool
	var inputProbe gcpmodels.ProbeInput

	inputProbe = gcpmodels.ProbeInput{
		Name:               healthProbeName,
		Type:               iaasProvider.Config.ProbeProtocol,
		HealthyThreshold:   iaasProvider.Config.ProbeHealthCheckCount,
		UnhealthyThreshold: iaasProvider.Config.ProbeHealthCheckCount,
		CheckIntervalSec:   iaasProvider.Config.ProbeIntervalInSeconds,
		TimeoutSec:         iaasProvider.Config.ProbeIntervalInSeconds,
		HTTPHealthCheck: gcpmodels.HTTPHC{
			Host: "",
			Port: iaasProvider.Config.ProbePort,
			RequestPath: iaasProvider.Config.ProbeRequestPath,
		},
	}
	returnValue = iaasProvider.hpService.CreateHealthProbe(inputProbe)
	if returnValue == false {
		log.Println("HealthProbeService: Failed to create health probe with name", healthProbeName)
		return false
	} else {
		log.Println("Health probe with name :", healthProbeName, "has been created successfully.")
		return true
	}

}
