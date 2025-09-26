# Proompteng Operator Chart

This chart deploys the Proompteng operator, a Go controller built with Kubebuilder that reconciles Agents, Memories, and other platform resources.

## Prerequisites

- Kubernetes v1.28+
- Proompteng CRDs installed (`charts/proompteng-crds`)
- Optional: Prometheus Operator for ServiceMonitor support

## Installing

```sh
helm upgrade --install proompteng charts/proompteng-operator -n proompteng-system
```

## Configuration Highlights

| Parameter | Description | Default |
|-----------|-------------|---------|
| `image.repository` | Operator image repository | `ghcr.io/proompteng/operator` |
| `image.tag` | Operator image tag | `0.1.0` |
| `replicaCount` | Number of operator pods | `1` |
| `serviceAccount.create` | Create a ServiceAccount | `true` |
| `rbac.create` | Create ClusterRole/Binding | `true` |
| `metrics.service.port` | Metrics service port | `9090` |
| `metrics.serviceMonitor.enabled` | Emit a ServiceMonitor | `false` |
| `featureFlags.enableAgentServices` | Enable AgentService reconciliation | `true` |

See `values.yaml` and `values.schema.json` for the full configuration surface.

## Upgrade Notes

- Upgrade the CRD chart before upgrading this chart when schemas change.
- Changes to feature flags are rolled out via the operator ConfigMap.

## Observability

- Metrics endpoint exposed on `/metrics` at port `9090`.
- Health checks on `/healthz` and `/readyz`.

## Uninstalling

```sh
helm uninstall proompteng -n proompteng-system
```

Removing the operator does not delete any managed resources.
