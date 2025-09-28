# proompteng

> ship ai agents with full control plane automation.

[![CI status](https://github.com/proompteng/proompteng/actions/workflows/ci.yaml/badge.svg)](https://github.com/proompteng/proompteng/actions/workflows/ci.yaml)
[![Latest release](https://img.shields.io/github/v/release/proompteng/proompteng?display_name=tag&logo=github)](https://github.com/proompteng/proompteng/releases)
[![License: Apache 2.0](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![GitHub stars](https://img.shields.io/github/stars/proompteng/proompteng?style=social)](https://github.com/proompteng/proompteng/stargazers)

## At a glance

- **Declarative agent control plane**: Define agent runtimes, memory backends, and policies with Kubernetes CustomResourceDefinitions.
- **Batteries included**: Operator, Helm charts, and sample agents help teams bootstrap in minutes.
- **Enterprise-ready foundations**: Pluggable observability, GitOps-friendly packaging, and reproducible releases.

## Why proompteng

Modern teams need to launch AI agents that are observable, repeatable, and safe. proompteng layers an operator-driven control plane onto Kubernetes so you can:

- Treat agents as first-class resources with desired-state reconciliation.
- Keep infrastructure GitOps-ready through Helm charts and apply-once manifests.
- Standardise on a Python SDK and Go operator primitives that scale from prototypes to production clusters.

## Feature highlights

- **CRD suite for agents, memories, and policies** with schema validation.
- **Kubebuilder-based operator** that reconciles workloads, secrets, and service accounts.
- **FastAPI echo-agent reference** showcasing the runtime SDK, memory adapters, and observability hooks.
- **Smoke-test examples** under `examples/` for GitOps and direct cluster usage.
- **Release automation** with GitHub Actions CI, versioned manifests, and quick rollback paths.

## Quickstart (5 minutes)

### Prerequisites

- Kubernetes cluster with `kubectl` context set (Kind, k3d, or managed).
- Helm 3.16.2+
- Docker Engine 28.4.0+ (optional if you build custom agents)
- Go 1.25.1+ for local operator runs

### 1. Install the control plane

Apply the bundled manifest for an end-to-end installation:

```sh
kubectl apply -f https://raw.githubusercontent.com/proompteng/proompteng/main/manifests/proompteng/install.yaml
```

This provisions the `proompteng-system` namespace, registers CustomResourceDefinitions, and deploys the operator with sensible defaults.

Prefer GitOps? Reference the same manifest inside your Kustomize base:

```yaml
resources:
  - https://raw.githubusercontent.com/proompteng/proompteng/<tag>/manifests/proompteng/install.yaml
```

Replace `<tag>` with `main` for live development or a release tag such as `v0.1.0`.

### 2. Install via Helm (optional)

Split the install into chart components when you need more control:

```sh
helm upgrade --install proompteng-crds charts/proompteng-crds -n proompteng-system --create-namespace
helm upgrade --install proompteng charts/proompteng-operator -n proompteng-system
```

### 3. Deploy the sample echo agent

```sh
kubectl apply -f examples/ns.yaml
kubectl apply -f examples/memory-mongodb.yaml
kubectl apply -f examples/agent-echo.yaml
```

### 4. Verify the rollout

```sh
kubectl get pods -n proompteng-system
kubectl get agents.proompteng.ai -A
```

You should see the operator ready and the echo agent running under your target namespace.

### 5. Iterate locally (optional)

Run the operator against your current `$KUBECONFIG` for live debugging:

```sh
make run-operator-local
```

## Architecture

proompteng pairs CustomResourceDefinitions with a Kubebuilder operator. Developers declare agent intents via GitOps or CLI, the Kubernetes API stores desired state, and the operator reconciles workloads, secrets, and dependencies onto the cluster. Observability feedback loops flow back through CRD status and logging hooks.

## Platform workflow

1. Model agents, memories, or connectors as CRDs checked into Git.
2. Ship manifests through GitOps or CI pipelines.
3. Allow the operator to reconcile workloads, inject secrets, and emit status.
4. Inspect health via `kubectl`, dashboards, or the SDK’s telemetry endpoints.

## Roadmap signals

- Graph-based plan orchestration across multiple agents.
- Pluggable memory backends beyond MongoDB.
- Managed control plane installation for hosted clusters.
- Observability quickstarts for OpenTelemetry and Grafana Agent.

Track detailed milestones in [Linear](https://linear.app/proompteng).

## Resources

- [Operator source](operator/) — Kubebuilder reconcilers and controller runtime.
- [Helm charts](charts/) — Install CRDs and the control plane with overrides.
- [FastAPI echo agent](apps/echo-agent/) — Example runtime with SDK usage and tests.
- [Examples](examples/) — Apply-once manifests for common deployments.
- [Docs](docs/) — Architecture notes, onboarding, and upgrade guides.

## Community & support

- Follow release notes in the GitHub Releases feed.
- File issues with reproducible steps or enhancement ideas — templates live in `.github/`.
- Report security concerns privately to <security@proompteng.ai>.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for branching conventions, toolchain requirements, and the pull request checklist.

## License

Distributed under the [Apache License 2.0](LICENSE). Commercial support inquiries can reach <partnerships@proompteng.ai>.

## Give proompteng a star

If this project streamlines your agent infrastructure, please star and watch the repository to keep updates flowing.
