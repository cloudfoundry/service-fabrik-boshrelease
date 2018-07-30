//TODO copyright header
package fake

import (
	v1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/clientset/typed/bind/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeBindV1alpha1 struct {
	*testing.Fake
}

func (c *FakeBindV1alpha1) DirectorBinds(namespace string) v1alpha1.DirectorBindInterface {
	return &FakeDirectorBinds{c, namespace}
}

func (c *FakeBindV1alpha1) DockerBinds(namespace string) v1alpha1.DockerBindInterface {
	return &FakeDockerBinds{c, namespace}
}

func (c *FakeBindV1alpha1) Virtualhostbinds(namespace string) v1alpha1.VirtualhostbindInterface {
	return &FakeVirtualhostbinds{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeBindV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
