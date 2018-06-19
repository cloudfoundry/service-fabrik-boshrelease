//TODO copyright header
package fake

import (
	clientset "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/clientset"
	backupv1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/clientset/typed/backup/v1alpha1"
	fakebackupv1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/clientset/typed/backup/v1alpha1/fake"
	deploymentv1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/clientset/typed/deployment/v1alpha1"
	fakedeploymentv1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/clientset/typed/deployment/v1alpha1/fake"
	lockv1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/clientset/typed/lock/v1alpha1"
	fakelockv1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/clientset/typed/lock/v1alpha1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	fakePtr := testing.Fake{}
	fakePtr.AddReactor("*", "*", testing.ObjectReaction(o))
	fakePtr.AddWatchReactor("*", testing.DefaultWatchReactor(watch.NewFake(), nil))

	return &Clientset{fakePtr, &fakediscovery.FakeDiscovery{Fake: &fakePtr}}
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

var _ clientset.Interface = &Clientset{}

// BackupV1alpha1 retrieves the BackupV1alpha1Client
func (c *Clientset) BackupV1alpha1() backupv1alpha1.BackupV1alpha1Interface {
	return &fakebackupv1alpha1.FakeBackupV1alpha1{Fake: &c.Fake}
}

// Backup retrieves the BackupV1alpha1Client
func (c *Clientset) Backup() backupv1alpha1.BackupV1alpha1Interface {
	return &fakebackupv1alpha1.FakeBackupV1alpha1{Fake: &c.Fake}
}

// DeploymentV1alpha1 retrieves the DeploymentV1alpha1Client
func (c *Clientset) DeploymentV1alpha1() deploymentv1alpha1.DeploymentV1alpha1Interface {
	return &fakedeploymentv1alpha1.FakeDeploymentV1alpha1{Fake: &c.Fake}
}

// Deployment retrieves the DeploymentV1alpha1Client
func (c *Clientset) Deployment() deploymentv1alpha1.DeploymentV1alpha1Interface {
	return &fakedeploymentv1alpha1.FakeDeploymentV1alpha1{Fake: &c.Fake}
}

// LockV1alpha1 retrieves the LockV1alpha1Client
func (c *Clientset) LockV1alpha1() lockv1alpha1.LockV1alpha1Interface {
	return &fakelockv1alpha1.FakeLockV1alpha1{Fake: &c.Fake}
}

// Lock retrieves the LockV1alpha1Client
func (c *Clientset) Lock() lockv1alpha1.LockV1alpha1Interface {
	return &fakelockv1alpha1.FakeLockV1alpha1{Fake: &c.Fake}
}
