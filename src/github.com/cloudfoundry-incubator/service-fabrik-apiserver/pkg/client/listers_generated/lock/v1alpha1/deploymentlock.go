//TODO copyright header

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/lock/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DeploymentLockLister helps list DeploymentLocks.
type DeploymentLockLister interface {
	// List lists all DeploymentLocks in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DeploymentLock, err error)
	// DeploymentLocks returns an object that can list and get DeploymentLocks.
	DeploymentLocks(namespace string) DeploymentLockNamespaceLister
	DeploymentLockListerExpansion
}

// deploymentLockLister implements the DeploymentLockLister interface.
type deploymentLockLister struct {
	indexer cache.Indexer
}

// NewDeploymentLockLister returns a new DeploymentLockLister.
func NewDeploymentLockLister(indexer cache.Indexer) DeploymentLockLister {
	return &deploymentLockLister{indexer: indexer}
}

// List lists all DeploymentLocks in the indexer.
func (s *deploymentLockLister) List(selector labels.Selector) (ret []*v1alpha1.DeploymentLock, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DeploymentLock))
	})
	return ret, err
}

// DeploymentLocks returns an object that can list and get DeploymentLocks.
func (s *deploymentLockLister) DeploymentLocks(namespace string) DeploymentLockNamespaceLister {
	return deploymentLockNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DeploymentLockNamespaceLister helps list and get DeploymentLocks.
type DeploymentLockNamespaceLister interface {
	// List lists all DeploymentLocks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DeploymentLock, err error)
	// Get retrieves the DeploymentLock from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DeploymentLock, error)
	DeploymentLockNamespaceListerExpansion
}

// deploymentLockNamespaceLister implements the DeploymentLockNamespaceLister
// interface.
type deploymentLockNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DeploymentLocks in the indexer for a given namespace.
func (s deploymentLockNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DeploymentLock, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DeploymentLock))
	})
	return ret, err
}

// Get retrieves the DeploymentLock from the indexer for a given namespace and name.
func (s deploymentLockNamespaceLister) Get(name string) (*v1alpha1.DeploymentLock, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("deploymentlock"), name)
	}
	return obj.(*v1alpha1.DeploymentLock), nil
}
