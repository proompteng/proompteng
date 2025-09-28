# Repository Guidelines

## Mission & Scope
Proompteng anchors Decacorn DevTools' agent infrastructure and this directory is the organization root; keep any sibling repositories here when they exist so shared automation finds consistent paths. All strategy, runbooks, and release notes live in Notion—archive any remaining Confluence pages to prevent drift.

## Project Structure & Module Organization
- `proompteng/apps/echo-agent/`: FastAPI reference agent and SDK primitives.
- `proompteng/operator/`: Kubebuilder controller; reconcilers in `controllers/`, CRDs in `api/v1alpha1/`.
- `proompteng/charts/`: Helm packages for CRDs and operator deployments, mirrored by `examples/` manifests.
- `proompteng/examples/`: Apply-once manifests for smoke testing.
- `proompteng/docs/`: Architecture notes, CI contract, onboarding checklists.
`cd proompteng` before invoking tooling so Makefile paths resolve.

## Toolchain & Environment
Install and pin Go 1.25.1, Python 3.13.7, Node.js 24.6.1, Docker Engine 28.4.0, Helm 3.16.2, kubectl 1.34.1, and GitHub CLI 2.80.0. Configure them via `asdf` or similar, then run `gh auth login --scopes repo,workflow` and `docker login ghcr.io` before pushing images.

## Build, Test & Development Commands
- `make lint`: Helm lint, yamllint, kubeconform, Ruff, and Go unit tests.
- `make go-test`: Executes controller suites in `operator/`.
- `make run-operator-local`: Runs the operator against `$KUBECONFIG` for iterative work.
- `ruff check apps/echo-agent --fix`: Applies Python formatting and linting.
- `make kubeconform`: Confirms rendered manifests match upstream schemas.

## Coding Style & Naming Conventions
Go code is formatted with `gofmt`/`goimports`; keep exported API types versioned in `api/v1alpha1/`. Python uses 4-space indentation, typing, `loguru` logging, and FastAPI handlers named `<verb>_<resource>`. YAML & Helm keep kebab-case keys, `.Values` for configuration, and explanatory template comments. Branch names follow `<area>/<ticket>-<summary>` (e.g., `operator/PLAT-142-scale-delay`).

## Testing Guidelines
Co-locate Go tests as `_test.go` files and favour table-driven cases. Add Python suites under `apps/echo-agent/tests/` and run `pytest` when behaviour expands. Re-run `make kubeconform` after chart or example edits to catch schema drift and keep coverage healthy before merging.

## Commit & Pull Request Guidelines
Use Conventional Commits (`feat:`, `fix:`, `docs:`, `chore:`) with present-tense summaries. PRs document behaviour changes, manual or automated test evidence, linked tasks or issues, and release impact. Communicate deployment windows alongside updates to the release log.
