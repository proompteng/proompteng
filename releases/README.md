# Release Manifests

This directory holds versioned, GitOps-friendly Kubernetes manifests so that other
repositories—such as [`gregkonush/lab`](https://github.com/gregkonush/lab)—can
reference Proompteng from a single Git URL.

## Layout

```
releases/
  kubernetes/
    proompteng/
      crds/                # Raw CustomResourceDefinitions bundled with every release
      install.yaml         # Concatenated manifest for `kubectl apply -f`
      kustomization.yaml   # Allows `kubectl apply -k` / ArgoCD remote bases
      operator-*.yaml      # Operator components used by the bundle
      namespace.yaml       # Namespace definition for `proompteng-system`
```

Consumers can either download `install.yaml` directly, use the mirrored copy at
`manifests/proompteng/install.yaml`, or add the kustomization as an
ArgoCD/Kustomize remote base, e.g.:

```yaml
resources:
  - github.com/proompteng/proompteng//releases/kubernetes/proompteng?ref=v0.1.0
```

If you prefer to mirror the single-file pattern used in
[`gregkonush/lab`](https://github.com/gregkonush/lab), you can reference the
raw manifest directly from GitHub:

```yaml
resources:
  - https://raw.githubusercontent.com/proompteng/proompteng/v0.1.0/releases/kubernetes/proompteng/install.yaml
```

## Updating the Bundle

1. Modify any of the manifests under `releases/kubernetes/proompteng/`.
2. Regenerate the concatenated `install.yaml`:
   ```sh
   make release-manifest
   ```
3. Commit the changes alongside chart or operator updates.
4. Cut a Git tag and GitHub release so downstream consumers can pin a version.

When publishing a new release, update the image tags and chart versions to match
before generating the bundle so the one-step install stays in sync with Helm
packages.

## Smoke Testing with Kind

Use the repository helper script to validate the manifest on a disposable
Kind cluster before tagging a release:

```sh
./hack/test_release_manifest.sh releases/kubernetes/proompteng/install.yaml
```

Pass a remote URL or alternate manifest path to exercise unpublished changes,
and set `KEEP_CLUSTER=1` to keep the Kind cluster alive for further manual
inspection. The script waits for the operator deployment to report as
Available, mirroring the checks our CI pipeline performs.
