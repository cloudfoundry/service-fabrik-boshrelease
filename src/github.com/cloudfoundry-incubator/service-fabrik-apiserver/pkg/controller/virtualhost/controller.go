
//TODO copyright header


package virtualhost

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/deployment/v1alpha1"
	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/controller/sharedinformers"
	listers "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/listers_generated/deployment/v1alpha1"
)

// +controller:group=deployment,version=v1alpha1,kind=Virtualhost,resource=virtualhosts
type VirtualhostControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about Virtualhost
	lister listers.VirtualhostLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *VirtualhostControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing virtualhosts labels
	c.lister = arguments.GetSharedInformers().Factory.Deployment().V1alpha1().Virtualhosts().Lister()
}

// Reconcile handles enqueued messages
func (c *VirtualhostControllerImpl) Reconcile(u *v1alpha1.Virtualhost) error {
	// Implement controller logic here
	log.Printf("Running reconcile Virtualhost for %s\n", u.Name)
	return nil
}

func (c *VirtualhostControllerImpl) Get(namespace, name string) (*v1alpha1.Virtualhost, error) {
	return c.lister.Virtualhosts(namespace).Get(name)
}
