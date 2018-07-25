//TODO copyright header
package v1alpha1

import (
	v1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/bind/v1alpha1"
	scheme "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DockerBindsGetter has a method to return a DockerBindInterface.
// A group's client should implement this interface.
type DockerBindsGetter interface {
	DockerBinds(namespace string) DockerBindInterface
}

// DockerBindInterface has methods to work with DockerBind resources.
type DockerBindInterface interface {
	Create(*v1alpha1.DockerBind) (*v1alpha1.DockerBind, error)
	Update(*v1alpha1.DockerBind) (*v1alpha1.DockerBind, error)
	UpdateStatus(*v1alpha1.DockerBind) (*v1alpha1.DockerBind, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DockerBind, error)
	List(opts v1.ListOptions) (*v1alpha1.DockerBindList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DockerBind, err error)
	DockerBindExpansion
}

// dockerBinds implements DockerBindInterface
type dockerBinds struct {
	client rest.Interface
	ns     string
}

// newDockerBinds returns a DockerBinds
func newDockerBinds(c *BindV1alpha1Client, namespace string) *dockerBinds {
	return &dockerBinds{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dockerBind, and returns the corresponding dockerBind object, and an error if there is any.
func (c *dockerBinds) Get(name string, options v1.GetOptions) (result *v1alpha1.DockerBind, err error) {
	result = &v1alpha1.DockerBind{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dockerbinds").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DockerBinds that match those selectors.
func (c *dockerBinds) List(opts v1.ListOptions) (result *v1alpha1.DockerBindList, err error) {
	result = &v1alpha1.DockerBindList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dockerbinds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dockerBinds.
func (c *dockerBinds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dockerbinds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a dockerBind and creates it.  Returns the server's representation of the dockerBind, and an error, if there is any.
func (c *dockerBinds) Create(dockerBind *v1alpha1.DockerBind) (result *v1alpha1.DockerBind, err error) {
	result = &v1alpha1.DockerBind{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dockerbinds").
		Body(dockerBind).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dockerBind and updates it. Returns the server's representation of the dockerBind, and an error, if there is any.
func (c *dockerBinds) Update(dockerBind *v1alpha1.DockerBind) (result *v1alpha1.DockerBind, err error) {
	result = &v1alpha1.DockerBind{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dockerbinds").
		Name(dockerBind.Name).
		Body(dockerBind).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dockerBinds) UpdateStatus(dockerBind *v1alpha1.DockerBind) (result *v1alpha1.DockerBind, err error) {
	result = &v1alpha1.DockerBind{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dockerbinds").
		Name(dockerBind.Name).
		SubResource("status").
		Body(dockerBind).
		Do().
		Into(result)
	return
}

// Delete takes name of the dockerBind and deletes it. Returns an error if one occurs.
func (c *dockerBinds) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dockerbinds").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dockerBinds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dockerbinds").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dockerBind.
func (c *dockerBinds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DockerBind, err error) {
	result = &v1alpha1.DockerBind{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dockerbinds").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
