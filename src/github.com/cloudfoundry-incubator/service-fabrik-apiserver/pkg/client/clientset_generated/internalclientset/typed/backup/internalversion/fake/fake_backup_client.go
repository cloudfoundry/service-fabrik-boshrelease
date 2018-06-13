//TODO copyright header
package fake

import (
	internalversion "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/internalclientset/typed/backup/internalversion"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeBackup struct {
	*testing.Fake
}

func (c *FakeBackup) DefaultBackups(namespace string) internalversion.DefaultBackupInterface {
	return &FakeDefaultBackups{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeBackup) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
