
# Image URL to use all building/pushing image targets
IMG ?= controller:latest
IMG_WEBHOOKS ?= amshumanraokaraya/manager-webhooks:latest
NAMESPACE ?= default

all: test manager

# Run tests
test: generate fmt vet manifests
	go test ./pkg/... ./cmd/... -coverprofile cover.out

# Build manager binary
manager: generate fmt vet
	go build -o bin/manager github.com/cloudfoundry-incubator/service-fabrik-crds/cmd/manager

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet
	go run ./cmd/manager/main.go

# Build managerp-webhooks binary
manager-webhooks: generate fmt vet
	go build -o bin/manager-webhooks github.com/cloudfoundry-incubator/service-fabrik-crds/cmd/webhooks/manager

# Run against the configured Kubernetes cluster in ~/.kube/config
run-webhooks: generate fmt vet
	go run ./cmd/webhooks/manager/main.go

# Install CRDs into a cluster
install: manifests
	kubectl apply -f config/crds

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy: manifests
	kubectl apply -f config/crds
	kustomize build config/default | kubectl apply -f -

deploy-secrets:
	sh config/default/webhook-create-signed-cert.sh --namespace ${NAMESPACE}
    
deploy-webhooks: docker-push-webhooks deploy-secrets
	kubectl apply -f config/default/manager_webhooks_image_patch.yaml
	cat config/default/mutating-webhooks-config.yaml | NAMESPACE=${NAMESPACE} sh config/default/webhook-patch-ca-bundle.sh | kubectl apply -f -

undeploy-webhooks:
	kubectl delete -f config/default/mutating-webhooks-config.yaml
	kubectl delete -f config/default/manager_webhooks_image_patch.yaml


# Generate manifests e.g. CRD, RBAC etc.
manifests:
	go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go all

# Run go fmt against code
fmt:
	go fmt ./pkg/... ./cmd/...

# Run go vet against code
vet:
	go vet ./pkg/... ./cmd/...

# Generate code
generate:
	go generate ./pkg/... ./cmd/...

# Build the docker image
docker-build: test
	docker build . -t ${IMG}
	@echo "updating kustomize image patch file for manager resource"
	sed -i 's@image: .*@image: '"${IMG}"'@' ./config/default/manager_image_patch.yaml

# Push the docker image
docker-push:
	docker push ${IMG}
# Build the docker image

docker-build-webhooks: manager-webhooks
	docker build . -f build/Dockerfile -t ${IMG_WEBHOOKS}
	@echo "updating kustomize image patch file for manager resource"
	sed -i 's@image: .*@image: '"${IMG_WEBHOOKS}"'@' ./config/default/manager_webhooks_image_patch.yaml

# Push the docker image
docker-push-webhooks: docker-build-webhooks
	docker push ${IMG_WEBHOOKS}
