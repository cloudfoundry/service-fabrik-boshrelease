//TODO copyright header
package internalversion

import (
	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/internalclientset/scheme"
	rest "k8s.io/client-go/rest"
)

type BindInterface interface {
	RESTClient() rest.Interface
	DirectorBindsGetter
	DockerBindsGetter
	VirtualhostbindsGetter
}

// BindClient is used to interact with features provided by the bind.servicefabrik.io group.
type BindClient struct {
	restClient rest.Interface
}

func (c *BindClient) DirectorBinds(namespace string) DirectorBindInterface {
	return newDirectorBinds(c, namespace)
}

func (c *BindClient) DockerBinds(namespace string) DockerBindInterface {
	return newDockerBinds(c, namespace)
}

func (c *BindClient) Virtualhostbinds(namespace string) VirtualhostbindInterface {
	return newVirtualhostbinds(c, namespace)
}

// NewForConfig creates a new BindClient for the given config.
func NewForConfig(c *rest.Config) (*BindClient, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &BindClient{client}, nil
}

// NewForConfigOrDie creates a new BindClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *BindClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new BindClient for the given RESTClient.
func New(c rest.Interface) *BindClient {
	return &BindClient{c}
}

func setConfigDefaults(config *rest.Config) error {
	g, err := scheme.Registry.Group("bind.servicefabrik.io")
	if err != nil {
		return err
	}

	config.APIPath = "/apis"
	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
	if config.GroupVersion == nil || config.GroupVersion.Group != g.GroupVersion.Group {
		gv := g.GroupVersion
		config.GroupVersion = &gv
	}
	config.NegotiatedSerializer = scheme.Codecs

	if config.QPS == 0 {
		config.QPS = 5
	}
	if config.Burst == 0 {
		config.Burst = 10
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *BindClient) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
