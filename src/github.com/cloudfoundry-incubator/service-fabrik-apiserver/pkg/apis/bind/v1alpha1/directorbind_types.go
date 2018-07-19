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

// DirectorBind
// +k8s:openapi-gen=true
// +resource:path=directorbinds,strategy=DirectorBindStrategy
type DirectorBind struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DirectorBindSpec   `json:"spec,omitempty"`
	Status DirectorBindStatus `json:"status,omitempty"`
}

// DirectorBindSpec defines the desired state of DirectorBind
type DirectorBindSpec struct {
	Instance string `json:"instance,omitempty"`
	Options  string `json:"options,omitempty"`
}

// DirectorBindStatus defines the observed state of DirectorBind
type DirectorBindStatus struct {
	State    string `json:"state,omitempty"`
	Response string `json:"response,omitempty"`
}

// Validate checks that an instance of DirectorBind is well formed
func (DirectorBindStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*bind.DirectorBind)
	log.Printf("Validating fields for DirectorBind %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default DirectorBind field values
func (DirectorBindSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*DirectorBind)
	// set default field values here
	log.Printf("Defaulting fields for DirectorBind %s\n", obj.Name)
}

// PrepareForUpdate sets the status labels during status update.
func (s DirectorBindStatusStrategy) PrepareForUpdate(ctx request.Context, obj, old runtime.Object) {
	s.DefaultStatusStorageStrategy.PrepareForUpdate(ctx, obj, old)
	dNew := obj.(*bind.DirectorBind)
	labels := dNew.GetObjectMeta().GetLabels()
	if labels == nil {
		labels = make(map[string]string)
		dNew.GetObjectMeta().SetLabels(labels)
	}
	labels["state"] = dNew.Status.State
}
