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

// FakeDockers implements DockerInterface
type FakeDockers struct {
	Fake *FakeDeploymentV1alpha1
	ns   string
}

var dockersResource = schema.GroupVersionResource{Group: "deployment.servicefabrik.io", Version: "v1alpha1", Resource: "dockers"}

var dockersKind = schema.GroupVersionKind{Group: "deployment.servicefabrik.io", Version: "v1alpha1", Kind: "Docker"}

// Get takes name of the docker, and returns the corresponding docker object, and an error if there is any.
func (c *FakeDockers) Get(name string, options v1.GetOptions) (result *v1alpha1.Docker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dockersResource, c.ns, name), &v1alpha1.Docker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Docker), err
}

// List takes label and field selectors, and returns the list of Dockers that match those selectors.
func (c *FakeDockers) List(opts v1.ListOptions) (result *v1alpha1.DockerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dockersResource, dockersKind, c.ns, opts), &v1alpha1.DockerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DockerList{}
	for _, item := range obj.(*v1alpha1.DockerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dockers.
func (c *FakeDockers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dockersResource, c.ns, opts))

}

// Create takes the representation of a docker and creates it.  Returns the server's representation of the docker, and an error, if there is any.
func (c *FakeDockers) Create(docker *v1alpha1.Docker) (result *v1alpha1.Docker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dockersResource, c.ns, docker), &v1alpha1.Docker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Docker), err
}

// Update takes the representation of a docker and updates it. Returns the server's representation of the docker, and an error, if there is any.
func (c *FakeDockers) Update(docker *v1alpha1.Docker) (result *v1alpha1.Docker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dockersResource, c.ns, docker), &v1alpha1.Docker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Docker), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDockers) UpdateStatus(docker *v1alpha1.Docker) (*v1alpha1.Docker, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dockersResource, "status", c.ns, docker), &v1alpha1.Docker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Docker), err
}

// Delete takes name of the docker and deletes it. Returns an error if one occurs.
func (c *FakeDockers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dockersResource, c.ns, name), &v1alpha1.Docker{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDockers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dockersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DockerList{})
	return err
}

// Patch applies the patch and returns the patched docker.
func (c *FakeDockers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Docker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dockersResource, c.ns, name, data, subresources...), &v1alpha1.Docker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Docker), err
}
