//TODO copyright header
package fake

import (
	internalversion "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/internalclientset/typed/bind/internalversion"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeBind struct {
	*testing.Fake
}

func (c *FakeBind) DirectorBinds(namespace string) internalversion.DirectorBindInterface {
	return &FakeDirectorBinds{c, namespace}
}

func (c *FakeBind) DockerBinds(namespace string) internalversion.DockerBindInterface {
	return &FakeDockerBinds{c, namespace}
}

func (c *FakeBind) Virtualhostbinds(namespace string) internalversion.VirtualhostbindInterface {
	return &FakeVirtualhostbinds{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeBind) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
