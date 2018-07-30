
//TODO copyright header


package virtualhostbind

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/bind/v1alpha1"
	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/controller/sharedinformers"
	listers "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/listers_generated/bind/v1alpha1"
)

// +controller:group=bind,version=v1alpha1,kind=Virtualhostbind,resource=virtualhostbinds
type VirtualhostbindControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about Virtualhostbind
	lister listers.VirtualhostbindLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *VirtualhostbindControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing virtualhostbinds labels
	c.lister = arguments.GetSharedInformers().Factory.Bind().V1alpha1().Virtualhostbinds().Lister()
}

// Reconcile handles enqueued messages
func (c *VirtualhostbindControllerImpl) Reconcile(u *v1alpha1.Virtualhostbind) error {
	// Implement controller logic here
	log.Printf("Running reconcile Virtualhostbind for %s\n", u.Name)
	return nil
}

func (c *VirtualhostbindControllerImpl) Get(namespace, name string) (*v1alpha1.Virtualhostbind, error) {
	return c.lister.Virtualhostbinds(namespace).Get(name)
}
