//TODO copyright header
package clientset

import (
	backupv1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/clientset/typed/backup/v1alpha1"
	bindv1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/clientset/typed/bind/v1alpha1"
	deploymentv1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/clientset/typed/deployment/v1alpha1"
	lockv1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/clientset_generated/clientset/typed/lock/v1alpha1"
	glog "github.com/golang/glog"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	BackupV1alpha1() backupv1alpha1.BackupV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Backup() backupv1alpha1.BackupV1alpha1Interface
	BindV1alpha1() bindv1alpha1.BindV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Bind() bindv1alpha1.BindV1alpha1Interface
	DeploymentV1alpha1() deploymentv1alpha1.DeploymentV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Deployment() deploymentv1alpha1.DeploymentV1alpha1Interface
	LockV1alpha1() lockv1alpha1.LockV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Lock() lockv1alpha1.LockV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	backupV1alpha1     *backupv1alpha1.BackupV1alpha1Client
	bindV1alpha1       *bindv1alpha1.BindV1alpha1Client
	deploymentV1alpha1 *deploymentv1alpha1.DeploymentV1alpha1Client
	lockV1alpha1       *lockv1alpha1.LockV1alpha1Client
}

// BackupV1alpha1 retrieves the BackupV1alpha1Client
func (c *Clientset) BackupV1alpha1() backupv1alpha1.BackupV1alpha1Interface {
	return c.backupV1alpha1
}

// Deprecated: Backup retrieves the default version of BackupClient.
// Please explicitly pick a version.
func (c *Clientset) Backup() backupv1alpha1.BackupV1alpha1Interface {
	return c.backupV1alpha1
}

// BindV1alpha1 retrieves the BindV1alpha1Client
func (c *Clientset) BindV1alpha1() bindv1alpha1.BindV1alpha1Interface {
	return c.bindV1alpha1
}

// Deprecated: Bind retrieves the default version of BindClient.
// Please explicitly pick a version.
func (c *Clientset) Bind() bindv1alpha1.BindV1alpha1Interface {
	return c.bindV1alpha1
}

// DeploymentV1alpha1 retrieves the DeploymentV1alpha1Client
func (c *Clientset) DeploymentV1alpha1() deploymentv1alpha1.DeploymentV1alpha1Interface {
	return c.deploymentV1alpha1
}

// Deprecated: Deployment retrieves the default version of DeploymentClient.
// Please explicitly pick a version.
func (c *Clientset) Deployment() deploymentv1alpha1.DeploymentV1alpha1Interface {
	return c.deploymentV1alpha1
}

// LockV1alpha1 retrieves the LockV1alpha1Client
func (c *Clientset) LockV1alpha1() lockv1alpha1.LockV1alpha1Interface {
	return c.lockV1alpha1
}

// Deprecated: Lock retrieves the default version of LockClient.
// Please explicitly pick a version.
func (c *Clientset) Lock() lockv1alpha1.LockV1alpha1Interface {
	return c.lockV1alpha1
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
	cs.backupV1alpha1, err = backupv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.bindV1alpha1, err = bindv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.deploymentV1alpha1, err = deploymentv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.lockV1alpha1, err = lockv1alpha1.NewForConfig(&configShallowCopy)
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
	cs.backupV1alpha1 = backupv1alpha1.NewForConfigOrDie(c)
	cs.bindV1alpha1 = bindv1alpha1.NewForConfigOrDie(c)
	cs.deploymentV1alpha1 = deploymentv1alpha1.NewForConfigOrDie(c)
	cs.lockV1alpha1 = lockv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.backupV1alpha1 = backupv1alpha1.New(c)
	cs.bindV1alpha1 = bindv1alpha1.New(c)
	cs.deploymentV1alpha1 = deploymentv1alpha1.New(c)
	cs.lockV1alpha1 = lockv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
