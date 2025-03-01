# Cloud Sync Operator


Cloud Sync is a Kubernetes operator that implements GitOps principles for continuous deployment and synchronization of Kubernetes resources. Similar to ArgoCD, it uses the gitops-engine to ensure that the desired state in Git repositories is synchronized with the actual state in Kubernetes clusters.

## Overview

Cloud Sync provides:
- Automated synchronization between Git repositories and Kubernetes clusters
- Declarative, version-controlled application deployment
- Multi-cluster management capabilities
- Drift detection and automatic reconciliation
- Support for multiple Git providers (GitHub, GitLab, Bitbucket)

## Features

- **GitOps Engine Integration**: Leverages the powerful gitops-engine for reliable state management
- **Declarative Configuration**: All configurations are declarative and version-controlled
- **Multi-Tenancy**: Supports multiple teams and projects with RBAC integration
- **Health Monitoring**: Real-time monitoring of application and sync status
- **Automated Sync**: Continuous synchronization with configurable intervals
- **Rollback Capabilities**: Easy rollback to previous versions
- **Resource Tracking**: Track and manage Kubernetes resources across clusters

## Architecture

Cloud Sync follows a controller-based architecture:
- **Operator Controller**: Manages the core operator functionality
- **Sync Controller**: Handles Git to cluster synchronization
- **Health Controller**: Monitors application health and status
- **Repository Controller**: Manages Git repository connections

## Prerequisites

- Kubernetes cluster (v1.19+)
- Go 1.19+
- kubectl
- Access to a Git repository

## Installation

```bash
# Clone the repository
git clone https://github.com/Galactic-Grid/cloud-sync.git

# Install CRDs
kubectl apply -f config/crd/bases

# Deploy the operator
kubectl apply -f config/manager/manager.yaml
```

## Quick Start

1. Create a Cloud Sync application:

### YAML-based Application
```yaml
apiVersion: sync.cloudsync.io/v1alpha1
kind: Application
metadata:
  name: example-yaml-app
spec:
  provider:
    repoUrl: https://github.com/example/repo.git
    targetPath: kubernetes/
    targetType: yaml
    targetRevision: HEAD
```

### Script-based Application
```yaml
apiVersion: sync.cloudsync.io/v1alpha1
kind: Application
metadata:
  name: example-script-app
spec:
  provider:
    repoUrl: https://github.com/example/repo.git
    targetPath: scripts/
    targetType: script
    scriptYamlOutput: manifests/output.yaml  # Required when targetType is script
    targetRevision: HEAD
```

2. Apply the configuration:
```bash
kubectl apply -f application.yaml
```

## Development

### Building from Source

```bash
# Build
make build

# Run tests
make test

# Generate manifests
make manifests
```

### Project Structure

```
cloud-sync/
├── api/
│   └── v1alpha1/          # API definitions
├── config/               # Kubernetes manifests
├── controllers/         # Controller implementation
└── Makefile             # Build and deployment targets
```

## Contributing

Contributions are welcome! Please read our [Contributing Guide](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to the ArgoCD team for inspiration
- GitOps Engine contributors
- Kubernetes community
