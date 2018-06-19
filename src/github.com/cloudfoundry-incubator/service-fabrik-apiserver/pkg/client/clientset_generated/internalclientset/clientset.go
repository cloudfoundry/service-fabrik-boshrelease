//TODO copyright header
package internalclientset

import (
	backupinternalversion "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/internalclientset/typed/backup/internalversion"
	deploymentinternalversion "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/internalclientset/typed/deployment/internalversion"
	lockinternalversion "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/internalclientset/typed/lock/internalversion"
	glog "github.com/golang/glog"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	Backup() backupinternalversion.BackupInterface
	Deployment() deploymentinternalversion.DeploymentInterface
	Lock() lockinternalversion.LockInterface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	backup     *backupinternalversion.BackupClient
	deployment *deploymentinternalversion.DeploymentClient
	lock       *lockinternalversion.LockClient
}

// Backup retrieves the BackupClient
func (c *Clientset) Backup() backupinternalversion.BackupInterface {
	return c.backup
}

// Deployment retrieves the DeploymentClient
func (c *Clientset) Deployment() deploymentinternalversion.DeploymentInterface {
	return c.deployment
}

// Lock retrieves the LockClient
func (c *Clientset) Lock() lockinternalversion.LockInterface {
	return c.lock
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.backup, err = backupinternalversion.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.deployment, err = deploymentinternalversion.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.lock, err = lockinternalversion.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		glog.Errorf("failed to create the DiscoveryClient: %v", err)
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.backup = backupinternalversion.NewForConfigOrDie(c)
	cs.deployment = deploymentinternalversion.NewForConfigOrDie(c)
	cs.lock = lockinternalversion.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.backup = backupinternalversion.New(c)
	cs.deployment = deploymentinternalversion.New(c)
	cs.lock = lockinternalversion.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
