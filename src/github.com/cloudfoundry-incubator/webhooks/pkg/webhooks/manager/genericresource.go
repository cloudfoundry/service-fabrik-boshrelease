package main

import (
	"bytes"
	"encoding/json"

	"github.com/golang/glog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ContextOptions struct {
	Platform         string `json:"platform"`
	OrganizationGuid string `json:"organization_guid"`
	SpaceGuid        string `json:"space_guid"`
}

type GenericOptions struct {
	ServiceId string         `json:"service_id"`
	PlanId    string         `json:"plan_id"`
	Context   ContextOptions `json:"context"`
}

type GenericLastOperation struct {
	Type  string `json:"type"`
	State string `json:"state"`
}

type GenericSpec struct {
	Options string `json:"options,omitempty"`
}

type GenericStatus struct {
	State            string `json:"state,omitempty"`
	LastOperationRaw string `json:"lastOperation,omitempty"`
	lastOperation    GenericLastOperation
}

type GenericResource struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Status            GenericStatus `json:"status,omitempty"`
	Spec              GenericSpec   `json:"spec,omitempty"`
}

func getGenericResource(object []byte) (GenericResource, error) {
	var crd GenericResource
	decoder := json.NewDecoder(bytes.NewReader(object))
	err := decoder.Decode(&crd)
	if err != nil {
		glog.Errorf("Could not unmarshal raw object: %v", err)
	}
	return crd, err
}

func getLastOperation(crd GenericResource) GenericLastOperation {
	var lo GenericLastOperation
	loDecoder := json.NewDecoder(bytes.NewReader([]byte(crd.Status.LastOperationRaw)))
	if err := loDecoder.Decode(&lo); err != nil {
		glog.Errorf("Could not unmarshal raw object: %v", err)
	}
	return lo
}

func getOptions(crd GenericResource) GenericOptions {
	var op GenericOptions
	opDecoder := json.NewDecoder(bytes.NewReader([]byte(crd.Spec.Options)))
	if err := opDecoder.Decode(&op); err != nil {
		glog.Errorf("Could not unmarshal raw object: %v", err)
	}
	return op
}
