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

// Director
// +k8s:openapi-gen=true
// +resource:path=directors,strategy=DirectorStrategy
type Director struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DirectorSpec   `json:"spec,omitempty"`
	Status DirectorStatus `json:"status,omitempty"`
}

// DirectorSpec defines the desired state of Director
type DirectorSpec struct {
	Options string `json:"options,omitempty"`
}

// DirectorStatus defines the observed state of Director
type DirectorStatus struct {
	State         string `json:"state,omitempty"`
	LastOperation string `json:"lastOperation,omitempty"`
	Response string	`json:"response,omitempty"`
}

// Validate checks that an instance of Director is well formed
func (DirectorStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*deployment.Director)
	log.Printf("Validating fields for Director %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default Director field values
func (DirectorSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*Director)
	// set default field values here
	log.Printf("Defaulting fields for Director %s\n", obj.Name)
}

// PrepareForUpdate sets the status labels during status update.
func (s DirectorStatusStrategy) PrepareForUpdate(ctx request.Context, obj, old runtime.Object) {
	s.DefaultStatusStorageStrategy.PrepareForUpdate(ctx, obj, old)
	dNew := obj.(*deployment.Director)
	labels := dNew.GetObjectMeta().GetLabels()
	if labels == nil {
		labels = make(map[string]string)
		dNew.GetObjectMeta().SetLabels(labels)
	}
	labels["state"] = dNew.Status.State
}
