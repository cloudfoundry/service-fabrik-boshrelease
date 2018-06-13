//TODO copyright header
package internalversion

import (
	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/internalclientset/scheme"
	rest "k8s.io/client-go/rest"
)

type BackupInterface interface {
	RESTClient() rest.Interface
	DefaultBackupsGetter
}

// BackupClient is used to interact with features provided by the backup.servicefabrik.io group.
type BackupClient struct {
	restClient rest.Interface
}

func (c *BackupClient) DefaultBackups(namespace string) DefaultBackupInterface {
	return newDefaultBackups(c, namespace)
}

// NewForConfig creates a new BackupClient for the given config.
func NewForConfig(c *rest.Config) (*BackupClient, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &BackupClient{client}, nil
}

// NewForConfigOrDie creates a new BackupClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *BackupClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new BackupClient for the given RESTClient.
func New(c rest.Interface) *BackupClient {
	return &BackupClient{c}
}

func setConfigDefaults(config *rest.Config) error {
	g, err := scheme.Registry.Group("backup.servicefabrik.io")
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
func (c *BackupClient) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
