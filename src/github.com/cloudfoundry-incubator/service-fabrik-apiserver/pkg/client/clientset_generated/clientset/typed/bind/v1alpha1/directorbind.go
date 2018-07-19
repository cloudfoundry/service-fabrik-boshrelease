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

// DirectorBindsGetter has a method to return a DirectorBindInterface.
// A group's client should implement this interface.
type DirectorBindsGetter interface {
	DirectorBinds(namespace string) DirectorBindInterface
}

// DirectorBindInterface has methods to work with DirectorBind resources.
type DirectorBindInterface interface {
	Create(*v1alpha1.DirectorBind) (*v1alpha1.DirectorBind, error)
	Update(*v1alpha1.DirectorBind) (*v1alpha1.DirectorBind, error)
	UpdateStatus(*v1alpha1.DirectorBind) (*v1alpha1.DirectorBind, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DirectorBind, error)
	List(opts v1.ListOptions) (*v1alpha1.DirectorBindList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DirectorBind, err error)
	DirectorBindExpansion
}

// directorBinds implements DirectorBindInterface
type directorBinds struct {
	client rest.Interface
	ns     string
}

// newDirectorBinds returns a DirectorBinds
func newDirectorBinds(c *BindV1alpha1Client, namespace string) *directorBinds {
	return &directorBinds{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the directorBind, and returns the corresponding directorBind object, and an error if there is any.
func (c *directorBinds) Get(name string, options v1.GetOptions) (result *v1alpha1.DirectorBind, err error) {
	result = &v1alpha1.DirectorBind{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("directorbinds").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DirectorBinds that match those selectors.
func (c *directorBinds) List(opts v1.ListOptions) (result *v1alpha1.DirectorBindList, err error) {
	result = &v1alpha1.DirectorBindList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("directorbinds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested directorBinds.
func (c *directorBinds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("directorbinds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a directorBind and creates it.  Returns the server's representation of the directorBind, and an error, if there is any.
func (c *directorBinds) Create(directorBind *v1alpha1.DirectorBind) (result *v1alpha1.DirectorBind, err error) {
	result = &v1alpha1.DirectorBind{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("directorbinds").
		Body(directorBind).
		Do().
		Into(result)
	return
}

// Update takes the representation of a directorBind and updates it. Returns the server's representation of the directorBind, and an error, if there is any.
func (c *directorBinds) Update(directorBind *v1alpha1.DirectorBind) (result *v1alpha1.DirectorBind, err error) {
	result = &v1alpha1.DirectorBind{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("directorbinds").
		Name(directorBind.Name).
		Body(directorBind).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *directorBinds) UpdateStatus(directorBind *v1alpha1.DirectorBind) (result *v1alpha1.DirectorBind, err error) {
	result = &v1alpha1.DirectorBind{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("directorbinds").
		Name(directorBind.Name).
		SubResource("status").
		Body(directorBind).
		Do().
		Into(result)
	return
}

// Delete takes name of the directorBind and deletes it. Returns an error if one occurs.
func (c *directorBinds) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("directorbinds").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *directorBinds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("directorbinds").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched directorBind.
func (c *directorBinds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DirectorBind, err error) {
	result = &v1alpha1.DirectorBind{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("directorbinds").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
