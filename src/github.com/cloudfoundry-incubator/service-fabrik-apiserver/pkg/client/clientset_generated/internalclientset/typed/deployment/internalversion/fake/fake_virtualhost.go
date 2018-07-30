//TODO copyright header
package fake

import (
	deployment "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/deployment"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVirtualhosts implements VirtualhostInterface
type FakeVirtualhosts struct {
	Fake *FakeDeployment
	ns   string
}

var virtualhostsResource = schema.GroupVersionResource{Group: "deployment.servicefabrik.io", Version: "", Resource: "virtualhosts"}

var virtualhostsKind = schema.GroupVersionKind{Group: "deployment.servicefabrik.io", Version: "", Kind: "Virtualhost"}

// Get takes name of the virtualhost, and returns the corresponding virtualhost object, and an error if there is any.
func (c *FakeVirtualhosts) Get(name string, options v1.GetOptions) (result *deployment.Virtualhost, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualhostsResource, c.ns, name), &deployment.Virtualhost{})

	if obj == nil {
		return nil, err
	}
	return obj.(*deployment.Virtualhost), err
}

// List takes label and field selectors, and returns the list of Virtualhosts that match those selectors.
func (c *FakeVirtualhosts) List(opts v1.ListOptions) (result *deployment.VirtualhostList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualhostsResource, virtualhostsKind, c.ns, opts), &deployment.VirtualhostList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &deployment.VirtualhostList{}
	for _, item := range obj.(*deployment.VirtualhostList).Items {
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
func (c *FakeVirtualhosts) Create(virtualhost *deployment.Virtualhost) (result *deployment.Virtualhost, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualhostsResource, c.ns, virtualhost), &deployment.Virtualhost{})

	if obj == nil {
		return nil, err
	}
	return obj.(*deployment.Virtualhost), err
}

// Update takes the representation of a virtualhost and updates it. Returns the server's representation of the virtualhost, and an error, if there is any.
func (c *FakeVirtualhosts) Update(virtualhost *deployment.Virtualhost) (result *deployment.Virtualhost, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualhostsResource, c.ns, virtualhost), &deployment.Virtualhost{})

	if obj == nil {
		return nil, err
	}
	return obj.(*deployment.Virtualhost), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualhosts) UpdateStatus(virtualhost *deployment.Virtualhost) (*deployment.Virtualhost, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualhostsResource, "status", c.ns, virtualhost), &deployment.Virtualhost{})

	if obj == nil {
		return nil, err
	}
	return obj.(*deployment.Virtualhost), err
}

// Delete takes name of the virtualhost and deletes it. Returns an error if one occurs.
func (c *FakeVirtualhosts) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualhostsResource, c.ns, name), &deployment.Virtualhost{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualhosts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualhostsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &deployment.VirtualhostList{})
	return err
}

// Patch applies the patch and returns the patched virtualhost.
func (c *FakeVirtualhosts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *deployment.Virtualhost, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualhostsResource, c.ns, name, data, subresources...), &deployment.Virtualhost{})

	if obj == nil {
		return nil, err
	}
	return obj.(*deployment.Virtualhost), err
}
