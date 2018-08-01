//TODO copyright header
package v1alpha1

import (
	v1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/bind/v1alpha1"
	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/clientset/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type BindV1alpha1Interface interface {
	RESTClient() rest.Interface
	DirectorBindsGetter
	DockerBindsGetter
}

// BindV1alpha1Client is used to interact with features provided by the bind.servicefabrik.io group.
type BindV1alpha1Client struct {
	restClient rest.Interface
}

func (c *BindV1alpha1Client) DirectorBinds(namespace string) DirectorBindInterface {
	return newDirectorBinds(c, namespace)
}

func (c *BindV1alpha1Client) DockerBinds(namespace string) DockerBindInterface {
	return newDockerBinds(c, namespace)
}

// NewForConfig creates a new BindV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*BindV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &BindV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new BindV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *BindV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new BindV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *BindV1alpha1Client {
	return &BindV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *BindV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
