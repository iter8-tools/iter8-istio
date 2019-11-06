# Image URL to use all building/pushing image targets
IMG ?= iter8-controller:latest

all: manager

# Build manager binary
manager: generate fmt vet
	go build -o bin/manager github.com/iter8-tools/iter8-controller/cmd/manager

# Run against the Kubernetes cluster configured in $KUBECONFIG or ~/.kube/config
run: generate fmt vet
	go run ./cmd/manager/main.go

# Prepare Kubernetes cluster for iter8 (running in cluster or locally):
#   install CRDs
#   install configmap/iter8-metrics is defined in namespace iter8 (creating namespace if needed)
load: 
	helm template install/helm/iter8-controller \
	  --name iter8-controller \
	  -x templates/default/namespace.yaml \
	  -x templates/crds/iter8_v1alpha1_experiment.yaml \
	  -x templates/metrics/iter8_metrics.yaml \
	| kubectl apply -f -

# Deploy controller to the Kubernetes cluster configured in $KUBECONFIG or ~/.kube/config
deploy: 
	helm template install/helm/iter8-controller \
	  --name iter8-controller \
	  --set image.repository=`echo ${IMG} | cut -f1 -d':'` \
	  --set image.tag=`echo ${IMG} | cut -f2 -d':'` \
	| kubectl apply -f -

# Run go fmt against code
fmt:
	go fmt ./pkg/... ./cmd/...

# Run go vet against code
vet:
	go vet ./pkg/... ./cmd/...

# Generate code
generate:
ifndef GOPATH
	$(error GOPATH not defined, please define GOPATH. Run "go help gopath" to learn more about GOPATH)
endif
	go generate ./pkg/... ./cmd/...

# Build the docker image
docker-build:
	docker build . -t ${IMG}

# Push the docker image
docker-push:
	docker push ${IMG}

build-default:
	helm template install/helm/iter8-controller \
   		--name iter8-controller \
	> install/iter8-controller.yaml