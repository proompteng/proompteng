# Proompteng CRDs Chart

This chart installs the CustomResourceDefinitions required by the Proompteng platform. It must be installed before the operator chart and upgraded first whenever CRD schemas change.

## Usage

```sh
helm upgrade --install proompteng-crds charts/proompteng-crds -n proompteng-system --create-namespace
```

## Upgrading CRDs

1. Bump the chart `version` in `Chart.yaml`.
2. Run `helm dependency update` if needed (none by default).
3. Apply the upgraded CRDs:
   ```sh
   helm upgrade proompteng-crds charts/proompteng-crds -n proompteng-system
   ```
4. Once the CRDs are upgraded, upgrade the operator chart.

## Notes

- CRDs are stored as raw YAML under `crds/` so Helm will apply them with server-side apply semantics.
- Helm hooks are not used; ordering is managed by installing this chart separately.
