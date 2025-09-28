# Proompteng Onboarding Guide

## Platform Overview
Proompteng is a Kubernetes-native platform for deploying AI agents. The repository bundles Helm charts for the control plane, a Go-based operator that reconciles custom resources, and a sample FastAPI agent showcasing the runtime SDK. Together they deliver end-to-end automation for provisioning agent workloads, wiring dependencies, and demonstrating runtime interactions.

## Repository Map
- **`charts/`** – Helm charts for shipping the CRDs and operator deployments.
- **`operator/`** – Kubebuilder-based controller manager that registers CRD schemes, configures health/metrics endpoints, and wires reconcilers for Agents and Memories.
- **`apps/echo-agent/`** – Reference FastAPI agent and lightweight Python SDK that exercises runtime loading, request handling, and health probes.
- **`examples/`** – Apply-once Kubernetes manifests for namespaces, memory backends, agent instances, and routing policies.
- **`docs/`** – Additional design notes and upgrade runbooks.
- **`Makefile`** – Aggregates linting, testing, Helm validation, and install helpers used in CI and local workflows.

## Key Components
### Operator
- Registers the Proompteng CRDs with the controller-runtime scheme and configures logging from environment variables or a mounted config file.
- Starts a manager with metrics (`:9090`) and health/ready probes (`:8081`), then installs dedicated reconcilers for `Agent` and `Memory` resources and exposes Kubernetes-native health checks.
- `Agent` resources declare display metadata, model identity, container runtime hints, optional transport/service wiring, and references to Memory or Tool resources, while status captures reconciliation phase, service name, and ready replica counts.

### Helm Charts
- `charts/proompteng-crds` packages the raw CustomResourceDefinition manifests so platform operators can install/update the API surface independently of workloads.
- `charts/proompteng-operator` deploys the controller manager image, exposes chart metadata, and is versioned alongside the operator release cadence.

### Example Agent & SDK
- `apps/echo-agent/app.py` exposes `/invoke`, `/healthz`, and `/readyz` endpoints. The `EchoAgent` echoes payloads, annotates them with environment-derived model metadata, and leverages a cached runtime loader to support hot-swapping agent implementations.
- The bundled `sdk/runtime.py` defines a simple `AgentRuntime` protocol and import-based loader that expects an `AGENT` instance exporting a `handle(input_payload, context)` method.

### Examples & Documentation
- `examples/` manifests demonstrate how to provision namespaces, configure a MongoDB-backed `Memory`, deploy the echo agent, register a default route, and apply baseline policy documents.
- The `docs/` folder includes deep dives into CRD design principles plus step-by-step guidance for upgrading CRD charts safely.

## Development Workflow
- Toolchain targets: Go 1.25.1, Python 3.13.7, Node.js 24.6.1, Docker 28.4.0, Helm 3.16.2, kubectl 1.34.1, and GitHub CLI 2.80.0.
- Primary automation entrypoints live in the `Makefile`: `make lint` chains Helm lint, yamllint, kubeconform, Ruff, and Go unit tests; individual phony targets exist for isolated runs. Additional helpers install CRDs, install the operator chart, or run the controller locally against your current kubeconfig.

## What to Explore Next
1. **CRD schemas** – Read `docs/designing-crds.md` and the Go type definitions under `operator/api/v1alpha1/` to understand how Agents, Memories, and future Tool/Route resources are modeled.
2. **Reconciliation logic** – Dive into `operator/controllers/*` to see how Agent and Memory reconcilers map CRDs to Kubernetes Deployments/Services and manage status updates.
3. **Runtime extensibility** – Experiment with `apps/echo-agent/sdk/` by swapping out the `AGENT` implementation or wiring a real memory backend using the example manifests.
4. **Release process** – Follow `docs/upgrade-crds.md` and the Helm chart versions to practice safe CRD upgrades and operator rollouts.
