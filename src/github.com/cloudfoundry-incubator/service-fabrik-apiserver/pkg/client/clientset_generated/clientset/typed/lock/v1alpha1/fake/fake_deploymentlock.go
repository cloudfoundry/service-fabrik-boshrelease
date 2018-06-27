//TODO copyright header
package fake

import (
	v1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/lock/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDeploymentLocks implements DeploymentLockInterface
type FakeDeploymentLocks struct {
	Fake *FakeLockV1alpha1
	ns   string
}

var deploymentlocksResource = schema.GroupVersionResource{Group: "lock.servicefabrik.io", Version: "v1alpha1", Resource: "deploymentlocks"}

var deploymentlocksKind = schema.GroupVersionKind{Group: "lock.servicefabrik.io", Version: "v1alpha1", Kind: "DeploymentLock"}

// Get takes name of the deploymentLock, and returns the corresponding deploymentLock object, and an error if there is any.
func (c *FakeDeploymentLocks) Get(name string, options v1.GetOptions) (result *v1alpha1.DeploymentLock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(deploymentlocksResource, c.ns, name), &v1alpha1.DeploymentLock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeploymentLock), err
}

// List takes label and field selectors, and returns the list of DeploymentLocks that match those selectors.
func (c *FakeDeploymentLocks) List(opts v1.ListOptions) (result *v1alpha1.DeploymentLockList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(deploymentlocksResource, deploymentlocksKind, c.ns, opts), &v1alpha1.DeploymentLockList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DeploymentLockList{}
	for _, item := range obj.(*v1alpha1.DeploymentLockList).Items {
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
func (c *FakeDeploymentLocks) Create(deploymentLock *v1alpha1.DeploymentLock) (result *v1alpha1.DeploymentLock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(deploymentlocksResource, c.ns, deploymentLock), &v1alpha1.DeploymentLock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeploymentLock), err
}

// Update takes the representation of a deploymentLock and updates it. Returns the server's representation of the deploymentLock, and an error, if there is any.
func (c *FakeDeploymentLocks) Update(deploymentLock *v1alpha1.DeploymentLock) (result *v1alpha1.DeploymentLock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(deploymentlocksResource, c.ns, deploymentLock), &v1alpha1.DeploymentLock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeploymentLock), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDeploymentLocks) UpdateStatus(deploymentLock *v1alpha1.DeploymentLock) (*v1alpha1.DeploymentLock, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(deploymentlocksResource, "status", c.ns, deploymentLock), &v1alpha1.DeploymentLock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeploymentLock), err
}

// Delete takes name of the deploymentLock and deletes it. Returns an error if one occurs.
func (c *FakeDeploymentLocks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(deploymentlocksResource, c.ns, name), &v1alpha1.DeploymentLock{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDeploymentLocks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(deploymentlocksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DeploymentLockList{})
	return err
}

// Patch applies the patch and returns the patched deploymentLock.
func (c *FakeDeploymentLocks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DeploymentLock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(deploymentlocksResource, c.ns, name, data, subresources...), &v1alpha1.DeploymentLock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeploymentLock), err
}
