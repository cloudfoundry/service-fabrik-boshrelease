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

// FakeDirectors implements DirectorInterface
type FakeDirectors struct {
	Fake *FakeDeploymentV1alpha1
	ns   string
}

var directorsResource = schema.GroupVersionResource{Group: "deployment.servicefabrik.io", Version: "v1alpha1", Resource: "directors"}

var directorsKind = schema.GroupVersionKind{Group: "deployment.servicefabrik.io", Version: "v1alpha1", Kind: "Director"}

// Get takes name of the director, and returns the corresponding director object, and an error if there is any.
func (c *FakeDirectors) Get(name string, options v1.GetOptions) (result *v1alpha1.Director, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(directorsResource, c.ns, name), &v1alpha1.Director{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Director), err
}

// List takes label and field selectors, and returns the list of Directors that match those selectors.
func (c *FakeDirectors) List(opts v1.ListOptions) (result *v1alpha1.DirectorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(directorsResource, directorsKind, c.ns, opts), &v1alpha1.DirectorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DirectorList{}
	for _, item := range obj.(*v1alpha1.DirectorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested directors.
func (c *FakeDirectors) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(directorsResource, c.ns, opts))

}

// Create takes the representation of a director and creates it.  Returns the server's representation of the director, and an error, if there is any.
func (c *FakeDirectors) Create(director *v1alpha1.Director) (result *v1alpha1.Director, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(directorsResource, c.ns, director), &v1alpha1.Director{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Director), err
}

// Update takes the representation of a director and updates it. Returns the server's representation of the director, and an error, if there is any.
func (c *FakeDirectors) Update(director *v1alpha1.Director) (result *v1alpha1.Director, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(directorsResource, c.ns, director), &v1alpha1.Director{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Director), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDirectors) UpdateStatus(director *v1alpha1.Director) (*v1alpha1.Director, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(directorsResource, "status", c.ns, director), &v1alpha1.Director{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Director), err
}

// Delete takes name of the director and deletes it. Returns an error if one occurs.
func (c *FakeDirectors) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(directorsResource, c.ns, name), &v1alpha1.Director{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDirectors) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(directorsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DirectorList{})
	return err
}

// Patch applies the patch and returns the patched director.
func (c *FakeDirectors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Director, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(directorsResource, c.ns, name, data, subresources...), &v1alpha1.Director{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Director), err
}
