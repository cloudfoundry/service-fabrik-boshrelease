//TODO copyright header
package fake

import (
	v1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/clientset/typed/deployment/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeDeploymentV1alpha1 struct {
	*testing.Fake
}

func (c *FakeDeploymentV1alpha1) Directors(namespace string) v1alpha1.DirectorInterface {
	return &FakeDirectors{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeDeploymentV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
