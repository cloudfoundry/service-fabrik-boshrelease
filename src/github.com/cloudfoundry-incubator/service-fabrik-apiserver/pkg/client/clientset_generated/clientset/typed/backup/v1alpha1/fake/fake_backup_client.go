//TODO copyright header
package fake

import (
	v1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/clientset/typed/backup/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeBackupV1alpha1 struct {
	*testing.Fake
}

func (c *FakeBackupV1alpha1) DefaultBackups(namespace string) v1alpha1.DefaultBackupInterface {
	return &FakeDefaultBackups{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeBackupV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
