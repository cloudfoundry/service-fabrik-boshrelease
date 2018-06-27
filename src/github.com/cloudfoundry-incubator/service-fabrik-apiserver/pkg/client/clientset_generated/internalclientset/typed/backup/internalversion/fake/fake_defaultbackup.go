//TODO copyright header
package fake

import (
	backup "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/backup"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDefaultBackups implements DefaultBackupInterface
type FakeDefaultBackups struct {
	Fake *FakeBackup
	ns   string
}

var defaultbackupsResource = schema.GroupVersionResource{Group: "backup.servicefabrik.io", Version: "", Resource: "defaultbackups"}

var defaultbackupsKind = schema.GroupVersionKind{Group: "backup.servicefabrik.io", Version: "", Kind: "DefaultBackup"}

// Get takes name of the defaultBackup, and returns the corresponding defaultBackup object, and an error if there is any.
func (c *FakeDefaultBackups) Get(name string, options v1.GetOptions) (result *backup.DefaultBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(defaultbackupsResource, c.ns, name), &backup.DefaultBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*backup.DefaultBackup), err
}

// List takes label and field selectors, and returns the list of DefaultBackups that match those selectors.
func (c *FakeDefaultBackups) List(opts v1.ListOptions) (result *backup.DefaultBackupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(defaultbackupsResource, defaultbackupsKind, c.ns, opts), &backup.DefaultBackupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &backup.DefaultBackupList{}
	for _, item := range obj.(*backup.DefaultBackupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested defaultBackups.
func (c *FakeDefaultBackups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(defaultbackupsResource, c.ns, opts))

}

// Create takes the representation of a defaultBackup and creates it.  Returns the server's representation of the defaultBackup, and an error, if there is any.
func (c *FakeDefaultBackups) Create(defaultBackup *backup.DefaultBackup) (result *backup.DefaultBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(defaultbackupsResource, c.ns, defaultBackup), &backup.DefaultBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*backup.DefaultBackup), err
}

// Update takes the representation of a defaultBackup and updates it. Returns the server's representation of the defaultBackup, and an error, if there is any.
func (c *FakeDefaultBackups) Update(defaultBackup *backup.DefaultBackup) (result *backup.DefaultBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(defaultbackupsResource, c.ns, defaultBackup), &backup.DefaultBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*backup.DefaultBackup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDefaultBackups) UpdateStatus(defaultBackup *backup.DefaultBackup) (*backup.DefaultBackup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(defaultbackupsResource, "status", c.ns, defaultBackup), &backup.DefaultBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*backup.DefaultBackup), err
}

// Delete takes name of the defaultBackup and deletes it. Returns an error if one occurs.
func (c *FakeDefaultBackups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(defaultbackupsResource, c.ns, name), &backup.DefaultBackup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDefaultBackups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(defaultbackupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &backup.DefaultBackupList{})
	return err
}

// Patch applies the patch and returns the patched defaultBackup.
func (c *FakeDefaultBackups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *backup.DefaultBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(defaultbackupsResource, c.ns, name, data, subresources...), &backup.DefaultBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*backup.DefaultBackup), err
}
