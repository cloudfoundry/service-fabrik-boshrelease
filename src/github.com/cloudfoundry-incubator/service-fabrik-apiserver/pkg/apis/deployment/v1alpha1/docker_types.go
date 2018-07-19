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

// Docker
// +k8s:openapi-gen=true
// +resource:path=dockers,strategy=DockerStrategy
type Docker struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DockerSpec   `json:"spec,omitempty"`
	Status DockerStatus `json:"status,omitempty"`
}

// DockerSpec defines the desired state of Docker
type DockerSpec struct {
	Options string `json:"options,omitempty"`
}

// DockerStatus defines the observed state of Docker
type DockerStatus struct {
	State         string `json:"state,omitempty"`
	LastOperation string `json:"lastOperation,omitempty"`
	Response      string `json:"response,omitempty"`
}

// Validate checks that an instance of Docker is well formed
func (DockerStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*deployment.Docker)
	log.Printf("Validating fields for Docker %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default Docker field values
func (DockerSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*Docker)
	// set default field values here
	log.Printf("Defaulting fields for Docker %s\n", obj.Name)
}

// PrepareForUpdate sets the status labels during status update.
func (s DockerStatusStrategy) PrepareForUpdate(ctx request.Context, obj, old runtime.Object) {
	s.DefaultStatusStorageStrategy.PrepareForUpdate(ctx, obj, old)
	dNew := obj.(*deployment.Docker)
	labels := dNew.GetObjectMeta().GetLabels()
	if labels == nil {
		labels = make(map[string]string)
		dNew.GetObjectMeta().SetLabels(labels)
	}
	labels["state"] = dNew.Status.State
}
