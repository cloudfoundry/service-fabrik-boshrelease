
//TODO copyright header


package v1alpha1

import (
	"log"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/bind"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Virtualhostbind
// +k8s:openapi-gen=true
// +resource:path=virtualhostbinds,strategy=VirtualhostbindStrategy
type Virtualhostbind struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualhostbindSpec   `json:"spec,omitempty"`
	Status VirtualhostbindStatus `json:"status,omitempty"`
}

// VirtualhostbindSpec defines the desired state of Virtualhostbind
type VirtualhostbindSpec struct {
	Instance string `json:"instance,omitempty"`
	Options  string `json:"options,omitempty"`
}

// VirtualhostbindStatus defines the observed state of Virtualhostbind
type VirtualhostbindStatus struct {
	State    string `json:"state,omitempty"`
	Response string `json:"response,omitempty"`
        Error    string `json:"error,omitempty"`
}

// Validate checks that an instance of Virtualhostbind is well formed
func (VirtualhostbindStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*bind.Virtualhostbind)
	log.Printf("Validating fields for Virtualhostbind %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default Virtualhostbind field values
func (VirtualhostbindSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*Virtualhostbind)
	// set default field values here
	log.Printf("Defaulting fields for Virtualhostbind %s\n", obj.Name)
}
// PrepareForUpdate sets the status labels during status update.
func (s VirtualhostbindStatusStrategy) PrepareForUpdate(ctx request.Context, obj, old runtime.Object) {
	s.DefaultStatusStorageStrategy.PrepareForUpdate(ctx, obj, old)
	dNew := obj.(*bind.Virtualhostbind)
	labels := dNew.GetObjectMeta().GetLabels()
	if labels == nil {
		labels = make(map[string]string)
		dNew.GetObjectMeta().SetLabels(labels)
	}
	labels["state"] = dNew.Status.State
}
