
//TODO copyright header


package dockerbind

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/bind/v1alpha1"
	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/controller/sharedinformers"
	listers "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/listers_generated/bind/v1alpha1"
)

// +controller:group=bind,version=v1alpha1,kind=DockerBind,resource=dockerbinds
type DockerBindControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about DockerBind
	lister listers.DockerBindLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *DockerBindControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing dockerbinds labels
	c.lister = arguments.GetSharedInformers().Factory.Bind().V1alpha1().DockerBinds().Lister()
}

// Reconcile handles enqueued messages
func (c *DockerBindControllerImpl) Reconcile(u *v1alpha1.DockerBind) error {
	// Implement controller logic here
	log.Printf("Running reconcile DockerBind for %s\n", u.Name)
	return nil
}

func (c *DockerBindControllerImpl) Get(namespace, name string) (*v1alpha1.DockerBind, error) {
	return c.lister.DockerBinds(namespace).Get(name)
}
