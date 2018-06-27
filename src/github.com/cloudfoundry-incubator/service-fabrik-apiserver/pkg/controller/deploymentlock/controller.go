
//TODO copyright header


package deploymentlock

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/lock/v1alpha1"
	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/controller/sharedinformers"
	listers "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/listers_generated/lock/v1alpha1"
)

// +controller:group=lock,version=v1alpha1,kind=DeploymentLock,resource=deploymentlocks
type DeploymentLockControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about DeploymentLock
	lister listers.DeploymentLockLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *DeploymentLockControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing deploymentlocks labels
	c.lister = arguments.GetSharedInformers().Factory.Lock().V1alpha1().DeploymentLocks().Lister()
}

// Reconcile handles enqueued messages
func (c *DeploymentLockControllerImpl) Reconcile(u *v1alpha1.DeploymentLock) error {
	// Implement controller logic here
	log.Printf("Running reconcile DeploymentLock for %s\n", u.Name)
	return nil
}

func (c *DeploymentLockControllerImpl) Get(namespace, name string) (*v1alpha1.DeploymentLock, error) {
	return c.lister.DeploymentLocks(namespace).Get(name)
}
