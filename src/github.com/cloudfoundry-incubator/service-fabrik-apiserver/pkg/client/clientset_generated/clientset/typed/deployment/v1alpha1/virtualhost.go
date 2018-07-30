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

// VirtualhostsGetter has a method to return a VirtualhostInterface.
// A group's client should implement this interface.
type VirtualhostsGetter interface {
	Virtualhosts(namespace string) VirtualhostInterface
}

// VirtualhostInterface has methods to work with Virtualhost resources.
type VirtualhostInterface interface {
	Create(*v1alpha1.Virtualhost) (*v1alpha1.Virtualhost, error)
	Update(*v1alpha1.Virtualhost) (*v1alpha1.Virtualhost, error)
	UpdateStatus(*v1alpha1.Virtualhost) (*v1alpha1.Virtualhost, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Virtualhost, error)
	List(opts v1.ListOptions) (*v1alpha1.VirtualhostList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Virtualhost, err error)
	VirtualhostExpansion
}

// virtualhosts implements VirtualhostInterface
type virtualhosts struct {
	client rest.Interface
	ns     string
}

// newVirtualhosts returns a Virtualhosts
func newVirtualhosts(c *DeploymentV1alpha1Client, namespace string) *virtualhosts {
	return &virtualhosts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualhost, and returns the corresponding virtualhost object, and an error if there is any.
func (c *virtualhosts) Get(name string, options v1.GetOptions) (result *v1alpha1.Virtualhost, err error) {
	result = &v1alpha1.Virtualhost{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualhosts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Virtualhosts that match those selectors.
func (c *virtualhosts) List(opts v1.ListOptions) (result *v1alpha1.VirtualhostList, err error) {
	result = &v1alpha1.VirtualhostList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualhosts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualhosts.
func (c *virtualhosts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualhosts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a virtualhost and creates it.  Returns the server's representation of the virtualhost, and an error, if there is any.
func (c *virtualhosts) Create(virtualhost *v1alpha1.Virtualhost) (result *v1alpha1.Virtualhost, err error) {
	result = &v1alpha1.Virtualhost{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualhosts").
		Body(virtualhost).
		Do().
		Into(result)
	return
}

// Update takes the representation of a virtualhost and updates it. Returns the server's representation of the virtualhost, and an error, if there is any.
func (c *virtualhosts) Update(virtualhost *v1alpha1.Virtualhost) (result *v1alpha1.Virtualhost, err error) {
	result = &v1alpha1.Virtualhost{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualhosts").
		Name(virtualhost.Name).
		Body(virtualhost).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *virtualhosts) UpdateStatus(virtualhost *v1alpha1.Virtualhost) (result *v1alpha1.Virtualhost, err error) {
	result = &v1alpha1.Virtualhost{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualhosts").
		Name(virtualhost.Name).
		SubResource("status").
		Body(virtualhost).
		Do().
		Into(result)
	return
}

// Delete takes name of the virtualhost and deletes it. Returns an error if one occurs.
func (c *virtualhosts) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualhosts").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualhosts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualhosts").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched virtualhost.
func (c *virtualhosts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Virtualhost, err error) {
	result = &v1alpha1.Virtualhost{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualhosts").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
