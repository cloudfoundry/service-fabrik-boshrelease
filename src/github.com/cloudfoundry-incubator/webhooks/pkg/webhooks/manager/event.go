package main

import (
	"context"
	"encoding/json"

	"k8s.io/client-go/rest"

	"github.com/golang/glog"
	"github.com/google/uuid"
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
		glog.Errorf("Could not get the GenericResource object", err)
		return nil, err
	}
	oldCrd, err := getGenericResource(ar.Request.OldObject.Raw)
	if err != nil {
		glog.Errorf("Could not get the old GenericResource object", err)
		return nil, err
	}
	return &Event{
		AdmissionReview: ar,
		crd:             crd,
		oldCrd:          oldCrd,
	}, nil
}

func (e *Event) isMeteringEvent() bool {
	loNew := getLastOperation(e.crd)
	loOld := getLastOperation(e.oldCrd)
	glog.Infof("New: type: %s, state: %s\n", loNew.Type, loNew.State)
	glog.Infof("Old: type: %s, state: %s\n", loOld.Type, loOld.State)

	if loNew.Type == loOld.Type && loNew.State != loOld.State {
		if loNew.Type == "update" || loNew.Type == "create" {
			if loNew.State == "succeeded" {
				return true
			}

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
		glog.Errorf("unable to set up overall controller manager", err)
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

func (e *Event) getDoc(opt GenericOptions, lo GenericLastOperation, crd GenericResource, signal string) (*unstructured.Unstructured, error) {
	m := Metering{
		Spec: MeteringSpec{
			Options: MeteringOptions{
				ServiceID:  opt.ServiceId,
				PlanID:     opt.PlanId,
				InstanceID: e.crd.Name,
				OrgID:      opt.Context.OrganizationGuid,
				SpaceID:    opt.Context.SpaceGuid,
				Type:       lo.Type,
			},
		},
	}
	values, err := ObjectToMapInterface(m)
	if err != nil {
		glog.Errorf("unable convert to map interface %v", err)
		return nil, err
	}
	meteringDoc := &unstructured.Unstructured{}
	meteringDoc.SetUnstructuredContent(values)
	meteringDoc.SetKind("Event")
	meteringDoc.SetAPIVersion("metering.servicefabrik.io/v1alpha1")
	meteringDoc.SetNamespace("default")
	name := uuid.New().String()
	meteringDoc.SetName(name)
	glog.Infof("Creating metering doc of Type: %s Signal: %s, uuid: %s", lo.Type, signal, name)
	return meteringDoc, nil
}

func (e *Event) getDocs() ([]*unstructured.Unstructured, error) {
	opt := getOptions(e.crd)
	lo := getLastOperation(e.crd)
	oldOpt := getOptions(e.oldCrd)
	oldLo := getLastOperation(e.oldCrd)
	var meteringDocs []*unstructured.Unstructured

	glog.Infof("Getting Metering Docs for Type %s", lo.Type)

	switch lo.Type {
	case "update":
		meteringDoc, err := e.getDoc(opt, lo, e.crd, "start")
		if err != nil {
			glog.Errorf("\nError getting: %v\n", err)
			return nil, err
		}
		meteringDocs = append(meteringDocs, meteringDoc)
		meteringDoc, err = e.getDoc(oldOpt, oldLo, e.oldCrd, "stop")
		if err != nil {
			glog.Errorf("\nError getting: %v\n", err)
			return nil, err
		}
		meteringDocs = append(meteringDocs, meteringDoc)
	case "create":
		meteringDoc, err := e.getDoc(opt, lo, e.crd, "start")
		if err != nil {
			glog.Errorf("\nError getting: %v\n", err)
			return nil, err
		}
		meteringDocs = append(meteringDocs, meteringDoc)
	}
	return meteringDocs, nil
}

func (e *Event) createMertering(cfg *rest.Config) error {
	apiserver, err := getClient(cfg)
	if err != nil {
		return err
	}
	docs, err := e.getDocs()
	if err != nil {
		return err
	}
	for _, doc := range docs {
		err := apiserver.Create(context.TODO(), doc)
		if err != nil {
			glog.Errorf("\nError creating: %v\n", err)
			return err
		}
		glog.Infof("Successfully created metering resource")
	}
	return nil
}
