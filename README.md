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

1. Install the CRDs:
   ```sh
   helm upgrade --install proompteng-crds charts/proompteng-crds -n proompteng-system --create-namespace
   ```
2. Install the operator:
   ```sh
   helm upgrade --install proompteng charts/proompteng-operator -n proompteng-system
   ```
3. Apply the sample resources:
   ```sh
   kubectl apply -f examples/ns.yaml
   kubectl apply -f examples/memory-mongodb.yaml
   kubectl apply -f examples/agent-echo.yaml
   ```
4. Run the operator locally (optional):
   ```sh
   KUBECONFIG=${KUBECONFIG:-~/.kube/config} go run ./operator/main.go
   ```
5. Build and push the example agent image (optional):
   ```sh
   docker build -t ghcr.io/proompteng/echo-agent:0.1.0 apps/echo-agent
   docker push ghcr.io/proompteng/echo-agent:0.1.0
   ```

## Contributing

Contributions are welcome. Please open an issue with proposed changes before submitting a pull request so we can align roadmap and expectations.

## License

Licensed under the Apache License, Version 2.0. See [LICENSE](LICENSE).
