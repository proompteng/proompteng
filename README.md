# Proompteng Platform

Proompteng provides a foundation for deploying AI agent infrastructure on Kubernetes. This repository contains:

- Helm charts for installing the Proompteng CRDs and operator
- A Go-based operator built with Kubebuilder
- An example echo agent application using the Proompteng SDK
- Example manifests, workflows, and documentation to help you get running quickly


## Repository Layout

- `charts/proompteng-crds`: Raw CustomResourceDefinition manifests packaged as a Helm chart
- `charts/proompteng-operator`: Helm chart for deploying the operator and related resources
- `operator`: Go operator codebase built with Kubebuilder
- `apps/echo-agent`: Example FastAPI agent demonstrating the runtime SDK
- `examples`: Example Kubernetes manifests for common resources
- `docs`: Design notes and operational guidance
- `.github/workflows/ci.yaml`: Continuous integration pipeline

## Quickstart

1. Install everything with a single manifest (ideal for quick starts or GitOps repos):
   ```sh
   kubectl apply -f https://raw.githubusercontent.com/proompteng/proompteng/main/manifests/proompteng/install.yaml
   ```
   The bundle creates the `proompteng-system` namespace, registers all CustomResourceDefinitions, and deploys the operator with default settings.
   For GitOps/Kustomize setups you can reference the same file directly:
   ```yaml
   resources:
     - https://raw.githubusercontent.com/proompteng/proompteng/<tag>/manifests/proompteng/install.yaml
   ```
   Replace `<tag>` with the Git tag or branch you want to track (for example, `main` or `v0.1.0`).
   When you want to pin to an immutable GitHub release, reference the
   versioned manifest published under `releases/`:
   ```yaml
   resources:
     - https://raw.githubusercontent.com/proompteng/proompteng/<tag>/releases/kubernetes/proompteng/install.yaml
   ```
   The [`examples/gitops/lab`](examples/gitops/lab) directory mirrors the
   structure used in [`gregkonush/lab`](https://github.com/gregkonush/lab) and
   shows how to drop the manifest URL straight into an ArgoCD application.
2. Install the CRDs with Helm (optional when using the single manifest):
   ```sh
   helm upgrade --install proompteng-crds charts/proompteng-crds -n proompteng-system --create-namespace
   ```
3. Install the operator with Helm (optional when using the single manifest):
   ```sh
   helm upgrade --install proompteng charts/proompteng-operator -n proompteng-system
   ```
4. Apply the sample resources:
   ```sh
   kubectl apply -f examples/ns.yaml
   kubectl apply -f examples/memory-mongodb.yaml
   kubectl apply -f examples/agent-echo.yaml
   ```
5. Run the operator locally (optional):
   ```sh
   KUBECONFIG=${KUBECONFIG:-~/.kube/config} go run ./operator/main.go
   ```
6. Build and push the example agent image (optional):
   ```sh
   docker build -t ghcr.io/proompteng/echo-agent:0.1.0 apps/echo-agent
   docker push ghcr.io/proompteng/echo-agent:0.1.0
   ```

### Smoke-test the manifest with Kind

Before pointing a real cluster at the bundled manifest you can validate the
resources on a local [Kind](https://kind.sigs.k8s.io/) cluster. The repository
includes a helper script that provisions a throwaway cluster, applies the
release manifest, and waits for the operator to become ready:

```sh
./hack/test_release_manifest.sh
```

Pass a different manifest path or URL to test prospective changes:

```sh
./hack/test_release_manifest.sh manifests/proompteng/install.yaml
```

Set `KEEP_CLUSTER=1` if you would like to inspect the resources after the
script completes:

```sh
KEEP_CLUSTER=1 ./hack/test_release_manifest.sh
kubectl get pods -n proompteng-system
```

The script deletes the Kind cluster by default so local testing stays fast and
repeatable.

## Contributing

Contributions are welcome. Please open an issue with proposed changes before submitting a pull request so we can align roadmap and expectations.

## License

Licensed under the Apache License, Version 2.0. See [LICENSE](LICENSE).
