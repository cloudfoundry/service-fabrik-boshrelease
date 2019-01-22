package main

import (
	"context"
	"encoding/json"

	"k8s.io/client-go/rest"

	"github.com/golang/glog"
	"k8s.io/api/admission/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// Event stores the event details
type Event struct {
	AdmissionReview *v1beta1.AdmissionReview
	crd             GenericResource
	oldCrd          GenericResource
}

// NewEvent is a constructor for Event
func NewEvent(ar *v1beta1.AdmissionReview) (*Event, error) {
	crd, err := getGenericResource(ar.Request.Object.Raw)
	if err != nil {
		glog.Errorf("Could not get the GenericResource object %v", err)
		return nil, err
	}
	oldCrd, err := getGenericResource(ar.Request.OldObject.Raw)
	if err != nil {
		glog.Errorf("Could not get the old GenericResource object %v", err)
		return nil, err
	}
	crd.Status.lastOperation = getLastOperation(crd)
	crd.Spec.options = getOptions(crd)
	oldCrd.Status.lastOperation = getLastOperation(oldCrd)
	oldCrd.Spec.options = getOptions(oldCrd)
	crd.Status.appliedOptions = getAppliedOptions(crd)
	oldCrd.Status.appliedOptions = getAppliedOptions(oldCrd)
	return &Event{
		AdmissionReview: ar,
		crd:             crd,
		oldCrd:          oldCrd,
	}, nil
}

func (e *Event) isStateChanged() bool {
	loNew := e.crd.Status.lastOperation
	loOld := e.oldCrd.Status.lastOperation
	glog.Infof("New: type: %s, state: %s\n", loNew.Type, loNew.State)
	glog.Infof("Old: type: %s, state: %s\n", loOld.Type, loOld.State)
	return loNew.Type == loOld.Type && loNew.State != loOld.State
}

func (e *Event) isPlanChanged() bool {
	appliedOptionsNew := e.crd.Status.appliedOptions
	appliedOptionsOld := e.oldCrd.Status.appliedOptions
	return appliedOptionsNew.PlanID != appliedOptionsOld.PlanID
}

func (e *Event) isCreate() bool {
	return e.crd.Status.lastOperation.Type == "create"
}

func (e *Event) isUpdate() bool {
	return e.crd.Status.lastOperation.Type == "update"
}

func (e *Event) isSucceeded() bool {
	return e.crd.Status.lastOperation.State == "succeeded"
}

func (e *Event) isMeteringEvent() bool {
	if e.isStateChanged() && e.isSucceeded() {
		if e.isUpdate() && e.isPlanChanged() {
			return true
		}
		if e.isCreate() {
			return true
		}
	}
	return false
}

// ObjectToMapInterface converts an Object to map[string]interface{}
func ObjectToMapInterface(obj interface{}) (map[string]interface{}, error) {
	values := make(map[string]interface{})
	options, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(options, &values)
	if err != nil {
		return nil, err
	}
	return values, nil
}

func getClient(cfg *rest.Config) (client.Client, error) {
	glog.Infof("setting up manager")
	mgr, err := manager.New(cfg, manager.Options{})
	if err != nil {
		glog.Errorf("unable to set up overall controller manager %v", err)
		return nil, err
	}
	options := client.Options{
		Scheme: mgr.GetScheme(),
		Mapper: mgr.GetRESTMapper(),
	}
	apiserver, err := client.New(cfg, options)
	if err != nil {
		glog.Errorf("unable create kubernetes client %v", err)
		return nil, err
	}
	return apiserver, err
}

func meteringToUnstructured(m *Metering) (*unstructured.Unstructured, error) {
	values, err := ObjectToMapInterface(m)
	if err != nil {
		glog.Errorf("unable convert to map interface %v", err)
		return nil, err
	}
	meteringDoc := &unstructured.Unstructured{}
	meteringDoc.SetUnstructuredContent(values)
	meteringDoc.SetKind("Sfevent")
	meteringDoc.SetAPIVersion("instance.servicefabrik.io/v1alpha1")
	meteringDoc.SetNamespace("default")
	meteringDoc.SetName(m.getName())
    labels := make(map[string]string)
    labels["meter_state"] = DEFAULT_METER_LABEL
	meteringDoc.SetLabels(labels);
	return meteringDoc, nil
}

func (e *Event) getMeteringEvent(opt GenericOptions, signal int) *Metering {
	return newMetering(opt, e.crd, signal)
}

func (e *Event) getMeteringEvents() ([]*Metering, error) {
	options := e.crd.Spec.options
	lo := e.crd.Status.lastOperation
	oldAppliedOptions := e.oldCrd.Status.appliedOptions
	var meteringDocs []*Metering

	glog.Infof("Getting Metering Docs for Type %s", lo.Type)

	switch lo.Type {
	case "update":
		meteringDocs = append(meteringDocs, e.getMeteringEvent(options, METER_START))
		meteringDocs = append(meteringDocs, e.getMeteringEvent(oldAppliedOptions, METER_STOP))
	case "create":
		meteringDocs = append(meteringDocs, e.getMeteringEvent(options, METER_START))
	}
	return meteringDocs, nil
}

func (e *Event) createMertering(cfg *rest.Config) error {
	apiserver, err := getClient(cfg)
	if err != nil {
		return err
	}
	events, err := e.getMeteringEvents()
	if err != nil {
		return err
	}
	for _, evt := range events {
		unstructuredDoc, err := meteringToUnstructured(evt)
		if err != nil {
			glog.Errorf("\nError converting event : %v\n", err)
			return err
		}
		err = apiserver.Create(context.TODO(), unstructuredDoc)
		if err != nil {
			glog.Errorf("\nError creating: %v\n", err)
			return err
		}
		glog.Infof("Successfully created metering resource")
	}
	return nil
}
