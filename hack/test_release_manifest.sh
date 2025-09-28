#!/usr/bin/env bash
set -euo pipefail

CLUSTER_NAME="proompteng-smoke"
KEEP_CLUSTER=false
MANIFEST_PATH="${1:-releases/kubernetes/proompteng/install.yaml}"

usage() {
  cat <<USAGE
Usage: ${0##*/} [manifest-path]

Creates a disposable kind cluster, applies the Proompteng release manifest, and waits
for the operator to become ready. Pass a custom manifest path or URL as the first
argument (defaults to releases/kubernetes/proompteng/install.yaml).

Environment variables:
  KIND_CLUSTER_NAME   Override the cluster name (default: proompteng-smoke)
  KEEP_CLUSTER        Set to 1 to skip deleting the cluster when the script exits

Requires kind >= 0.20.0 and kubectl in $PATH.
USAGE
}

if [[ "${1:-}" == "-h" || "${1:-}" == "--help" ]]; then
  usage
  exit 0
fi

if [[ -n "${KIND_CLUSTER_NAME:-}" ]]; then
  CLUSTER_NAME="${KIND_CLUSTER_NAME}"
fi

if [[ "${KEEP_CLUSTER:-0}" == "1" ]]; then
  KEEP_CLUSTER=true
fi

cleanup() {
  if ! $KEEP_CLUSTER; then
    echo "\nDeleting kind cluster \"${CLUSTER_NAME}\"..."
    kind delete cluster --name "${CLUSTER_NAME}" >/dev/null 2>&1 || true
  fi
}

trap cleanup EXIT

if ! command -v kind >/dev/null; then
  echo "kind is required but was not found in PATH" >&2
  exit 1
fi

if ! command -v kubectl >/dev/null; then
  echo "kubectl is required but was not found in PATH" >&2
  exit 1
fi

echo "Creating kind cluster \"${CLUSTER_NAME}\"..."
kind create cluster --name "${CLUSTER_NAME}" >/dev/null

echo "Applying Proompteng manifest: ${MANIFEST_PATH}"
kubectl apply -f "${MANIFEST_PATH}"

echo "Waiting for operator deployment to become available..."
kubectl -n proompteng-system wait \
  --for=condition=Available deployment/proompteng-operator \
  --timeout=180s

echo "\nProompteng manifest applied successfully. Check the namespace status with:\n  kubectl get all -n proompteng-system"

if $KEEP_CLUSTER; then
  echo "KEEP_CLUSTER=1, leaving cluster \"${CLUSTER_NAME}\" running"
fi
