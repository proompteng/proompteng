# Lab-style GitOps Example

This example mirrors the structure used in
[`gregkonush/lab`](https://github.com/gregkonush/lab). By pointing a
`kustomization.yaml` at the Proompteng release manifest you can fold the
operator, namespace, and CRDs into an ArgoCD application with a single URL.

```
argocd/
  applications/
    proompteng/
      kustomization.yaml
```

The provided [`argocd/applications/proompteng/kustomization.yaml`](argocd/applications/proompteng/kustomization.yaml)
sets the namespace and references the tagged release manifest hosted in this
repository. Update the URL to match the version you want ArgoCD to track.
