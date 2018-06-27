//TODO copyright header
package internalversion

import (
	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/internalclientset/scheme"
	rest "k8s.io/client-go/rest"
)

type LockInterface interface {
	RESTClient() rest.Interface
	DeploymentLocksGetter
}

// LockClient is used to interact with features provided by the lock.servicefabrik.io group.
type LockClient struct {
	restClient rest.Interface
}

func (c *LockClient) DeploymentLocks(namespace string) DeploymentLockInterface {
	return newDeploymentLocks(c, namespace)
}

// NewForConfig creates a new LockClient for the given config.
func NewForConfig(c *rest.Config) (*LockClient, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &LockClient{client}, nil
}

// NewForConfigOrDie creates a new LockClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *LockClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new LockClient for the given RESTClient.
func New(c rest.Interface) *LockClient {
	return &LockClient{c}
}

func setConfigDefaults(config *rest.Config) error {
	g, err := scheme.Registry.Group("lock.servicefabrik.io")
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
func (c *LockClient) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
