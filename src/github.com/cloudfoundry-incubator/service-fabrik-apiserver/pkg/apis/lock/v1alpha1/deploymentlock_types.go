
//TODO copyright header


package v1alpha1

import (
	"log"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/lock"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DeploymentLock
// +k8s:openapi-gen=true
// +resource:path=deploymentlocks,strategy=DeploymentLockStrategy
type DeploymentLock struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeploymentLockSpec   `json:"spec,omitempty"`
	Status DeploymentLockStatus `json:"status,omitempty"`
}

// DeploymentLockSpec defines the desired state of DeploymentLock
type DeploymentLockSpec struct {
	Options string `json:"options,omitempty"`
}

// DeploymentLockStatus defines the observed state of DeploymentLock
type DeploymentLockStatus struct {
}

// Validate checks that an instance of DeploymentLock is well formed
func (DeploymentLockStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*lock.DeploymentLock)
	log.Printf("Validating fields for DeploymentLock %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default DeploymentLock field values
func (DeploymentLockSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*DeploymentLock)
	// set default field values here
	log.Printf("Defaulting fields for DeploymentLock %s\n", obj.Name)
}
