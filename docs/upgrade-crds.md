# Upgrading Proompteng CRDs

CustomResourceDefinitions must be upgraded before any controller relying on their schemas. Follow this procedure whenever CRD manifests change.

1. **Review breaking changes**
   - Document new fields or removals in the changelog.
   - Confirm that conversion webhooks are not required for the change scope.

2. **Update the CRD chart version**
   - Bump `version` in `charts/proompteng-crds/Chart.yaml`.
   - Commit the CRD manifest updates (raw YAML in `charts/proompteng-crds/crds/`).

3. **Apply CRDs first**
   ```sh
   helm upgrade --install proompteng-crds charts/proompteng-crds -n proompteng-system --create-namespace
   ```
   Helm applies CRDs before template rendering, ensuring API availability for subsequent releases.

4. **Upgrade workloads**
   - After CRDs complete, upgrade the operator chart:
     ```sh
     helm upgrade --install proompteng charts/proompteng-operator -n proompteng-system
     ```
   - Verify `kubectl get crd` reflects the expected `AGE`/`VERSION` transitions.

5. **Validate compatibility**
   - Reconcile a representative Agent to ensure status updates succeed.
   - Run `kubeconform` or CI-equivalent validation against sample manifests.

6. **Roll back**
   - If problems arise, roll back the operator first (`helm rollback proompteng`).
   - Restore prior CRDs (`helm rollback proompteng-crds`); be mindful that downgrading CRDs may require manual data cleanup.

> **Tip:** Keep CRD changes backwards compatible when possible to avoid disruptive upgrades.
