# Prometheus exporter

This is sample Prometheus exporter

## Quick start

### Prerequisites
- [go](https://golang.org/dl/) version v1.15+
- minikube

### Steps to run

```
$ mkdir $GOPATH/src/github.com/
$ cd $GOPATH/src/github.com/
$ git clone https://github.com/karanparhar/exporter.git
$ cd exporter
$ make release-image
$ make deploy
```

## get metrics
```
$ curl -X GET $(minikube service --url exporter)/metrics

```