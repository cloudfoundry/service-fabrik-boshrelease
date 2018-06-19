
//TODO copyright header


package director

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/deployment/v1alpha1"
	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/controller/sharedinformers"
	listers "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/listers_generated/deployment/v1alpha1"
)

// +controller:group=deployment,version=v1alpha1,kind=Director,resource=directors
type DirectorControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about Director
	lister listers.DirectorLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *DirectorControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing directors labels
	c.lister = arguments.GetSharedInformers().Factory.Deployment().V1alpha1().Directors().Lister()
}

// Reconcile handles enqueued messages
func (c *DirectorControllerImpl) Reconcile(u *v1alpha1.Director) error {
	// Implement controller logic here
	log.Printf("Running reconcile Director for %s\n", u.Name)
	return nil
}

func (c *DirectorControllerImpl) Get(namespace, name string) (*v1alpha1.Director, error) {
	return c.lister.Directors(namespace).Get(name)
}
