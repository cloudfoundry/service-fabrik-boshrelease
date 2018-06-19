
//TODO copyright header


package defaultbackup

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/backup/v1alpha1"
	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/controller/sharedinformers"
	listers "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/client/listers_generated/backup/v1alpha1"
)

// +controller:group=backup,version=v1alpha1,kind=DefaultBackup,resource=defaultbackups
type DefaultBackupControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about DefaultBackup
	lister listers.DefaultBackupLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *DefaultBackupControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing defaultbackups labels
	c.lister = arguments.GetSharedInformers().Factory.Backup().V1alpha1().DefaultBackups().Lister()
}

// Reconcile handles enqueued messages
func (c *DefaultBackupControllerImpl) Reconcile(u *v1alpha1.DefaultBackup) error {
	// Implement controller logic here
	log.Printf("Running reconcile DefaultBackup for %s\n", u.Name)
	return nil
}

func (c *DefaultBackupControllerImpl) Get(namespace, name string) (*v1alpha1.DefaultBackup, error) {
	return c.lister.DefaultBackups(namespace).Get(name)
}
