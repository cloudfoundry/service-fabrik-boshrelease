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

// FakeDockers implements DockerInterface
type FakeDockers struct {
	Fake *FakeDeployment
	ns   string
}

var dockersResource = schema.GroupVersionResource{Group: "deployment.servicefabrik.io", Version: "", Resource: "dockers"}

var dockersKind = schema.GroupVersionKind{Group: "deployment.servicefabrik.io", Version: "", Kind: "Docker"}

// Get takes name of the docker, and returns the corresponding docker object, and an error if there is any.
func (c *FakeDockers) Get(name string, options v1.GetOptions) (result *deployment.Docker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dockersResource, c.ns, name), &deployment.Docker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*deployment.Docker), err
}

// List takes label and field selectors, and returns the list of Dockers that match those selectors.
func (c *FakeDockers) List(opts v1.ListOptions) (result *deployment.DockerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dockersResource, dockersKind, c.ns, opts), &deployment.DockerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &deployment.DockerList{}
	for _, item := range obj.(*deployment.DockerList).Items {
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
func (c *FakeDockers) Create(docker *deployment.Docker) (result *deployment.Docker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dockersResource, c.ns, docker), &deployment.Docker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*deployment.Docker), err
}

// Update takes the representation of a docker and updates it. Returns the server's representation of the docker, and an error, if there is any.
func (c *FakeDockers) Update(docker *deployment.Docker) (result *deployment.Docker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dockersResource, c.ns, docker), &deployment.Docker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*deployment.Docker), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDockers) UpdateStatus(docker *deployment.Docker) (*deployment.Docker, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dockersResource, "status", c.ns, docker), &deployment.Docker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*deployment.Docker), err
}

// Delete takes name of the docker and deletes it. Returns an error if one occurs.
func (c *FakeDockers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dockersResource, c.ns, name), &deployment.Docker{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDockers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dockersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &deployment.DockerList{})
	return err
}

// Patch applies the patch and returns the patched docker.
func (c *FakeDockers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *deployment.Docker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dockersResource, c.ns, name, data, subresources...), &deployment.Docker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*deployment.Docker), err
}
