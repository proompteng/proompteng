.PHONY: lint helm-lint yamllint kubeconform ruff go-test install-crds install-operator run-operator-local release-manifest

PYTHON ?= python3

helm-lint:
	helm lint charts/proompteng-crds
	helm lint charts/proompteng-operator

yamllint:
	yamllint charts examples docs .github/workflows/ci.yaml

kubeconform:
	helm template proompteng charts/proompteng-operator --namespace proompteng-system | kubeconform -strict -ignore-missing-schemas -exit-on-error -
	kubeconform -strict -ignore-missing-schemas -exit-on-error examples/*.yaml

ruff:
	ruff check apps/echo-agent

go-test:
	cd operator && go test ./...

lint: helm-lint yamllint kubeconform ruff go-test

install-crds:
	helm upgrade --install proompteng-crds charts/proompteng-crds -n proompteng-system --create-namespace

install-operator:
	helm upgrade --install proompteng charts/proompteng-operator -n proompteng-system

run-operator-local:
	KUBECONFIG=$${KUBECONFIG:-$$HOME/.kube/config} go run ./operator/main.go

release-manifest:
	$(PYTHON) hack/update_release_manifest.py --extra-output manifests/proompteng/install.yaml

