//TODO copyright header
package internalversion

import (
	deployment "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/deployment"
	scheme "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DirectorsGetter has a method to return a DirectorInterface.
// A group's client should implement this interface.
type DirectorsGetter interface {
	Directors(namespace string) DirectorInterface
}

// DirectorInterface has methods to work with Director resources.
type DirectorInterface interface {
	Create(*deployment.Director) (*deployment.Director, error)
	Update(*deployment.Director) (*deployment.Director, error)
	UpdateStatus(*deployment.Director) (*deployment.Director, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*deployment.Director, error)
	List(opts v1.ListOptions) (*deployment.DirectorList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *deployment.Director, err error)
	DirectorExpansion
}

// directors implements DirectorInterface
type directors struct {
	client rest.Interface
	ns     string
}

// newDirectors returns a Directors
func newDirectors(c *DeploymentClient, namespace string) *directors {
	return &directors{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the director, and returns the corresponding director object, and an error if there is any.
func (c *directors) Get(name string, options v1.GetOptions) (result *deployment.Director, err error) {
	result = &deployment.Director{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("directors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Directors that match those selectors.
func (c *directors) List(opts v1.ListOptions) (result *deployment.DirectorList, err error) {
	result = &deployment.DirectorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("directors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested directors.
func (c *directors) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("directors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a director and creates it.  Returns the server's representation of the director, and an error, if there is any.
func (c *directors) Create(director *deployment.Director) (result *deployment.Director, err error) {
	result = &deployment.Director{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("directors").
		Body(director).
		Do().
		Into(result)
	return
}

// Update takes the representation of a director and updates it. Returns the server's representation of the director, and an error, if there is any.
func (c *directors) Update(director *deployment.Director) (result *deployment.Director, err error) {
	result = &deployment.Director{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("directors").
		Name(director.Name).
		Body(director).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *directors) UpdateStatus(director *deployment.Director) (result *deployment.Director, err error) {
	result = &deployment.Director{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("directors").
		Name(director.Name).
		SubResource("status").
		Body(director).
		Do().
		Into(result)
	return
}

// Delete takes name of the director and deletes it. Returns an error if one occurs.
func (c *directors) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("directors").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *directors) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("directors").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched director.
func (c *directors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *deployment.Director, err error) {
	result = &deployment.Director{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("directors").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
