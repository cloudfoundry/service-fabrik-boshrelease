//TODO copyright header
package internalversion

import (
	lock "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/lock"
	scheme "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DeploymentLocksGetter has a method to return a DeploymentLockInterface.
// A group's client should implement this interface.
type DeploymentLocksGetter interface {
	DeploymentLocks(namespace string) DeploymentLockInterface
}

// DeploymentLockInterface has methods to work with DeploymentLock resources.
type DeploymentLockInterface interface {
	Create(*lock.DeploymentLock) (*lock.DeploymentLock, error)
	Update(*lock.DeploymentLock) (*lock.DeploymentLock, error)
	UpdateStatus(*lock.DeploymentLock) (*lock.DeploymentLock, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*lock.DeploymentLock, error)
	List(opts v1.ListOptions) (*lock.DeploymentLockList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *lock.DeploymentLock, err error)
	DeploymentLockExpansion
}

// deploymentLocks implements DeploymentLockInterface
type deploymentLocks struct {
	client rest.Interface
	ns     string
}

// newDeploymentLocks returns a DeploymentLocks
func newDeploymentLocks(c *LockClient, namespace string) *deploymentLocks {
	return &deploymentLocks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the deploymentLock, and returns the corresponding deploymentLock object, and an error if there is any.
func (c *deploymentLocks) Get(name string, options v1.GetOptions) (result *lock.DeploymentLock, err error) {
	result = &lock.DeploymentLock{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deploymentlocks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DeploymentLocks that match those selectors.
func (c *deploymentLocks) List(opts v1.ListOptions) (result *lock.DeploymentLockList, err error) {
	result = &lock.DeploymentLockList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deploymentlocks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested deploymentLocks.
func (c *deploymentLocks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("deploymentlocks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a deploymentLock and creates it.  Returns the server's representation of the deploymentLock, and an error, if there is any.
func (c *deploymentLocks) Create(deploymentLock *lock.DeploymentLock) (result *lock.DeploymentLock, err error) {
	result = &lock.DeploymentLock{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("deploymentlocks").
		Body(deploymentLock).
		Do().
		Into(result)
	return
}

// Update takes the representation of a deploymentLock and updates it. Returns the server's representation of the deploymentLock, and an error, if there is any.
func (c *deploymentLocks) Update(deploymentLock *lock.DeploymentLock) (result *lock.DeploymentLock, err error) {
	result = &lock.DeploymentLock{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deploymentlocks").
		Name(deploymentLock.Name).
		Body(deploymentLock).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *deploymentLocks) UpdateStatus(deploymentLock *lock.DeploymentLock) (result *lock.DeploymentLock, err error) {
	result = &lock.DeploymentLock{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deploymentlocks").
		Name(deploymentLock.Name).
		SubResource("status").
		Body(deploymentLock).
		Do().
		Into(result)
	return
}

// Delete takes name of the deploymentLock and deletes it. Returns an error if one occurs.
func (c *deploymentLocks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deploymentlocks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *deploymentLocks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deploymentlocks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched deploymentLock.
func (c *deploymentLocks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *lock.DeploymentLock, err error) {
	result = &lock.DeploymentLock{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("deploymentlocks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
