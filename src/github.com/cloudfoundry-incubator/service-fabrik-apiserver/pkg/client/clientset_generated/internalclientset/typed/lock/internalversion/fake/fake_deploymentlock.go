//TODO copyright header
package fake

import (
	lock "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/lock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDeploymentLocks implements DeploymentLockInterface
type FakeDeploymentLocks struct {
	Fake *FakeLock
	ns   string
}

var deploymentlocksResource = schema.GroupVersionResource{Group: "lock.servicefabrik.io", Version: "", Resource: "deploymentlocks"}

var deploymentlocksKind = schema.GroupVersionKind{Group: "lock.servicefabrik.io", Version: "", Kind: "DeploymentLock"}

// Get takes name of the deploymentLock, and returns the corresponding deploymentLock object, and an error if there is any.
func (c *FakeDeploymentLocks) Get(name string, options v1.GetOptions) (result *lock.DeploymentLock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(deploymentlocksResource, c.ns, name), &lock.DeploymentLock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*lock.DeploymentLock), err
}

// List takes label and field selectors, and returns the list of DeploymentLocks that match those selectors.
func (c *FakeDeploymentLocks) List(opts v1.ListOptions) (result *lock.DeploymentLockList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(deploymentlocksResource, deploymentlocksKind, c.ns, opts), &lock.DeploymentLockList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &lock.DeploymentLockList{}
	for _, item := range obj.(*lock.DeploymentLockList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested deploymentLocks.
func (c *FakeDeploymentLocks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(deploymentlocksResource, c.ns, opts))

}

// Create takes the representation of a deploymentLock and creates it.  Returns the server's representation of the deploymentLock, and an error, if there is any.
func (c *FakeDeploymentLocks) Create(deploymentLock *lock.DeploymentLock) (result *lock.DeploymentLock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(deploymentlocksResource, c.ns, deploymentLock), &lock.DeploymentLock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*lock.DeploymentLock), err
}

// Update takes the representation of a deploymentLock and updates it. Returns the server's representation of the deploymentLock, and an error, if there is any.
func (c *FakeDeploymentLocks) Update(deploymentLock *lock.DeploymentLock) (result *lock.DeploymentLock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(deploymentlocksResource, c.ns, deploymentLock), &lock.DeploymentLock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*lock.DeploymentLock), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDeploymentLocks) UpdateStatus(deploymentLock *lock.DeploymentLock) (*lock.DeploymentLock, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(deploymentlocksResource, "status", c.ns, deploymentLock), &lock.DeploymentLock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*lock.DeploymentLock), err
}

// Delete takes name of the deploymentLock and deletes it. Returns an error if one occurs.
func (c *FakeDeploymentLocks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(deploymentlocksResource, c.ns, name), &lock.DeploymentLock{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDeploymentLocks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(deploymentlocksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &lock.DeploymentLockList{})
	return err
}

// Patch applies the patch and returns the patched deploymentLock.
func (c *FakeDeploymentLocks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *lock.DeploymentLock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(deploymentlocksResource, c.ns, name, data, subresources...), &lock.DeploymentLock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*lock.DeploymentLock), err
}
