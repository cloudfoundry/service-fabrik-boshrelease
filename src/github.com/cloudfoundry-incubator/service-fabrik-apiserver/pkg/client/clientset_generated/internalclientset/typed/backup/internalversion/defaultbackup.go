//TODO copyright header
package internalversion

import (
	backup "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/backup"
	scheme "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DefaultBackupsGetter has a method to return a DefaultBackupInterface.
// A group's client should implement this interface.
type DefaultBackupsGetter interface {
	DefaultBackups(namespace string) DefaultBackupInterface
}

// DefaultBackupInterface has methods to work with DefaultBackup resources.
type DefaultBackupInterface interface {
	Create(*backup.DefaultBackup) (*backup.DefaultBackup, error)
	Update(*backup.DefaultBackup) (*backup.DefaultBackup, error)
	UpdateStatus(*backup.DefaultBackup) (*backup.DefaultBackup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*backup.DefaultBackup, error)
	List(opts v1.ListOptions) (*backup.DefaultBackupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *backup.DefaultBackup, err error)
	DefaultBackupExpansion
}

// defaultBackups implements DefaultBackupInterface
type defaultBackups struct {
	client rest.Interface
	ns     string
}

// newDefaultBackups returns a DefaultBackups
func newDefaultBackups(c *BackupClient, namespace string) *defaultBackups {
	return &defaultBackups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the defaultBackup, and returns the corresponding defaultBackup object, and an error if there is any.
func (c *defaultBackups) Get(name string, options v1.GetOptions) (result *backup.DefaultBackup, err error) {
	result = &backup.DefaultBackup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("defaultbackups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DefaultBackups that match those selectors.
func (c *defaultBackups) List(opts v1.ListOptions) (result *backup.DefaultBackupList, err error) {
	result = &backup.DefaultBackupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("defaultbackups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested defaultBackups.
func (c *defaultBackups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("defaultbackups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a defaultBackup and creates it.  Returns the server's representation of the defaultBackup, and an error, if there is any.
func (c *defaultBackups) Create(defaultBackup *backup.DefaultBackup) (result *backup.DefaultBackup, err error) {
	result = &backup.DefaultBackup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("defaultbackups").
		Body(defaultBackup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a defaultBackup and updates it. Returns the server's representation of the defaultBackup, and an error, if there is any.
func (c *defaultBackups) Update(defaultBackup *backup.DefaultBackup) (result *backup.DefaultBackup, err error) {
	result = &backup.DefaultBackup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("defaultbackups").
		Name(defaultBackup.Name).
		Body(defaultBackup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *defaultBackups) UpdateStatus(defaultBackup *backup.DefaultBackup) (result *backup.DefaultBackup, err error) {
	result = &backup.DefaultBackup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("defaultbackups").
		Name(defaultBackup.Name).
		SubResource("status").
		Body(defaultBackup).
		Do().
		Into(result)
	return
}

// Delete takes name of the defaultBackup and deletes it. Returns an error if one occurs.
func (c *defaultBackups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("defaultbackups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *defaultBackups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("defaultbackups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched defaultBackup.
func (c *defaultBackups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *backup.DefaultBackup, err error) {
	result = &backup.DefaultBackup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("defaultbackups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
