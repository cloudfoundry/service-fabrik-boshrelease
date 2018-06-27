package utils

import (
	"encoding/json"
	commoninterfaces "iaas-utils/ha/common/interfaces"
	gcpmodels "iaas-utils/ha/gcp/models"
	"log"
	"time"
)

func IsResourceProvisioningSuccessful(operation gcpmodels.Operation, svc commoninterfaces.IServiceClient) bool {

	var startTime time.Time
	var returnValue bool
	var responseStr string
	var currentOperation *gcpmodels.Operation = &gcpmodels.Operation{}

	startTime = time.Now()
	for timeElapsed := 0.0; timeElapsed < svc.GetProvisioningWaitTime(); timeElapsed = time.Since(startTime).Seconds() {

		responseStr, _, returnValue = svc.InvokeAPI("GET", operation.SelfLink, svc.GetCommonRequestHeaders(), nil)
		if returnValue == true {
			err := json.Unmarshal([]byte(responseStr), currentOperation)
			if err != nil {
				log.Println("Error occurred while unmarshalling operation details ", err.Error())
				return false
			}
			log.Println("Operation details with name :", currentOperation.Name, " ID: ", currentOperation.ID,
				" and OpType:", currentOperation.OperationType, " retrieved successfully.")
		} else {
			log.Println("Failed to fetch operation details - ", operation.ID, ". Op status fetch will be retried after ",
				svc.GetProvisioningPollTime, "seconds")
			time.Sleep(time.Duration(svc.GetProvisioningPollTime()) * time.Second)
			continue
		}

		if currentOperation.Status == "DONE" {
			log.Println("Operation details with name: ", currentOperation.Name, " ID: ", currentOperation.ID, " Status: ", currentOperation.Status,
				" and OpType: ", currentOperation.OperationType, " has been completed successfully.")
			return true
		}
		log.Println("( CurrentOperation Status:", currentOperation.Status, "): Sleeping for some time before checking operation status again ", time.Now())
		// Sleep for a configured period of time, before identifying operation status.
		time.Sleep(time.Duration(svc.GetProvisioningPollTime()) * time.Second)
	}

	log.Println("Operation details with name: ", operation.Name, " ID: ", operation.ID, " Status: ", operation.Status,
		" and OpType: ", operation.OperationType, " failed.")
	return false

}
