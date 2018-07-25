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

// DockerBind
// +k8s:openapi-gen=true
// +resource:path=dockerbinds,strategy=DockerBindStrategy
type DockerBind struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DockerBindSpec   `json:"spec,omitempty"`
	Status DockerBindStatus `json:"status,omitempty"`
}

// DockerBindSpec defines the desired state of DockerBind
type DockerBindSpec struct {
	Instance string `json:"instance,omitempty"`
	Options  string `json:"options,omitempty"`
}

// DockerBindStatus defines the observed state of DockerBind
type DockerBindStatus struct {
	State    string `json:"state,omitempty"`
	Response string `json:"response,omitempty"`
}

// Validate checks that an instance of DockerBind is well formed
func (DockerBindStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*bind.DockerBind)
	log.Printf("Validating fields for DockerBind %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default DockerBind field values
func (DockerBindSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*DockerBind)
	// set default field values here
	log.Printf("Defaulting fields for DockerBind %s\n", obj.Name)
}

// PrepareForUpdate sets the status labels during status update.
func (s DockerBindStatusStrategy) PrepareForUpdate(ctx request.Context, obj, old runtime.Object) {
	s.DefaultStatusStorageStrategy.PrepareForUpdate(ctx, obj, old)
	dNew := obj.(*bind.DockerBind)
	labels := dNew.GetObjectMeta().GetLabels()
	if labels == nil {
		labels = make(map[string]string)
		dNew.GetObjectMeta().SetLabels(labels)
	}
	labels["state"] = dNew.Status.State
}
