//TODO copyright header
package fake

import (
	internalversion "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/internalclientset/typed/deployment/internalversion"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeDeployment struct {
	*testing.Fake
}

func (c *FakeDeployment) Directors(namespace string) internalversion.DirectorInterface {
	return &FakeDirectors{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeDeployment) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
