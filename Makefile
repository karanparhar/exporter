# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=exporter
BINARY_UNIX=$(BINARY_NAME)_unix
VERSION=0.1.0
all: test build
fmt:
		$(GOCMD) fmt ./...
build:

		$(GOBUILD) -o $(BINARY_NAME) -v	
test:
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)
release-image:
		
		eval $(minikube docker-env)
		docker build -t ${BINARY_NAME}:${VERSION} --build-arg RELEASE=${VERSION} .
deploy:
		
		kubectl delete secret configfile
		kubectl create secret generic configfile  --from-file=config.json
		kubectl apply -f deploy/service.yaml
		kubectl apply -f deploy/deployment.yaml


.PHONY: all fmt test build clean release-image deploy
