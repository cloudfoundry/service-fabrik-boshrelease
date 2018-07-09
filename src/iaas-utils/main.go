package main

import (
	"flag"
	IAASProviderFactory "iaas-utils/ha/common/iaasproviderfactory"
	"iaas-utils/ha/common/interfaces"
	"iaas-utils/ha/common/models"
	"log"
	"os"
	"time"
)

func main() {
	os.Exit(mainWithReturnCode())
}

func mainWithReturnCode() int {

	var provider interfaces.IIAASProvider
	var startTime time.Time
	var returnValue int

	startTime = time.Now()
	log.Println("Starting Post-start activities for keepalived at ", startTime)

	config := ReadConfigurationParameters()
	log.Println(config)
	provider = IAASProviderFactory.GetProvider(config)
	if provider == nil {
		log.Println("Could not find any provider for this landscape. HA will not be available for this deployment")
		// post-start should not fail as this could be due to configuration.
		return 0
	}
	if provider.IsHAEnabled() == false {
		// post-start should not fail as this could be due to configuration.
		log.Println("HA is not enabled for this IAAS provider.")
		return 0
	}
	returnValue = provider.Initialize(config)
	if returnValue != 0 {
		log.Println("Provider initialization failed with status", returnValue)
		return returnValue
	}
	log.Println("HA is enabled for this deployment. Associated provider is being invoked for creating/managing resources")
	// Lets call manage resources to create / update resources required for HA.
	returnValue = provider.ManageResources()
	// Identify the time taken for record keeping.
	if returnValue != 0 {
		log.Println("Provider resource management failed with return status:", returnValue)
	}
	log.Println("Post-start activities finished with status :", returnValue, "at ", time.Now(), "Total time taken for Post-start - ", time.Since(startTime))
	return returnValue

}

func ReadConfigurationParameters() models.ConfigParams {

	var config models.ConfigParams

	flag.StringVar(&config.AuthorizationBaseURL, "authbaseurl", "", "base url to be used for authorization")
	flag.StringVar(&config.ClientId, "clientid", "", "clientid of iaas provider")
	flag.StringVar(&config.ClientSecret, "clientsecret", "", "client secret related to clientid")
	flag.StringVar(&config.TenantId, "tenantid", "", "tenant id")
	flag.StringVar(&config.ClientEmailId, "clientemailid", "", "email id of the client")
	flag.StringVar(&config.PrivateKeyId, "privatekeyid", "", "id of the private key used")
	flag.StringVar(&config.PrivateKey, "privatekey", "", "private key")
	flag.StringVar(&config.Scopes, "scopes", "", "scopes to be used")
	flag.StringVar(&config.VirtualPrivateNetworkName, "network", "", "network name")
	flag.StringVar(&config.SubnetName, "subnetwork", "", "subnetwork name to which floating ip belongs to")

	flag.StringVar(&config.GCPBaseURL, "gcpbaseurl", "", "base url to be used for gcp api calls")
	flag.StringVar(&config.ProjectId, "projectid", "", "gcp project id")
	flag.IntVar(&config.SFBrokerPort, "sfbrokerport", 9293, "service-fabrik-broker port")
	flag.IntVar(&config.SFReportPort, "sfreportport", 9294, "service-fabrik-report port")
	flag.IntVar(&config.SFDephooksPort, "sfdephooksport", 9295, "service-fabrik deployment hooks port")
	flag.IntVar(&config.SFExternalPort, "sfexternalport", 9292, "service-fabrik-broker external port")
	flag.IntVar(&config.SFApiServerPort, "sfapiserverport", 9443, "service-fabrik-broker api-server port")
	flag.StringVar(&config.DeploymentGuid, "deploymentguid", "", "guid of the deployment")
	flag.StringVar(&config.FloatingIP, "floatingip", "", "floating ip used for this deployment")
	flag.StringVar(&config.CurrentInstanceIP, "instanceip", "", "ip address of the current instance")

	flag.IntVar(&config.ProbeIntervalInSeconds, "probeinterval", 5, "probe interval in seconds")
	flag.IntVar(&config.ProbeHealthCheckCount, "probehealthcheck", 2, "probe health check count")
	flag.StringVar(&config.ProbeProtocol, "probeprotocol", "Http", "protocol to be used by probe")
	flag.IntVar(&config.ProbePort, "probeport", 9293, "port to be used by probe")
	flag.StringVar(&config.ProbeRequestPath, "proberequestpath", "", "probe request path")
	flag.StringVar(&config.Landscape, "landscape", "", "current landscape")

	flag.Parse()
	return config

}
