# Contributing to proompteng

Thanks for investing time in improving the proompteng agent platform. This guide outlines the expectations and workflows that keep contributions predictable, reviewable, and production-ready.

## Getting Started

1. Review the [README](README.md) for the latest value proposition, architecture overview, and 5-minute quickstart.
2. Install the required toolchain versions via `asdf` or your preferred manager:
   - Go 1.25.1
   - Python 3.13.7
   - Node.js 24.6.1
   - Docker Engine 28.4.0
   - Helm 3.16.2
   - kubectl 1.34.1
   - GitHub CLI 2.80.0
3. Authenticate with `gh auth login --scopes repo,workflow` and `docker login ghcr.io` before pushing container images or interacting with private resources.
4. Fork the repository if you do not have direct push access.

## Branching & Commits

- Create feature branches using `<area>/<ticket>-<summary>` (for example, `operator/PLAT-142-scale-delay`).
- Follow [Conventional Commits](https://www.conventionalcommits.org/) with present-tense summaries (`feat:`, `fix:`, `docs:`, `chore:`, etc.).
- Reference the corresponding Linear issue in the pull request description to preserve traceability.

## Development Workflow

Before opening a pull request, run the following commands from the repository root (`cd proompteng`):

```sh
make lint
make kubeconform
make go-test
ruff check apps/echo-agent --fix
```

Use `make run-operator-local` for iterative debugging against the Kubernetes cluster referenced by `$KUBECONFIG`.

## Roadmap & Support

- Track planned work in the [Linear workspace](https://linear.app/proompteng). Open a ticket before starting work that is not already scoped.
- For security disclosures, email <security@proompteng.ai> instead of filing a public issue.

## Pull Request Checklist

- [ ] Tests and linters from the "Development Workflow" section pass locally.
- [ ] Documentation and examples are updated when behaviour changes.
- [ ] PR description details behaviour deltas, test evidence, and release impact.

Welcome aboard! We appreciate every fix, feature, and feedback loop that sharpens proompteng for the community.
