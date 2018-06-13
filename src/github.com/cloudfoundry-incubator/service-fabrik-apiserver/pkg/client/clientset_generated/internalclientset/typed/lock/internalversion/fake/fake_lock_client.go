//TODO copyright header
package fake

import (
	internalversion "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/internalclientset/typed/lock/internalversion"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeLock struct {
	*testing.Fake
}

func (c *FakeLock) DeploymentLocks(namespace string) internalversion.DeploymentLockInterface {
	return &FakeDeploymentLocks{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeLock) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
