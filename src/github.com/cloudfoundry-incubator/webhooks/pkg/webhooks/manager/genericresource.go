package main

import (
	"bytes"
	"encoding/json"

	"github.com/golang/glog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ContextOptions represents the contex information in GenericOptions
type ContextOptions struct {
	Platform         string `json:"platform"`
	OrganizationGUID string `json:"organization_guid"`
	SpaceGUID        string `json:"space_guid"`
}

// GenericOptions represents the option information in Spec
type GenericOptions struct {
	ServiceID string         `json:"service_id"`
	PlanID    string         `json:"plan_id"`
	Context   ContextOptions `json:"context"`
}

// GenericLastOperation represents the last option information in Status
type GenericLastOperation struct {
	Type  string `json:"type"`
	State string `json:"state"`
}

// GenericSpec represents the Spec in GenericResource
type GenericSpec struct {
	Options string `json:"options,omitempty"`
	options GenericOptions
}

// GenericStatus type represents the status in GenericResource
type GenericStatus struct {
	AppliedOptions   string `json:"appliedOptions"`
	State            string `json:"state,omitempty"`
	LastOperationRaw string `json:"lastOperation,omitempty"`
	lastOperation    GenericLastOperation
	appliedOptions   GenericOptions
}

// GenericResource type represents a generic resource
type GenericResource struct {
    Kind string `json:"kind"`
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

func getAppliedOptions(crd GenericResource) GenericOptions {
	var op GenericOptions
	opDecoder := json.NewDecoder(bytes.NewReader([]byte(crd.Status.AppliedOptions)))
	if err := opDecoder.Decode(&op); err != nil {
		glog.Errorf("Could not unmarshal raw object: %v", err)
	}
	return op
}
