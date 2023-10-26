# Audit-Logging on Kubernetes with Webhook PoC

## Overview

- audit-server inspired by https://dev.bitolog.com/implement-audits-webhook/
- use a fix `ClusterIP` for the kubernetes service of the audit webhook
- using [googles audit-policy](https://github.com/kubernetes/kubernetes/blob/master/cluster/gce/gci/configure-helper.sh#L1101) -> change to only `Request` instead of `RequestResponse`

## Setup kind

```bash
kind create cluster --config kind.yaml
```

## Start audit-webhook

```bash
# build docker image
docker build -t audit-webhook:0.0.1 .

# do not use "latest" -> won't work in kind if you import a local image
# see https://iximiuz.com/en/posts/kubernetes-kind-load-docker-image/
kind load docker-image audit-webhook:0.0.1

# apply deployment & service
kubectl apply -f kubernetes-manifests.yaml

# follow logs
kubectl logs -l app.kubernetes.io/name=audit-webhook -f
```