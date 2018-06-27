//TODO copyright header
package fake

import (
	v1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/clientset/typed/lock/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeLockV1alpha1 struct {
	*testing.Fake
}

func (c *FakeLockV1alpha1) DeploymentLocks(namespace string) v1alpha1.DeploymentLockInterface {
	return &FakeDeploymentLocks{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeLockV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
