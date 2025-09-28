#!/usr/bin/env python3
"""Regenerate the consolidated install manifest for release consumers."""
from __future__ import annotations

import argparse
from pathlib import Path
from typing import Iterable

ROOT = Path(__file__).resolve().parents[1]
DEFAULT_RELEASE_DIR = ROOT / "releases" / "kubernetes" / "proompteng"


def load_manifest_paths(release_dir: Path) -> list[Path]:
    crd_dir = release_dir / "crds"
    paths: list[Path] = [
        release_dir / "namespace.yaml",
        *sorted(crd_dir.glob("*.yaml")),
        release_dir / "operator-serviceaccount.yaml",
        release_dir / "operator-rbac.yaml",
        release_dir / "operator-configmap.yaml",
        release_dir / "operator-deployment.yaml",
        release_dir / "operator-service.yaml",
    ]
    missing = [p for p in paths if not p.exists()]
    if missing:
        missing_str = "\n".join(str(p.relative_to(ROOT)) for p in missing)
        raise FileNotFoundError(
            "Missing expected manifest(s):\n" f"{missing_str}"
        )
    return paths


def concatenate_manifests(paths: Iterable[Path]) -> str:
    docs: list[str] = []
    for path in paths:
        docs.append(path.read_text().rstrip())
    return "\n---\n".join(docs) + "\n"


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument(
        "--release-dir",
        type=Path,
        default=DEFAULT_RELEASE_DIR,
        help="Path to the release manifest directory (default: %(default)s)",
    )
    parser.add_argument(
        "--output",
        type=Path,
        default=None,
        help="Optional output file path (defaults to <release-dir>/install.yaml)",
    )
    parser.add_argument(
        "--extra-output",
        action="append",
        default=[],
        type=Path,
        help=(
            "Additional file paths that should also receive the bundled manifest. "
            "Paths are resolved relative to the repository root."
        ),
    )
    args = parser.parse_args()

    release_dir = args.release_dir.resolve()
    if not release_dir.exists():
        raise FileNotFoundError(f"Release directory {release_dir} does not exist")

    manifest_paths = load_manifest_paths(release_dir)
    bundle = concatenate_manifests(manifest_paths)

    default_output = args.output or (release_dir / "install.yaml")
    extra_outputs = [
        path if path.is_absolute() else ROOT / path
        for path in (args.extra_output or [])
    ]

    for path in [default_output, *extra_outputs]:
        path.parent.mkdir(parents=True, exist_ok=True)
        path.write_text(bundle)
        print(f"Wrote {path}")


if __name__ == "__main__":
    main()
