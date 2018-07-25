//TODO copyright header
package fake

import (
	bind "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/bind"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDirectorBinds implements DirectorBindInterface
type FakeDirectorBinds struct {
	Fake *FakeBind
	ns   string
}

var directorbindsResource = schema.GroupVersionResource{Group: "bind.servicefabrik.io", Version: "", Resource: "directorbinds"}

var directorbindsKind = schema.GroupVersionKind{Group: "bind.servicefabrik.io", Version: "", Kind: "DirectorBind"}

// Get takes name of the directorBind, and returns the corresponding directorBind object, and an error if there is any.
func (c *FakeDirectorBinds) Get(name string, options v1.GetOptions) (result *bind.DirectorBind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(directorbindsResource, c.ns, name), &bind.DirectorBind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*bind.DirectorBind), err
}

// List takes label and field selectors, and returns the list of DirectorBinds that match those selectors.
func (c *FakeDirectorBinds) List(opts v1.ListOptions) (result *bind.DirectorBindList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(directorbindsResource, directorbindsKind, c.ns, opts), &bind.DirectorBindList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &bind.DirectorBindList{}
	for _, item := range obj.(*bind.DirectorBindList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested directorBinds.
func (c *FakeDirectorBinds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(directorbindsResource, c.ns, opts))

}

// Create takes the representation of a directorBind and creates it.  Returns the server's representation of the directorBind, and an error, if there is any.
func (c *FakeDirectorBinds) Create(directorBind *bind.DirectorBind) (result *bind.DirectorBind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(directorbindsResource, c.ns, directorBind), &bind.DirectorBind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*bind.DirectorBind), err
}

// Update takes the representation of a directorBind and updates it. Returns the server's representation of the directorBind, and an error, if there is any.
func (c *FakeDirectorBinds) Update(directorBind *bind.DirectorBind) (result *bind.DirectorBind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(directorbindsResource, c.ns, directorBind), &bind.DirectorBind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*bind.DirectorBind), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDirectorBinds) UpdateStatus(directorBind *bind.DirectorBind) (*bind.DirectorBind, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(directorbindsResource, "status", c.ns, directorBind), &bind.DirectorBind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*bind.DirectorBind), err
}

// Delete takes name of the directorBind and deletes it. Returns an error if one occurs.
func (c *FakeDirectorBinds) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(directorbindsResource, c.ns, name), &bind.DirectorBind{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDirectorBinds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(directorbindsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &bind.DirectorBindList{})
	return err
}

// Patch applies the patch and returns the patched directorBind.
func (c *FakeDirectorBinds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *bind.DirectorBind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(directorbindsResource, c.ns, name, data, subresources...), &bind.DirectorBind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*bind.DirectorBind), err
}
