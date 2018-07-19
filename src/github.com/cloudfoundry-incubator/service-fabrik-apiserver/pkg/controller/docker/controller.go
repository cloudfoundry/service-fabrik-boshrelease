
//TODO copyright header


package docker

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/deployment/v1alpha1"
	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/controller/sharedinformers"
	listers "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/listers_generated/deployment/v1alpha1"
)

// +controller:group=deployment,version=v1alpha1,kind=Docker,resource=dockers
type DockerControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about Docker
	lister listers.DockerLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *DockerControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing dockers labels
	c.lister = arguments.GetSharedInformers().Factory.Deployment().V1alpha1().Dockers().Lister()
}

// Reconcile handles enqueued messages
func (c *DockerControllerImpl) Reconcile(u *v1alpha1.Docker) error {
	// Implement controller logic here
	log.Printf("Running reconcile Docker for %s\n", u.Name)
	return nil
}

func (c *DockerControllerImpl) Get(namespace, name string) (*v1alpha1.Docker, error) {
	return c.lister.Dockers(namespace).Get(name)
}
