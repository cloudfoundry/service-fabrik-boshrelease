//TODO copyright header
package fake

import (
	v1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/bind/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDockerBinds implements DockerBindInterface
type FakeDockerBinds struct {
	Fake *FakeBindV1alpha1
	ns   string
}

var dockerbindsResource = schema.GroupVersionResource{Group: "bind.servicefabrik.io", Version: "v1alpha1", Resource: "dockerbinds"}

var dockerbindsKind = schema.GroupVersionKind{Group: "bind.servicefabrik.io", Version: "v1alpha1", Kind: "DockerBind"}

// Get takes name of the dockerBind, and returns the corresponding dockerBind object, and an error if there is any.
func (c *FakeDockerBinds) Get(name string, options v1.GetOptions) (result *v1alpha1.DockerBind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dockerbindsResource, c.ns, name), &v1alpha1.DockerBind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DockerBind), err
}

// List takes label and field selectors, and returns the list of DockerBinds that match those selectors.
func (c *FakeDockerBinds) List(opts v1.ListOptions) (result *v1alpha1.DockerBindList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dockerbindsResource, dockerbindsKind, c.ns, opts), &v1alpha1.DockerBindList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DockerBindList{}
	for _, item := range obj.(*v1alpha1.DockerBindList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dockerBinds.
func (c *FakeDockerBinds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dockerbindsResource, c.ns, opts))

}

// Create takes the representation of a dockerBind and creates it.  Returns the server's representation of the dockerBind, and an error, if there is any.
func (c *FakeDockerBinds) Create(dockerBind *v1alpha1.DockerBind) (result *v1alpha1.DockerBind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dockerbindsResource, c.ns, dockerBind), &v1alpha1.DockerBind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DockerBind), err
}

// Update takes the representation of a dockerBind and updates it. Returns the server's representation of the dockerBind, and an error, if there is any.
func (c *FakeDockerBinds) Update(dockerBind *v1alpha1.DockerBind) (result *v1alpha1.DockerBind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dockerbindsResource, c.ns, dockerBind), &v1alpha1.DockerBind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DockerBind), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDockerBinds) UpdateStatus(dockerBind *v1alpha1.DockerBind) (*v1alpha1.DockerBind, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dockerbindsResource, "status", c.ns, dockerBind), &v1alpha1.DockerBind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DockerBind), err
}

// Delete takes name of the dockerBind and deletes it. Returns an error if one occurs.
func (c *FakeDockerBinds) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dockerbindsResource, c.ns, name), &v1alpha1.DockerBind{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDockerBinds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dockerbindsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DockerBindList{})
	return err
}

// Patch applies the patch and returns the patched dockerBind.
func (c *FakeDockerBinds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DockerBind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dockerbindsResource, c.ns, name, data, subresources...), &v1alpha1.DockerBind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DockerBind), err
}
