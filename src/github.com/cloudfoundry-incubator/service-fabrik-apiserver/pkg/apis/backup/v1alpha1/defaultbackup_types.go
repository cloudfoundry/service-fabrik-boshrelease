//TODO copyright header

package v1alpha1

import (
	"log"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/backup"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DefaultBackup
// +k8s:openapi-gen=true
// +resource:path=defaultbackups,strategy=DefaultBackupStrategy
type DefaultBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DefaultBackupSpec   `json:"spec,omitempty"`
	Status DefaultBackupStatus `json:"status,omitempty"`
}

// DefaultBackupSpec defines the desired state of DefaultBackup
type DefaultBackupSpec struct {
	Options string `json:"options,omitempty"`
}

// DefaultBackupStatus defines the observed state of DefaultBackup
type DefaultBackupStatus struct {
	State         string `json:"state,omitempty"`
	Error         string `json:"error,omitempty"`
	LastOperation string `json:"lastOperation,omitempty"`
	Response      string `json:"response,omitempty"`
}

// Validate checks that an instance of DefaultBackup is well formed
func (DefaultBackupStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*backup.DefaultBackup)
	log.Printf("Validating fields for DefaultBackup %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default DefaultBackup field values
func (DefaultBackupSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*DefaultBackup)
	// set default field values here
	log.Printf("Defaulting fields for DefaultBackup %s\n", obj.Name)
}

// PrepareForUpdate sets the status labels during status update.
func (s DefaultBackupStatusStrategy) PrepareForUpdate(ctx request.Context, obj, old runtime.Object) {
	s.DefaultStatusStorageStrategy.PrepareForUpdate(ctx, obj, old)
	dNew := obj.(*backup.DefaultBackup)
	labels := dNew.GetObjectMeta().GetLabels()
	if labels == nil {
		labels = make(map[string]string)
		dNew.GetObjectMeta().SetLabels(labels)
	}
	labels["state"] = dNew.Status.State
}
