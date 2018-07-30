//TODO copyright header
package fake

import (
	v1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/deployment/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVirtualhosts implements VirtualhostInterface
type FakeVirtualhosts struct {
	Fake *FakeDeploymentV1alpha1
	ns   string
}

var virtualhostsResource = schema.GroupVersionResource{Group: "deployment.servicefabrik.io", Version: "v1alpha1", Resource: "virtualhosts"}

var virtualhostsKind = schema.GroupVersionKind{Group: "deployment.servicefabrik.io", Version: "v1alpha1", Kind: "Virtualhost"}

// Get takes name of the virtualhost, and returns the corresponding virtualhost object, and an error if there is any.
func (c *FakeVirtualhosts) Get(name string, options v1.GetOptions) (result *v1alpha1.Virtualhost, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualhostsResource, c.ns, name), &v1alpha1.Virtualhost{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Virtualhost), err
}

// List takes label and field selectors, and returns the list of Virtualhosts that match those selectors.
func (c *FakeVirtualhosts) List(opts v1.ListOptions) (result *v1alpha1.VirtualhostList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualhostsResource, virtualhostsKind, c.ns, opts), &v1alpha1.VirtualhostList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VirtualhostList{}
	for _, item := range obj.(*v1alpha1.VirtualhostList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualhosts.
func (c *FakeVirtualhosts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualhostsResource, c.ns, opts))

}

// Create takes the representation of a virtualhost and creates it.  Returns the server's representation of the virtualhost, and an error, if there is any.
func (c *FakeVirtualhosts) Create(virtualhost *v1alpha1.Virtualhost) (result *v1alpha1.Virtualhost, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualhostsResource, c.ns, virtualhost), &v1alpha1.Virtualhost{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Virtualhost), err
}

// Update takes the representation of a virtualhost and updates it. Returns the server's representation of the virtualhost, and an error, if there is any.
func (c *FakeVirtualhosts) Update(virtualhost *v1alpha1.Virtualhost) (result *v1alpha1.Virtualhost, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualhostsResource, c.ns, virtualhost), &v1alpha1.Virtualhost{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Virtualhost), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualhosts) UpdateStatus(virtualhost *v1alpha1.Virtualhost) (*v1alpha1.Virtualhost, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualhostsResource, "status", c.ns, virtualhost), &v1alpha1.Virtualhost{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Virtualhost), err
}

// Delete takes name of the virtualhost and deletes it. Returns an error if one occurs.
func (c *FakeVirtualhosts) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualhostsResource, c.ns, name), &v1alpha1.Virtualhost{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualhosts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualhostsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.VirtualhostList{})
	return err
}

// Patch applies the patch and returns the patched virtualhost.
func (c *FakeVirtualhosts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Virtualhost, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualhostsResource, c.ns, name, data, subresources...), &v1alpha1.Virtualhost{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Virtualhost), err
}
