# skaffold example
This example uses nats-streaming and go modules.

# setup
    go mod vendor
    skaffold build | run | dev

# Details
1. Dockerfile creates a static binary
1. k8s-pod.yaml is a kubernetes pod manifes
1. skaffold.yaml uses the above and injects the image name
