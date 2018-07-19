//TODO copyright header
package v1alpha1

import (
	v1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/deployment/v1alpha1"
	scheme "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DockersGetter has a method to return a DockerInterface.
// A group's client should implement this interface.
type DockersGetter interface {
	Dockers(namespace string) DockerInterface
}

// DockerInterface has methods to work with Docker resources.
type DockerInterface interface {
	Create(*v1alpha1.Docker) (*v1alpha1.Docker, error)
	Update(*v1alpha1.Docker) (*v1alpha1.Docker, error)
	UpdateStatus(*v1alpha1.Docker) (*v1alpha1.Docker, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Docker, error)
	List(opts v1.ListOptions) (*v1alpha1.DockerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Docker, err error)
	DockerExpansion
}

// dockers implements DockerInterface
type dockers struct {
	client rest.Interface
	ns     string
}

// newDockers returns a Dockers
func newDockers(c *DeploymentV1alpha1Client, namespace string) *dockers {
	return &dockers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the docker, and returns the corresponding docker object, and an error if there is any.
func (c *dockers) Get(name string, options v1.GetOptions) (result *v1alpha1.Docker, err error) {
	result = &v1alpha1.Docker{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dockers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Dockers that match those selectors.
func (c *dockers) List(opts v1.ListOptions) (result *v1alpha1.DockerList, err error) {
	result = &v1alpha1.DockerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dockers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dockers.
func (c *dockers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dockers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a docker and creates it.  Returns the server's representation of the docker, and an error, if there is any.
func (c *dockers) Create(docker *v1alpha1.Docker) (result *v1alpha1.Docker, err error) {
	result = &v1alpha1.Docker{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dockers").
		Body(docker).
		Do().
		Into(result)
	return
}

// Update takes the representation of a docker and updates it. Returns the server's representation of the docker, and an error, if there is any.
func (c *dockers) Update(docker *v1alpha1.Docker) (result *v1alpha1.Docker, err error) {
	result = &v1alpha1.Docker{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dockers").
		Name(docker.Name).
		Body(docker).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dockers) UpdateStatus(docker *v1alpha1.Docker) (result *v1alpha1.Docker, err error) {
	result = &v1alpha1.Docker{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dockers").
		Name(docker.Name).
		SubResource("status").
		Body(docker).
		Do().
		Into(result)
	return
}

// Delete takes name of the docker and deletes it. Returns an error if one occurs.
func (c *dockers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dockers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dockers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dockers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched docker.
func (c *dockers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Docker, err error) {
	result = &v1alpha1.Docker{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dockers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
