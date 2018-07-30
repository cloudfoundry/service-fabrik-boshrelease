
//TODO copyright header


package v1alpha1

import (
	"log"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/deployment"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Virtualhost
// +k8s:openapi-gen=true
// +resource:path=virtualhosts,strategy=VirtualhostStrategy
type Virtualhost struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualhostSpec   `json:"spec,omitempty"`
	Status VirtualhostStatus `json:"status,omitempty"`
}

// VirtualhostSpec defines the desired state of Virtualhost
type VirtualhostSpec struct {
	Options string `json:"options,omitempty"`
}

// VirtualhostStatus defines the observed state of Virtualhost
type VirtualhostStatus struct {
	State         string `json:"state,omitempty"`
        Error         string `json:"error,omitempty"`
	LastOperation string `json:"lastOperation,omitempty"`
	Response      string `json:"response,omitempty"`
}

// Validate checks that an instance of Virtualhost is well formed
func (VirtualhostStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*deployment.Virtualhost)
	log.Printf("Validating fields for Virtualhost %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default Virtualhost field values
func (VirtualhostSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*Virtualhost)
	// set default field values here
	log.Printf("Defaulting fields for Virtualhost %s\n", obj.Name)
}

// PrepareForUpdate sets the status labels during status update.
func (s VirtualhostStatusStrategy) PrepareForUpdate(ctx request.Context, obj, old runtime.Object) {
	s.DefaultStatusStorageStrategy.PrepareForUpdate(ctx, obj, old)
	dNew := obj.(*deployment.Virtualhost)
	labels := dNew.GetObjectMeta().GetLabels()
	if labels == nil {
		labels = make(map[string]string)
		dNew.GetObjectMeta().SetLabels(labels)
	}
	labels["state"] = dNew.Status.State
}
