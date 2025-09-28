# Kubernetes Manifests

The `manifests/` tree mirrors selected release assets so GitOps tools can fetch
single files straight from GitHub. The primary entry point is
`manifests/proompteng/install.yaml`, which bundles the namespace, CRDs, and
operator components for a `kubectl apply` or remote `kustomization.yaml`
reference:

```yaml
resources:
  - https://raw.githubusercontent.com/proompteng/proompteng/<tag>/manifests/proompteng/install.yaml
```

Run `make release-manifest` after changing any manifests in
`releases/kubernetes/proompteng/` to refresh the mirrored file. Remote
consumers who prefer `apply -k` can point at the accompanying
[`manifests/proompteng/kustomization.yaml`](proompteng/kustomization.yaml)
instead of the raw manifest:

```yaml
resources:
  - github.com/proompteng/proompteng//manifests/proompteng?ref=<tag>
```

Downstream repositories that prefer to reference only tagged releases can use
the canonical copy under `releases/` instead:

```yaml
resources:
  - https://raw.githubusercontent.com/proompteng/proompteng/<tag>/releases/kubernetes/proompteng/install.yaml
```
