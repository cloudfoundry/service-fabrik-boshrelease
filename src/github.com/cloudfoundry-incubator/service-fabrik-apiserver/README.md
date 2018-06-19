# service-farbrik-apiserver
Kubernetes-style apiserver for Service Fabrik 2.0

## Getting Started

### From source code

#### Dependencies
* Kubernetes [apiserver-builder](https://github.com/kubernetes-incubator/apiserver-builder). Please follow the [installation instructions](https://github.com/kubernetes-incubator/apiserver-builder/blob/master/docs/installing.md).

#### Get source code
* Clone this project under $GOPATH/src/github.com/cloudfoundry-incubator

#### Build
```
$ make build
```

#### Run locally
```
$ make run-local    # runs apiserver as well as etcd locally
```
Or
```
apiserver-boot run local    # for more options for finer grained control
```

#### Test locally
```
$ curl https://localhost:9443/swagger.json
$ curl https://localhost:9443/apis/deployment.servicefabrik.io/v1alpha1/directors
$ curl https://localhost:9443/apis/backup.servicefabrik.io/v1alpha1/boshbackuprestores
```
