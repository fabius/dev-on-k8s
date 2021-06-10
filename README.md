# Kubernetes dev workflow

Showcasing the inner & the outer dev loop when developing on Kubernetes as
described on
[Medium](https://fabiandeifuss.medium.com/developing-on-kubernetes-the-inner-outer-loop-6957f9597f7a).

1. the inner loop: _local development_
2. the outer loop: _deployment strategies & configuration management_

## The inner dev loop

Building, deploying and watching for filechanges is as simple as running

```sh
# use the base configuration
skaffold dev

# or use the local patch
skaffold dev -p local
```

## The outer loop

Similarly, building and deploying the service to your K8s instance does not have
to be any more complicated

```sh
# use the base configuration
skaffold run

# or use the local patch
skaffold run -p local
```
