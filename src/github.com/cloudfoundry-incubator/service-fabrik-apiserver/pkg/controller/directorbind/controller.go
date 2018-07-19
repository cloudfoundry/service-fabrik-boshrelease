
//TODO copyright header


package directorbind

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/bind/v1alpha1"
	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/controller/sharedinformers"
	listers "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/listers_generated/bind/v1alpha1"
)

// +controller:group=bind,version=v1alpha1,kind=DirectorBind,resource=directorbinds
type DirectorBindControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about DirectorBind
	lister listers.DirectorBindLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *DirectorBindControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing directorbinds labels
	c.lister = arguments.GetSharedInformers().Factory.Bind().V1alpha1().DirectorBinds().Lister()
}

// Reconcile handles enqueued messages
func (c *DirectorBindControllerImpl) Reconcile(u *v1alpha1.DirectorBind) error {
	// Implement controller logic here
	log.Printf("Running reconcile DirectorBind for %s\n", u.Name)
	return nil
}

func (c *DirectorBindControllerImpl) Get(namespace, name string) (*v1alpha1.DirectorBind, error) {
	return c.lister.DirectorBinds(namespace).Get(name)
}
